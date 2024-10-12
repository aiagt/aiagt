package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"go.uber.org/zap"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/common/logger"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
)

func (m *Middleware) Logger(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		var (
			serviceName   string
			methodName, _ = kitexutil.GetMethod(ctx)
		)

		ri := rpcinfo.GetRPCInfo(ctx)
		if ri.To() != nil {
			serviceName = ri.To().ServiceName()
		}

		// streaming api
		if ctxutil.IsStreaming(ctx) {
			logger.With(
				zap.String("type", "streaming"),
				zap.String("service", serviceName),
				zap.String("method", methodName),
			).CtxInfof(ctx, "streaming request")

			return next(ctx, req, resp)
		}

		// pin pong api
		logger.With(
			zap.String("type", "request"),
			zap.String("service", serviceName),
			zap.String("method", methodName),
			zap.String("body", utils.Pretty(req, 1<<10)),
		).CtxInfof(ctx, "request")

		err = next(ctx, req, resp)
		if err != nil {
			logger.With(
				zap.String("type", "response"),
				zap.String("service", serviceName),
				zap.String("method", methodName),
				zap.String("error", err.Error()),
			).CtxErrorf(ctx, "response error")

			return err
		}

		logger.With(
			zap.String("type", "response"),
			zap.String("service", serviceName),
			zap.String("method", methodName),
			zap.String("body", utils.Pretty(resp, 1<<10)),
		).CtxInfof(ctx, "response")

		return nil
	}
}
