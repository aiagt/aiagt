package serversuite

import (
	"github.com/aiagt/aiagt/common/kitex/serversuite/metahandler"
	"github.com/aiagt/aiagt/common/kitex/serversuite/middleware"
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

	m := middleware.NewMiddleware(authSvc)
	opts = append(opts, server.WithMiddleware(m.Transaction))
	opts = append(opts, server.WithMiddleware(m.Auth))
	opts = append(opts, server.WithMetaHandler(metahandler.NewStreamingMetaHandler()))

	return &ServerSuite{opts: opts}
}
