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
	var (
		id  = ctxutil.UserID(ctx)
		app = mapper.NewModelCreateApp(req, id)
	)

	err = s.appDao.Create(ctx, app)
	if err != nil {
		return nil, bizCreateApp.NewErr(err)
	}

	return
}
