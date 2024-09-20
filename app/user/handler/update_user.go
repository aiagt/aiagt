package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/user/mapper"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *usersvc.UpdateUserReq) (resp *base.Empty, err error) {
	user := mapper.NewModelUpdateUser(req)

	err = s.userDao.Update(ctx, req.Id, user)
	if err != nil {
		return nil, bizUpdateUser.NewErr(err)
	}

	return
}
