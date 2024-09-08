package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteSecret(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	secret, err := s.secretDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteSecret.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, secret.UserID) {
		return nil, bizDeleteSecret.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.secretDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteSecret.NewErr(err)
	}

	return
}
