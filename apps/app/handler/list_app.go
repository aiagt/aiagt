package handler

import (
	"context"

	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/pkg/utils"

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

	var authorMap hmap.Map[int64, *usersvc.User]

	if utils.Value(req.WithAuthor) {
		authorIDs := lists.Map(apps, func(t *model.App) int64 { return t.AuthorID })

		authors, err := s.userCli.GetUserByIds(ctx, &base.IDsReq{Ids: authorIDs})
		if err != nil {
			return nil, bizListApp.NewErr(err)
		}

		authorMap = hmap.FromSliceEntries(authors, func(t *usersvc.User) (int64, *usersvc.User, bool) {
			return t.Id, t, true
		})
	}

	resp = &appsvc.ListAppResp{
		Apps:       mapper.NewGenListApp(apps, labelMap, authorMap),
		Pagination: page,
	}

	return
}
