package main

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/common/cos"
	"github.com/aiagt/aiagt/pkg/closer"
	"github.com/aiagt/aiagt/pkg/jsonutil"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aiagt/aiagt/common/confutil"
	"github.com/aiagt/aiagt/common/hertz/result"
	"github.com/aiagt/aiagt/pkg/logerr"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/google/uuid"

	ktlog "github.com/aiagt/kitextool/option/server/log"
	ktutils "github.com/aiagt/kitextool/utils"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	appcontroller "github.com/aiagt/aiagt/apps/app/controller"
	chatcontroller "github.com/aiagt/aiagt/apps/chat/controller"
	modelcontroller "github.com/aiagt/aiagt/apps/model/controller"
	plugincontroller "github.com/aiagt/aiagt/apps/plugin/controller"
	usercontroller "github.com/aiagt/aiagt/apps/user/controller"
	"github.com/aiagt/aiagt/common/observability"
	"github.com/aiagt/aiagt/rpc"
	ktconf "github.com/aiagt/kitextool/conf"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
	prometheus "github.com/hertz-contrib/monitor-prometheus"
	hertzzap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
)

var conf = new(ServerConf)

func init() {
	confutil.LoadConf(conf,
		".",
		filepath.Join("apps", "gateway"),
	)
}

func main() {
	logger := hertzzap.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelDebug)

	asyncWriter := SetLoggerOutput(&conf.ServerConf)

	consul, registryInfo := observability.InitMetrics(conf.Server.Name, conf.Metrics.Addr, conf.Registry.Address[0])
	p := observability.InitTracing(conf.Server.Name, conf.Tracing.ExportAddr)
	tracer, cfg := tracing.NewServerTracer()

	bindConfig := binding.NewBindConfig()
	bindConfig.UseThirdPartyJSONUnmarshaler(func(data []byte, v interface{}) error {
		return jsonutil.Unmarshal(data, v)
	})

	h := server.Default(server.WithHostPorts(conf.Server.Address),
		server.WithBindConfig(bindConfig),
		server.WithTracer(prometheus.NewServerTracer(
			"",
			"",
			prometheus.WithRegistry(observability.Registry),
			prometheus.WithDisableServer(true),
		)),
		tracer)

	h.Use(func(ctx context.Context, c *app.RequestContext) {
		c.Response.Header.Set("Access-Control-Allow-Origin", "*")

		if string(c.Method()) == http.MethodOptions {
			c.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-Forwarded-For, X-Frame-Options, Accept, Cache-Control")
			c.Response.Header.Set("Access-Control-Allow-Credentials", "true")
			c.Response.Header.Set("Access-Control-Max-Age", "86400")

			c.AbortWithStatus(http.StatusNoContent)

			return
		}

		c.Next(ctx)
	})
	h.Use(tracing.ServerMiddleware(cfg))
	h.Use(accesslog.New())

	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		logerr.Log(consul.Deregister(registryInfo))
		logerr.Log(p.Shutdown(ctx))
		logerr.Log(asyncWriter.Sync())
	})

	r := h.Group("/api/v1")
	usercontroller.RegisterRouter(r, rpc.UserCli)
	modelcontroller.RegisterRouter(r, rpc.ModelCli)
	plugincontroller.RegisterRouter(r, rpc.PluginCli)
	appcontroller.RegisterRouter(r, rpc.AppCli)
	chatcontroller.RegisterRouter(r, rpc.ChatCli, rpc.ChatStreamCli)

	cos.InitCos(conf.Cos.URL, conf.Cos.SecretID, conf.Cos.SecretKey)

	r.GET("/assets/*name", CosAssets)

	r.POST("/assets/avatar", UploadCos(cos.AvatarDir))
	r.POST("/assets/app_logo", UploadCos(cos.AppLogoDir))
	r.POST("/assets/plugin_logo", UploadCos(cos.PluginLogoDir))

	h.Spin()
}

type ServerConf struct {
	ktconf.ServerConf

	Metrics Metrics `yaml:"metrics"`
	Tracing Tracing `yaml:"tracing"`
	Cos     Cos     `yaml:"cos"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}

type Tracing struct {
	ExportAddr string `yaml:"export_addr"`
}

type Cos struct {
	URL       string `yaml:"url"`
	SecretID  string `yaml:"secret_id"`
	SecretKey string `yaml:"secret_key"`
}

func SetLoggerOutput(conf *ktconf.ServerConf) *zapcore.BufferedWriteSyncer {
	if conf.Log.EnableFile != nil && !*conf.Log.EnableFile {
		return nil
	}

	confLog := conf.Log
	ktutils.SetDefault(&confLog.FileName, filepath.Join("log", fmt.Sprintf("%s.log", conf.Server.Name)))
	ktutils.SetDefault(&confLog.MaxSize, ktlog.DefaultMaxSize)
	ktutils.SetDefault(&confLog.MaxAge, ktlog.DefaultMaxAge)
	ktutils.SetDefault(&confLog.MaxBackups, ktlog.DefaultMaxBackups)
	ktutils.SetDefault(&confLog.MaxSize, ktlog.DefaultMaxSize)
	ktutils.SetDefault(
		&confLog.FlushInterval,
		ktutils.Ternary(ktconf.GetEnv() == ktconf.EnvProd, ktlog.DefaultProdFlushInterval, ktlog.DefaultDevFlushInterval),
	)

	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   confLog.FileName,
			MaxSize:    confLog.MaxSize,
			MaxBackups: confLog.MaxBackups,
			MaxAge:     confLog.MaxAge,
		}),
		FlushInterval: time.Duration(confLog.FlushInterval) * time.Second,
	}

	output := io.MultiWriter(os.Stdout, asyncWriter)
	hlog.SetOutput(output)

	return asyncWriter
}

func UploadCos(dir string) app.HandlerFunc {
	const maxFileSize = 5 * 1024 * 1024

	var uploadLimiter = []*rate.Limiter{
		rate.NewLimiter(rate.Every(time.Second), 10),
		rate.NewLimiter(rate.Every(time.Hour), 1000),
	}

	return func(ctx context.Context, c *app.RequestContext) {
		for _, limiter := range uploadLimiter {
			if !limiter.Allow() {
				c.AbortWithStatus(http.StatusTooManyRequests)
				return
			}
		}

		file, err := c.FormFile("file")
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if file.Size > maxFileSize {
			c.AbortWithMsg(fmt.Sprintf("file size exceeds the limit of %d bytes", maxFileSize), http.StatusRequestEntityTooLarge)
			return
		}

		src, err := file.Open()
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer closer.Close(src)

		filename := uuid.New().String() + filepath.Ext(file.Filename)
		path := filepath.Join(dir, filename)

		_, err = cos.Cli.Object.Put(ctx, path, src, nil)
		if err != nil {
			hlog.CtxErrorf(ctx, "[COS] put cos file err: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		hlog.CtxErrorf(ctx, "[COS] put cos file success")
		c.JSON(http.StatusOK, result.Success(utils.H{"file_name": filename, "file_path": path}))
	}
}

func CosAssets(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	if len(name) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	presignedURL, err := cos.Cli.Object.GetPresignedURL(
		ctx,
		http.MethodGet,
		name,
		conf.Cos.SecretID,
		conf.Cos.SecretKey,
		time.Hour,
		nil)
	if err != nil {
		hlog.CtxErrorf(ctx, "[COS] get file %s error: %v", name, err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Redirect(http.StatusFound, []byte(presignedURL.String()))
}
