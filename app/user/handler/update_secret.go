package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// UpdateSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateSecret(ctx context.Context, req *usersvc.UpdateSecretReq) (resp *base.Empty, err error) {
	return
}
