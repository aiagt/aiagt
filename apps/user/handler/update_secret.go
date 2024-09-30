package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/user/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// UpdateSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateSecret(ctx context.Context, req *usersvc.UpdateSecretReq) (resp *base.Empty, err error) {
	secret, err := s.secretDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizUpdateSecret.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, secret.UserID) {
		return nil, bizUpdateSecret.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.secretDao.Update(ctx, req.Id, mapper.NewModelUpdateSecret(req))
	if err != nil {
		return nil, bizUpdateSecret.NewErr(err)
	}

	return
}
