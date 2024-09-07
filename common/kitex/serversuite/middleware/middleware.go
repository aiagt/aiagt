package middleware

import (
	"github.com/cloudwego/kitex/pkg/endpoint"
)

type Middleware struct {
	authSvc AuthService
}

func NewMiddleware(authSvc AuthService) *Middleware {
	return &Middleware{authSvc: authSvc}
}

func (m *Middleware) Middlewares() []endpoint.Middleware {
	return []endpoint.Middleware{m.Auth, m.Transaction}
}
