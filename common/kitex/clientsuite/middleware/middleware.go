package middleware

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Middlewares() []client.Option {
	middles := []endpoint.Middleware{
		m.RequestID,
	}

	opts := make([]client.Option, len(middles))
	for i, middle := range middles {
		opts[i] = client.WithMiddleware(middle)
	}

	return opts
}
