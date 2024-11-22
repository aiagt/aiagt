package handler

import (
	"context"

	"github.com/aiagt/aiagt/pkg/lists"

	"github.com/aiagt/aiagt/apps/app/mapper"
	"github.com/aiagt/aiagt/apps/app/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
)

// ListApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) ListApp(ctx context.Context, req *appsvc.ListAppReq) (resp *appsvc.ListAppResp, err error) {
	userID := ctxutil.UserID(ctx)

	apps, page, err := s.appDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListApp.NewErr(err)
	}

	labelIDs := lists.FlatMap(apps, func(t *model.App) []int64 {
		return t.LabelIDs
	})
	labels, err := s.labelDao.GetByIDs(ctx, labelIDs)

	labelMap := hmap.FromSliceEntries(labels, func(t *model.AppLabel) (int64, *appsvc.AppLabel, bool) {
		return t.ID, mapper.NewGenAppLabel(t), true
	})

	resp = &appsvc.ListAppResp{
		Apps:       mapper.NewGenListApp(apps, labelMap),
		Pagination: page,
	}

	return
}
