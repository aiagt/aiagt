package ctxutil

import (
	"context"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
)

const (
	Authorization     = "Authorization"
	AuthorizationUser = "AuthorizationUser"
)

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, Authorization, token)
}

func Token(ctx context.Context) string {
	token, _ := ctx.Value(Authorization).(string)
	return token
}

func WithUser(ctx context.Context, user *usersvc.User) context.Context {
	return context.WithValue(ctx, AuthorizationUser, user)
}

func User(ctx context.Context) *usersvc.User {
	user, _ := ctx.Value(AuthorizationUser).(*usersvc.User)
	return user
}

func UserID(ctx context.Context) int64 {
	user, ok := ctx.Value(Authorization).(*usersvc.User)
	if !ok {
		return 0
	}
	return user.Id
}

func Forbidden(ctx context.Context, id int64) bool {
	user, ok := ctx.Value(AuthorizationUser).(*usersvc.User)
	if !ok {
		return true
	}
	return user.Id != id
}
