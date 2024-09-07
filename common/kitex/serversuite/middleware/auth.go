package middleware

import (
	"context"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
	"github.com/pkg/errors"
)

type AuthService interface {
	GetUser(ctx context.Context, callOptions ...callopt.Option) (r *usersvc.User, err error)
}

func (m *Middleware) Auth(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		serviceName, _ := kitexutil.GetCaller(ctx)
		methodName, _ := kitexutil.GetMethod(ctx)

		switch serviceName {
		case "user":
			switch methodName {
			case "login", "register":
				return next(ctx, req, resp)
			}
		}

		user, err := m.authSvc.GetUser(ctx)
		if err != nil {
			return errors.Wrap(err, "auth")
		}

		return next(ctxutil.WithUser(ctx, user), req, resp)
	}
}
