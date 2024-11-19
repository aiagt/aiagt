package handler

import (
	"context"
	"errors"

	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/apps/user/mapper"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *usersvc.UpdateUserReq) (resp *base.Empty, err error) {
	userID := ctxutil.UserID(ctx)
	user := mapper.NewModelUpdateUser(req)

	if req.Email != nil && !validateEmail(*req.Email) {
		return nil, bizRegister.NewCodeErr(11, errors.New("invalid email"))
	}

	if req.Username != nil && !validateUsername(*req.Username) {
		return nil, bizRegister.NewCodeErr(13, errors.New("invalid username"))
	}

	err = s.userDao.Update(ctx, userID, user)
	if err != nil {
		return nil, bizUpdateUser.NewErr(err)
	}

	return
}
