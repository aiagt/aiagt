package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/user/dal/cache"

	"github.com/aiagt/aiagt/apps/user/dal/db"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/client/callopt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	userDao      *db.UserDao
	secretDao    *db.SecretDao
	captchaCache *cache.CaptchaCache
}

func NewUserService(userDao *db.UserDao, secretDao *db.SecretDao, captchaCache *cache.CaptchaCache) *UserServiceImpl {
	initServiceBusiness(1)

	return &UserServiceImpl{userDao: userDao, secretDao: secretDao, captchaCache: captchaCache}
}

type AuthServiceImpl struct {
	handler usersvc.UserService
}

func NewAuthService(handler usersvc.UserService) *AuthServiceImpl {
	return &AuthServiceImpl{handler: handler}
}

func (a *AuthServiceImpl) ParseToken(ctx context.Context, token string, _ ...callopt.Option) (resp int64, err error) {
	return a.handler.ParseToken(ctx, token)
}
