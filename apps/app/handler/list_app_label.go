package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/app/mapper"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
)

// ListAppLabel implements the AppServiceImpl interface.
func (s *AppServiceImpl) ListAppLabel(ctx context.Context, req *appsvc.ListAppLabelReq) (resp *appsvc.ListAppLabelResp, err error) {
	labels, page, err := s.labelDao.List(ctx, req)
	if err != nil {
		return nil, bizListAppLabel.NewErr(err)
	}

	resp = &appsvc.ListAppLabelResp{
		Labels:     mapper.NewGenListAppLabel(labels),
		Pagination: page,
	}

	return
}
