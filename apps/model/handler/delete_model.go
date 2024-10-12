package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteModel implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) DeleteModel(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	err = s.modelDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteModel.NewErr(err)
	}

	return
}
