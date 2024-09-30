package main

import (
	"context"
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
	"github.com/hertz-contrib/monitor-prometheus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"path/filepath"
)

var conf = new(ServerConf)

func init() {
	ktconf.LoadFiles(conf, "conf.yaml",
		filepath.Join("apps", "gateway", "conf.yaml"))
}

func main() {
	consul, registryInfo := observability.InitMetrics(conf.Server.Name, conf.Metrics.Addr, conf.Registry.Address[0])
	p := observability.InitTracing(conf.Server.Name)

	tracer, cfg := tracing.NewServerTracer()

	h := server.Default(server.WithHostPorts(conf.Server.Address),
		server.WithTracer(prometheus.NewServerTracer(
			"",
			"",
			prometheus.WithRegistry(observability.Registry),
			prometheus.WithDisableServer(true),
		)),
		tracer)

	h.Use(accesslog.New())
	h.Use(tracing.ServerMiddleware(cfg))

	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		_ = consul.Deregister(registryInfo)
		_ = p.Shutdown(ctx)
	})

	r := h.Group("/api/v1")
	usercontroller.RegisterRouter(r, rpc.UserCli)
	modelcontroller.RegisterRouter(r, rpc.ModelCli)
	plugincontroller.RegisterRouter(r, rpc.PluginCli)
	appcontroller.RegisterRouter(r, rpc.AppCli)
	chatcontroller.RegisterRouter(r, rpc.ChatCli, rpc.ChatStreamCli)

	h.Spin()
}

type ServerConf struct {
	ktconf.ServerConf

	Metrics Metrics `yaml:"metrics"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}
