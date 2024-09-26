package ctxutil

import (
	"context"

	"github.com/cloudwego/kitex/pkg/utils/contextmap"
)

const StreamingStatus CtxKey = "STREAMING_STATUS"

func WithStreamingStatus(ctx context.Context) context.Context {
	if m, ok := contextmap.GetContextMap(ctx); ok {
		m.Store(StreamingStatus, true)
	}

	return ctx
}

func IsStreaming(ctx context.Context) bool {
	if m, ok := contextmap.GetContextMap(ctx); ok {
		v, _ := m.Load(StreamingStatus)
		return v.(bool)
	}

	return false
}
