package handler

import (
	"context"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// ListSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) ListSecret(ctx context.Context, req *usersvc.ListSecretReq) (resp *usersvc.ListSecretResp, err error) {
	return
}
