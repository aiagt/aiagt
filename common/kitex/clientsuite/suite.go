package clientsuite

import (
	"github.com/aiagt/aiagt/common/kitex/clientsuite/middleware"
	"github.com/cloudwego/kitex/client"
)

type ClientSuite struct {
	opts []client.Option
}

func (s *ClientSuite) Options() []client.Option {
	return s.opts
}

func NewClientSuite() *ClientSuite {
	var opts []client.Option

	m := middleware.NewMiddleware()
	opts = append(opts, m.Middlewares()...)

	return &ClientSuite{opts: opts}
}
