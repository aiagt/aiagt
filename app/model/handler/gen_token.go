package handler

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"

	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// GenToken implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) GenToken(ctx context.Context, req *modelsvc.GenTokenReq) (resp *modelsvc.GenTokenResp, err error) {
	return nil, kerrors.NewBizStatusError(1, "hello")
}
