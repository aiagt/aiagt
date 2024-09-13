package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) DeleteApp(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	app, err := s.appDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteApp.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, app.AuthorID) {
		return nil, bizDeleteApp.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.appDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteApp.NewErr(err)
	}

	return
}
