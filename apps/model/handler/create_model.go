package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/model/mapper"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// CreateModel implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) CreateModel(ctx context.Context, req *modelsvc.CreateModelReq) (resp *base.Empty, err error) {
	model := mapper.NewModelCreateModel(req)

	err = s.modelDao.Create(ctx, model)
	if err != nil {
		return nil, bizCreateModel.NewErr(err)
	}

	return
}
