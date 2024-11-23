package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aiagt/aiagt/pkg/jsonutil"
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

	r.Use(StaticFSLimiter)

	r.StaticFS("/assets", &app.FS{
		Root:        "./assets",
		PathRewrite: app.NewPathSlashesStripper(3),
	})
	r.POST("/assets", UploadAssets)

	h.Spin()
}

type ServerConf struct {
	ktconf.ServerConf

	Metrics Metrics `yaml:"metrics"`
	Tracing Tracing `yaml:"tracing"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}

type Tracing struct {
	ExportAddr string `yaml:"export_addr"`
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

func UploadAssets(ctx context.Context, c *app.RequestContext) {
	file, err := c.FormFile("file")
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)

	err = c.SaveUploadedFile(file, fmt.Sprintf("assets/%s", filename))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, result.Success(utils.H{"filename": filename}))
}

func StaticFSLimiter(ctx context.Context, c *app.RequestContext) {
	const maxFileSize = 5 * 1024 * 1024

	if !bytes.HasPrefix(c.Request.URI().Path(), []byte("/api/v1/assets")) {
		c.Next(ctx)
		return
	}

	filePath := "./assets" + string(c.Request.URI().Path())
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if fileInfo.Size() > maxFileSize {
		c.AbortWithMsg(fmt.Sprintf("file size exceeds the maximum limit of %d bytes", maxFileSize), http.StatusForbidden)
		return
	}

	c.Next(ctx)
}
