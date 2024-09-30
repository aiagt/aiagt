package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/app/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// UpdateApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) UpdateApp(ctx context.Context, req *appsvc.UpdateAppReq) (resp *base.Empty, err error) {
	app, err := s.appDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizUpdateApp.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, app.AuthorID) {
		return nil, bizUpdateApp.CodeErr(bizerr.ErrCodeForbidden)
	}

	labelIDs, err := s.labelDao.UpdateLabels(ctx, req.LabelIds, req.LabelTexts)
	if err != nil {
		return nil, bizCreateApp.NewErr(err)
	}

	err = s.appDao.Update(ctx, req.Id, mapper.NewModelUpdateApp(req, labelIDs))
	if err != nil {
		return nil, bizUpdateApp.NewErr(err)
	}

	return
}
