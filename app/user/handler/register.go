package handler

import (
	"context"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *usersvc.RegisterReq) (resp *usersvc.RegisterResp, err error) {
	return
}
