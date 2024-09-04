package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// CreateSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateSecret(ctx context.Context, req *usersvc.CreateSecretReq) (resp *base.Empty, err error) {
	return
}
