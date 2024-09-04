package handler

import (
	"context"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// ForgotPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ForgotPassword(ctx context.Context, req *usersvc.ForgotPasswordReq) (resp *usersvc.ForgotPasswordResp, err error) {
	return
}
