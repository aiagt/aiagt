package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/model/mapper"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// ListModel implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) ListModel(ctx context.Context, req *modelsvc.ListModelReq) (resp *modelsvc.ListModelResp, err error) {
	models, page, err := s.modelDao.List(ctx, req)
	if err != nil {
		return nil, bizListModel.NewErr(err)
	}

	resp = &modelsvc.ListModelResp{
		Models:     mapper.NewGenListModel(models),
		Pagination: page,
	}

	return
}
