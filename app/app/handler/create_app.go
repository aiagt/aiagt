package handler

import (
	"context"

	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// CreateApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) CreateApp(ctx context.Context, req *appsvc.CreateAppReq) (resp *base.Empty, err error) {
	return
}
