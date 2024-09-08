package ctxutil

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

const (
	Authorization       = "Authorization"
	AuthorizationUserID = "AuthorizationUserID"
)

func WithToken(ctx context.Context, token string) context.Context {
	return metainfo.WithPersistentValue(ctx, Authorization, token)
}

func Token(ctx context.Context) string {
	token, _ := metainfo.GetPersistentValue(ctx, Authorization)
	return token
}

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, AuthorizationUserID, userID)
}

func UserID(ctx context.Context) int64 {
	userID, _ := ctx.Value(AuthorizationUserID).(int64)
	return userID
}

func GetUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(AuthorizationUserID).(int64)
	return userID, ok
}

func Forbidden(ctx context.Context, id int64) bool {
	userID, ok := ctx.Value(AuthorizationUserID).(int64)
	if !ok {
		return true
	}
	return userID != id
}
