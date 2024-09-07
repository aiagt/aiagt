package clientsuite

import (
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
	//opts = append(opts, client.WithMiddleware(middleware.Auth))
	return &ClientSuite{opts: opts}
}
