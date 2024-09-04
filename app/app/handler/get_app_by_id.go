package handler

import (
	"context"

	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// GetAppByID implements the AppServiceImpl interface.
func (s *AppServiceImpl) GetAppByID(ctx context.Context, req *base.IDReq) (resp *appsvc.App, err error) {
	return
}
