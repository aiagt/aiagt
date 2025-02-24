package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/model/mapper"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// GetModelByID implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) GetModelByID(ctx context.Context, req *base.IDReq) (resp *modelsvc.Model, err error) {
	model, err := s.modelDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetModelById.NewErr(err)
	}

	resp = mapper.NewGenModel(model)

	return
}
