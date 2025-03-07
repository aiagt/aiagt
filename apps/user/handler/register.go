package handler

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aiagt/aiagt/apps/user/mapper"
	"github.com/aiagt/aiagt/apps/user/pkg/encrypt"
	"github.com/aiagt/aiagt/apps/user/pkg/jwt"
	"github.com/aiagt/aiagt/common/baseutil"

	"github.com/aiagt/aiagt/apps/user/dal/cache"

	"github.com/aiagt/aiagt/apps/user/model"
	"github.com/aiagt/aiagt/common/bizerr"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/pkg/errors"
)

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *usersvc.RegisterReq) (resp *usersvc.RegisterResp, err error) {
	if !validateEmail(req.Email) {
		return nil, bizRegister.NewCodeErr(11, errors.New("invalid email"))
	}

	if req.Password != nil && !validatePassword(*req.Password) {
		return nil, bizRegister.NewCodeErr(12, errors.New("invalid password"))
	}

	if req.Username != nil && !validateUsername(*req.Username) {
		return nil, bizRegister.NewCodeErr(13, errors.New("invalid username"))
	}

	valid, err := s.captchaCache.Verify(ctx, cache.CaptchaTypeAuth, req.Email, req.Captcha)
	if err != nil {
		return nil, bizRegister.NewErr(err)
	}

	if !valid {
		return nil, bizRegister.CodeErr(bizerr.ErrCodeWrongAuth)
	}

	user := &model.User{
		Email:  req.Email,
		Avatar: fmt.Sprintf("https://api.multiavatar.com/%s.svg", req.Email),
	}

	if req.Username != nil {
		user.Username = *req.Username
	} else {
		user.Username = req.Email
	}

	if req.Password != nil {
		user.Password = encrypt.Encrypt(*req.Password)
	}

	err = s.userDao.Create(ctx, user)
	if err != nil {
		return nil, bizRegister.NewErr(err)
	}

	token, expire, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, bizRegister.NewErr(err)
	}

	resp = &usersvc.RegisterResp{
		Token:  token,
		Expire: baseutil.NewBaseTime(*expire),
		User:   mapper.NewGenUser(user),
	}

	return
}

func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func validatePassword(password string) bool {
	// at least 8 characters, must contain both letters + numbers or symbols
	if len(password) < 8 {
		return false
	}

	if !regexp.MustCompile(`[a-zA-Z]`).MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`[\d\W]`).MatchString(password) {
		return false
	}

	return true
}

func validateUsername(username string) bool {
	return len(username) > 0
}
