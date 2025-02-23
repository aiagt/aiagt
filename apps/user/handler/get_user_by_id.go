package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/user/mapper"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// GetUserByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByID(ctx context.Context, req *base.IDReq) (resp *usersvc.User, err error) {
	user, err := s.userDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetUserById.NewErr(err)
	}

	resp = mapper.NewGenUser(user)

	return
}
