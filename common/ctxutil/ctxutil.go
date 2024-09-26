package ctxutil

import (
	"context"

	"github.com/cloudwego/kitex/pkg/utils/contextmap"
)

type CtxKey string

func WithMapValue[T any](ctx context.Context, key CtxKey, val T) context.Context {
	if m, ok := contextmap.GetContextMap(ctx); ok {
		m.Store(key, val)
	}

	return ctx
}

func GetMapValue[T any](ctx context.Context, key CtxKey) (T, bool) {
	if m, ok := contextmap.GetContextMap(ctx); ok {
		raw, ok := m.Load(key)
		if ok {
			val, ok := raw.(T)
			return val, ok
		}
	}

	var zero T

	return zero, false
}
