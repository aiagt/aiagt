package serversuite

import (
	"github.com/aiagt/aiagt/common/kitex/serversuite/metahandler"
	"github.com/aiagt/aiagt/common/kitex/serversuite/middleware"
	"github.com/aiagt/aiagt/common/observability"
	ktconf "github.com/aiagt/kitextool/conf"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type ServerSuite struct {
	opts []server.Option
}

func (s *ServerSuite) Options() []server.Option {
	return s.opts
}

func NewServerSuite(conf *ktconf.ServerConf, authSvc middleware.AuthService) *ServerSuite {
	var opts []server.Option

	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	opts = append(opts, server.WithMetaHandler(metahandler.NewStreamingMetaHandler()))
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.Server.Name}))

	opts = append(opts, server.WithTracer(prometheus.NewServerTracer(
		"",
		"",
		prometheus.WithDisableServer(true), // disable built-in server
		prometheus.WithRegistry(observability.Registry)),
	))

	opts = append(opts, server.WithSuite(tracing.NewServerSuite()))

	m := middleware.NewMiddleware(authSvc)
	opts = append(opts, m.Middlewares()...)

	return &ServerSuite{opts: opts}
}
