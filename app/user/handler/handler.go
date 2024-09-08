package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/user/dal/cache"

	"github.com/aiagt/aiagt/app/user/dal/db"
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

func (a *AuthServiceImpl) GetUser(ctx context.Context, _ ...callopt.Option) (r *usersvc.User, err error) {
	return a.handler.GetUser(ctx)
}
