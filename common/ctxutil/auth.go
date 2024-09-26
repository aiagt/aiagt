package ctxutil

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

const (
	Authorization              = "AUTHORIZATION"
	AuthorizationUserID CtxKey = "AUTHORIZATION_USER_ID"
)

func WithToken(ctx context.Context, token string) context.Context {
	return metainfo.WithPersistentValue(ctx, Authorization, token)
}

func Token(ctx context.Context) string {
	token, _ := metainfo.GetPersistentValue(ctx, Authorization)
	return token
}

func WithUserID(ctx context.Context, userID int64) context.Context {
	if IsStreaming(ctx) {
		ctx = WithMapValue(ctx, AuthorizationUserID, userID)
	} else {
		ctx = context.WithValue(ctx, AuthorizationUserID, userID)
	}

	return ctx
}

func GetUserID(ctx context.Context) (int64, bool) {
	if IsStreaming(ctx) {
		return GetMapValue[int64](ctx, AuthorizationUserID)
	}

	userID, ok := ctx.Value(AuthorizationUserID).(int64)
	return userID, ok
}

func UserID(ctx context.Context) int64 {
	userID, _ := GetUserID(ctx)
	return userID
}

func Forbidden(ctx context.Context, id int64) bool {
	userID, ok := GetUserID(ctx)
	if !ok {
		return false
	}

	return userID != id
}
