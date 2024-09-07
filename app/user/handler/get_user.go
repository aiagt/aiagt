package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/user/mapping"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context) (resp *usersvc.User, err error) {
	// TODO: parse token

	id := int64(1)

	user, err := s.userDao.GetByID(ctx, id)
	resp = mapping.NewGenUser(user)

	return
}
