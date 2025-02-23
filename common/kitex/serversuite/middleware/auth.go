package middleware

import (
	"context"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/hash/hset"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
)

type AuthService interface {
	GenToken(ctx context.Context, id int64, options ...callopt.Option) (resp string, err error)
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
			case "Login", "Register", "ParseToken", "GenToken", "SendCaptcha", "ResetPassword":
				return next(ctx, req, resp)
			}
		}

		token := ctxutil.Token(ctx)

		id, err := m.authSvc.ParseToken(ctx, token)
		if err == nil {
			return next(ctxutil.WithUserID(ctx, id), req, resp)
		}

		biz := bizerr.NewBiz(serviceName, "auth", 40000)

		if serviceMethods, ok := withoutAuthenticationMethod[serviceName]; ok {
			if _, ok := serviceMethods[methodName]; ok {
				token, err = m.authSvc.GenToken(ctx, id)
				if err != nil {
					return ReturnBizErr(ctx, biz.CodeErr(bizerr.ErrCodeServerFailure).Logf(ctx, "gen token error: %v", err.Error()))
				}

				ctx = ctxutil.WithToken(ctx, token)
				ctx = ctxutil.WithUserID(ctx, id)

				return next(ctx, req, resp)
			}
		}

		return ReturnBizErr(ctx, biz.CodeErr(bizerr.ErrCodeUnauthorized).Logf(ctx, "parse token error: %v", err.Error()))
	}
}

var withoutAuthenticationMethod = hmap.Map[string, hset.Set[string]]{
	"app":   hset.FromValues("ListApp", "ListAppLabel", "GetAppByID"),
	"chat":  hset.FromValues("ListConversation"),
	"model": hset.FromValues("ListModel"),
}
