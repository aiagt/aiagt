package middleware

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
)

type AuthService interface {
	ParseToken(ctx context.Context, token string, callOptions ...callopt.Option) (resp int64, err error)
}

func (m *Middleware) Auth(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		var (
			serviceName   string
			methodName, _ = kitexutil.GetMethod(ctx)
		)

		ri := rpcinfo.GetRPCInfo(ctx)
		if ri.To() != nil {
			serviceName = ri.To().ServiceName()
		}

		switch serviceName {
		case "user":
			switch methodName {
			case "Login", "Register", "ParseToken", "SendCaptcha", "ResetPassword":
				return next(ctx, req, resp)
			}
		}

		token := ctxutil.Token(ctx)

		id, err := m.authSvc.ParseToken(ctx, token)
		if err != nil {
			biz := bizerr.NewBiz(serviceName, "auth", 40000)
			return ReturnBizErr(ctx, biz.CodeErr(bizerr.ErrCodeUnauthorized).Logf(ctx, "parse token error: %v", err.Error()))
		}

		return next(ctxutil.WithUserID(ctx, id), req, resp)
	}
}
