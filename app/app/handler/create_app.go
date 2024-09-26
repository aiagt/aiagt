package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/app/mapper"
	"github.com/aiagt/aiagt/common/ctxutil"

	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// CreateApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) CreateApp(ctx context.Context, req *appsvc.CreateAppReq) (resp *base.Empty, err error) {
	id := ctxutil.UserID(ctx)

	labelIDs, err := s.labelDao.UpdateLabels(ctx, req.LabelIds, req.LabelTexts)
	if err != nil {
		return nil, bizCreateApp.NewErr(err)
	}

	app := mapper.NewModelCreateApp(req, id, labelIDs)

	err = s.appDao.Create(ctx, app)
	if err != nil {
		return nil, bizCreateApp.NewErr(err)
	}

	return
}
