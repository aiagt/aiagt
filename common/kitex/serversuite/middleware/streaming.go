package middleware

import (
	"context"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/streaming"
	"go.opentelemetry.io/otel/trace"
)

func (m *Middleware) StreamingStatus(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		if _, ok := req.(*streaming.Args); ok {
			// inject streaming status in streaming api
			ctx = ctxutil.WithStreamingStatus(ctx)

			// kitex obs-telemetry does not support thrift streaming scenarios, needs to be injected manually
			span := trace.SpanFromContext(ctx)
			ctx = ctxutil.WithSpan(ctx, span)
		}

		return next(ctx, req, resp)
	}
}
