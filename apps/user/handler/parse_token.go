package handler

import (
	"context"
	"github.com/aiagt/aiagt/apps/user/pkg/jwt"
)

// ParseToken implements the UserServiceImpl interface.
func (s *UserServiceImpl) ParseToken(ctx context.Context, token string) (resp int64, err error) {
	resp, err = jwt.ParseToken(token)
	if err != nil {
		err = bizParseToken.NewErr(err).Log(ctx, "parse json web token error")
	}

	return
}
