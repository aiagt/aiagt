package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/user/mapper"
	"github.com/aiagt/aiagt/common/ctxutil"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// ListSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) ListSecret(ctx context.Context, req *usersvc.ListSecretReq) (resp *usersvc.ListSecretResp, err error) {
	userID := ctxutil.UserID(ctx)

	list, page, err := s.secretDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListSecret.NewErr(err)
	}

	resp = &usersvc.ListSecretResp{
		Pagination: page,
		Secrets:    mapper.NewGenListSecret(list),
	}

	return
}
