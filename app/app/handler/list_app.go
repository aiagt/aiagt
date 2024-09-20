package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/app/mapper"
	"github.com/aiagt/aiagt/common/ctxutil"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
)

// ListApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) ListApp(ctx context.Context, req *appsvc.ListAppReq) (resp *appsvc.ListAppResp, err error) {
	userID := ctxutil.UserID(ctx)

	apps, page, err := s.appDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListApp.NewErr(err)
	}

	resp = &appsvc.ListAppResp{
		Apps:       mapper.NewGenListApp(apps),
		Pagination: page,
	}

	return
}
