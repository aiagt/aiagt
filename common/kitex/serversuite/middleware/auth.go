package middleware

import (
	"context"
	"github.com/aiagt/aiagt/common/bizerr"

	"github.com/aiagt/aiagt/app/user/pkg/jwt"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
)

type AuthService interface {
	GetUser(ctx context.Context, callOptions ...callopt.Option) (r *usersvc.User, err error)
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
			case "Login", "Register", "SendCaptcha":
				return next(ctx, req, resp)
			}
		}

		token := ctxutil.Token(ctx)
		id, err := jwt.ParseToken(token)
		if err != nil {
			biz := bizerr.NewBiz("auth", "auth", 4000000)
			return ReturnBizErr(ctx, biz.CodeErr(bizerr.ErrCodeUnauthorized))
		}

		return next(ctxutil.WithUserID(ctx, id), req, resp)
	}
}
