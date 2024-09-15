package handler

import (
	"context"
	"github.com/aiagt/aiagt/app/user/dal/cache"

	"github.com/aiagt/aiagt/app/user/mapper"
	"github.com/aiagt/aiagt/app/user/pkg/encrypt"
	"github.com/aiagt/aiagt/app/user/pkg/jwt"
	"github.com/aiagt/aiagt/common/baseutil"

	"github.com/aiagt/aiagt/app/user/model"
	"github.com/aiagt/aiagt/common/bizerr"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *usersvc.LoginReq) (resp *usersvc.LoginResp, err error) {
	var user *model.User

	switch {
	case req.Captcha != nil:
		var captcha string

		captcha, err = s.captchaCache.GetAndDel(ctx, cache.CaptchaTypeAuth, req.Email)
		if err != nil {
			return nil, bizLogin.NewErr(err)
		}

		if captcha != *req.Captcha {
			return nil, bizLogin.CodeErr(bizerr.ErrCodeWrongAuth)
		}

		user, err = s.userDao.GetByEmail(ctx, req.Email)
		if err != nil {
			return nil, bizLogin.NewErr(err)
		}
	case req.Password != nil:
		user, err = s.userDao.GetByEmail(ctx, req.Email)
		if err != nil {
			return nil, bizLogin.NewErr(err)
		}

		if user.Password != encrypt.Encrypt(*req.Password) {
			return nil, bizLogin.CodeErr(bizerr.ErrCodeWrongAuth)
		}
	default:
		return nil, bizLogin.CodeErr(bizerr.ErrCodeWrongAuth)
	}

	token, expire, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, bizLogin.NewErr(err)
	}

	resp = &usersvc.LoginResp{
		Token:  token,
		Expire: baseutil.NewBaseTime(*expire),
		User:   mapper.NewGenUser(user),
	}

	return
}
