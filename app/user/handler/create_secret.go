package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/app/user/mapper"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// CreateSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateSecret(ctx context.Context, req *usersvc.CreateSecretReq) (resp *base.Empty, err error) {
	secret := mapper.NewModelCreateSecret(req)
	secret.UserID = ctxutil.UserID(ctx)

	err = s.secretDao.Create(ctx, secret)
	if err != nil {
		return nil, bizCreateSecret.NewErr(err)
	}

	return
}
