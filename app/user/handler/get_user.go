package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/user/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context) (resp *usersvc.User, err error) {
	id, ok := ctxutil.GetUserID(ctx)
	if !ok {
		return nil, bizGetUser.CodeErr(bizerr.ErrCodeUnauthorized)
	}

	user, err := s.userDao.GetByID(ctx, id)
	if err != nil {
		return nil, bizGetUser.NewErr(err)
	}

	resp = mapper.NewGenUser(user)

	return
}
