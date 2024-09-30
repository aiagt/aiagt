package middleware

import (
	"context"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
)

func (m *Middleware) Logger(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		methodName, _ := kitexutil.GetMethod(ctx)
		requestID := ctxutil.RequestID(ctx)

		// streaming api
		if ctxutil.IsStreaming(ctx) {
			klog.CtxInfof(ctx, "[STREAMING] %s, request_id: %s", methodName, requestID)
			return next(ctx, req, resp)
		}

		// pin pong api
		klog.CtxInfof(ctx, "[REQUEST] %s, request_id: %s, body: %v", methodName, requestID, utils.Pretty(req, 1<<10))

		err = next(ctx, req, resp)
		if err != nil {
			klog.CtxErrorf(ctx, "[RESPONSE] %s, request_id: %s, error: %v", methodName, requestID, err.Error())
			return err
		}

		klog.CtxInfof(ctx, "[RESPONSE] %s, request_id: %s, body: %v", methodName, requestID, utils.Pretty(resp, 1<<10))

		return nil
	}
}
