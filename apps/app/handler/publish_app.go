package handler

import (
	"context"
	"time"

	"github.com/aiagt/aiagt/pkg/utils"

	"github.com/aiagt/aiagt/apps/app/model"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// PublishApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) PublishApp(ctx context.Context, req *appsvc.PublishAppReq) (resp *base.Empty, err error) {
	app, err := s.appDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizPublishApp.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, app.AuthorID) {
		return nil, bizPublishApp.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.appDao.Update(ctx, req.Id, &model.AppOptional{
		PublishedAt: utils.Pointer(time.Now()),
		Version:     utils.Pointer(req.Version),
	})
	if err != nil {
		return nil, bizPublishApp.NewErr(err)
	}

	return
}
