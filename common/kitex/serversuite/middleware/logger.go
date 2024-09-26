package middleware

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/pkg/safe"
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
		klog.CtxInfof(ctx, "[REQUEST] %s, request_id: %s, body: %v", methodName, requestID, pretty(req, 1<<10))

		err = next(ctx, req, resp)
		if err != nil {
			klog.CtxErrorf(ctx, "[RESPONSE] %s, request_id: %s, error: %v", methodName, requestID, err.Error())
			return err
		}

		klog.CtxInfof(ctx, "[RESPONSE] %s, request_id: %s, body: %v", methodName, requestID, pretty(resp, 1<<10))

		return nil
	}
}

func pretty(v any, max int) string {
	resultBytes := safe.UnsafeValue(json.Marshal(v))

	if max > 0 && len(resultBytes) > max {
		builder := bytes.NewBuffer(resultBytes[:max])
		builder.WriteString("...")

		return builder.String()
	}

	return string(resultBytes)
}
