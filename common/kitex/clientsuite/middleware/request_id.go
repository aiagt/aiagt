package middleware

import (
	"context"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func (m *Middleware) RequestID(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ctx = ctxutil.WithRequestID(ctx)
		return next(ctx, req, resp)
	}
}
