package middleware

import (
	"context"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/streaming"
)

func (m *Middleware) StreamingStatus(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		if _, ok := req.(*streaming.Args); ok {
			// inject streaming status in streaming api
			ctx = ctxutil.WithStreamingStatus(ctx)
		}
		return next(ctx, req, resp)
	}
}
