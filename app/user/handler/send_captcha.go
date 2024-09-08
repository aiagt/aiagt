package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/user/dal/cache"

	"github.com/aiagt/aiagt/app/user/pkg/captcha"
	"github.com/aiagt/aiagt/app/user/pkg/email"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// SendCaptcha implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendCaptcha(ctx context.Context, req *usersvc.SendCaptchaReq) (resp *usersvc.SendCaptchaResp, err error) {
	resp = new(usersvc.SendCaptchaResp)

	_, err = s.userDao.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, bizSendCaptcha.NewErr(err)
	}
	if err == nil {
		resp.Exists = true
	}

	cpt := captcha.Generate()

	err = s.captchaCache.Set(ctx, cache.NewCaptchaType(req.Type), req.Email, cpt)
	if err != nil {
		return nil, bizSendCaptcha.NewErr(err)
	}

	err = email.SendAuthCaptcha(cpt, req.Email)
	if err != nil {
		return nil, bizSendCaptcha.NewErr(err)
	}

	return
}
