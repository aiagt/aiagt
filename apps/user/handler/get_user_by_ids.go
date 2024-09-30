package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/user/mapper"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// GetUserByIds implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByIds(ctx context.Context, req *base.IDsReq) (resp []*usersvc.User, err error) {
	users, err := s.userDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		return nil, bizGetUserByIds.NewCodeErr(1, err)
	}

	resp = mapper.NewGenListUser(users)

	return
}
