package ctxutil

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

const NoAuthKey = "NO_AUTH"

func WithNoAuth(ctx context.Context) context.Context {
	return metainfo.WithPersistentValue(ctx, NoAuthKey, NoAuthKey)
}

func RemoveNoAuth(ctx context.Context) context.Context {
	return metainfo.DelPersistentValue(ctx, NoAuthKey)
}

func NoAuth(ctx context.Context) bool {
	_, ok := metainfo.GetPersistentValue(ctx, NoAuthKey)
	return ok
}
