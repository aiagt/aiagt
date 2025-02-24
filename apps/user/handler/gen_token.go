package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/user/pkg/jwt"
)

// GenToken implements the UserServiceImpl interface.
func (s *UserServiceImpl) GenToken(ctx context.Context, token int64) (resp string, err error) {
	resp, _, err = jwt.GenerateToken(token)
	if err != nil {
		err = bizGenToken.NewErr(err).Log(ctx, "generate json web token error")
	}

	return
}
