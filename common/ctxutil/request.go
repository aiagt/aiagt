package ctxutil

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/google/uuid"
)

const (
	RequestIDKey = "REQUEST_ID"
)

func WithRequestID(ctx context.Context) context.Context {
	if _, ok := GetRequestID(ctx); ok {
		return ctx
	}

	id := uuid.New().String()

	return metainfo.WithPersistentValue(ctx, RequestIDKey, id)
}

func RequestID(ctx context.Context) string {
	id, _ := GetRequestID(ctx)
	return id
}

func GetRequestID(ctx context.Context) (string, bool) {
	return metainfo.GetPersistentValue(ctx, RequestIDKey)
}
