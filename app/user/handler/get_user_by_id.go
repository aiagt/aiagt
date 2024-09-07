package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/user/mapping"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// GetUserByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByID(ctx context.Context, req *base.IDReq) (resp *usersvc.User, err error) {
	user, err := s.userDao.GetByID(ctx, req.Id)
	resp = mapping.NewGenUser(user)

	return
}
