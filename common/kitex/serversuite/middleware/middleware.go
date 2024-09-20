package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type Middleware struct {
	authSvc AuthService
}

func NewMiddleware(authSvc AuthService) *Middleware {
	return &Middleware{authSvc: authSvc}
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
