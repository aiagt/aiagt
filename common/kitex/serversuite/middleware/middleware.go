package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type Middleware struct {
	authSvc AuthService
}

func NewMiddleware(authSvc AuthService) *Middleware {
	return &Middleware{authSvc: authSvc}
}

func (m *Middleware) Middlewares() []server.Option {
	middles := []endpoint.Middleware{
		m.StreamingStatus,
		m.Logger,
		//m.Transaction,
		m.Auth,
	}

	opts := make([]server.Option, len(middles))
	for i, middle := range middles {
		opts[i] = server.WithMiddleware(middle)
	}

	return opts
}

func ReturnBizErr(ctx context.Context, err error) error {
	ri := rpcinfo.GetRPCInfo(ctx)

	setter, ok := ri.Invocation().(rpcinfo.InvocationSetter)
	if !ok {
		return err
	}

	bizErr, ok := kerrors.FromBizStatusError(err)
	if !ok {
		return err
	}

	setter.SetBizStatusErr(bizErr)

	return nil
}
