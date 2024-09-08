package handler

import (
	"context"

	"github.com/aiagt/aiagt/kitex_gen/base"

	"github.com/aiagt/aiagt/app/user/dal/cache"
	"github.com/aiagt/aiagt/app/user/model"
	"github.com/aiagt/aiagt/app/user/pkg/encrypt"
	"github.com/aiagt/aiagt/common/bizerr"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// ResetPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ResetPassword(ctx context.Context, req *usersvc.ResetPasswordReq) (resp *base.Empty, err error) {
	captcha, err := s.captchaCache.GetAndDel(ctx, cache.CaptchaTypeReset, req.Email)
	if err != nil {
		return nil, bizResetPassword.NewErr(err)
	}

	if captcha == req.Captcha {
		return nil, bizResetPassword.CodeErr(bizerr.ErrCodeWrongAuth)
	}

	user, err := s.userDao.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, bizResetPassword.NewErr(err)
	}

	passwordHashed := encrypt.Encrypt(req.Password)

	err = s.userDao.Update(ctx, user.ID, &model.UserOptional{
		Password: &passwordHashed,
	})
	if err != nil {
		return nil, bizResetPassword.NewErr(err)
	}

	return
}
