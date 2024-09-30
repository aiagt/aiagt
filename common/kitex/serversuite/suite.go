package serversuite

import (
	"github.com/aiagt/aiagt/common/kitex/serversuite/metahandler"
	"github.com/aiagt/aiagt/common/kitex/serversuite/middleware"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
)

type ServerSuite struct {
	opts []server.Option
}

func (s *ServerSuite) Options() []server.Option {
	return s.opts
}

func NewServerSuite(authSvc middleware.AuthService) *ServerSuite {
	var opts []server.Option

	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	opts = append(opts, server.WithMetaHandler(metahandler.NewStreamingMetaHandler()))

	m := middleware.NewMiddleware(authSvc)
	opts = append(opts, m.Middlewares()...)

	return &ServerSuite{opts: opts}
}
