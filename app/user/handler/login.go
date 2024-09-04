package handler

import (
	"context"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *usersvc.LoginReq) (resp *usersvc.LoginResp, err error) {
	return
}
