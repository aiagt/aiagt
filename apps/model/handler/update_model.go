package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/model/mapper"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// UpdateModel implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) UpdateModel(ctx context.Context, req *modelsvc.UpdateModelReq) (resp *base.Empty, err error) {
	model := mapper.NewModelUpdateModel(req)

	err = s.modelDao.Update(ctx, req.Id, model)
	if err != nil {
		return nil, bizUpdateModel.NewErr(err)
	}

	return
}
