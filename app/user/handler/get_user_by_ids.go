package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/user/mapping"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// GetUserByIds implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByIds(ctx context.Context, req *base.IDsReq) (resp []*usersvc.User, err error) {
	users, err := s.userDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		return nil, bizGetUserByIDs.NewCodeErr(1, err)
	}

	resp = mapping.NewGenListUser(users)

	return
}
