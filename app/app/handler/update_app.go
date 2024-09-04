package handler

import (
	"context"

	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// UpdateApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) UpdateApp(ctx context.Context, req *appsvc.UpdateAppReq) (resp *base.Empty, err error) {
	return
}
