package handler

import (
	"context"

	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// GenToken implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) GenToken(ctx context.Context, req *modelsvc.GenTokenReq) (resp *modelsvc.GenTokenResp, err error) {
	resp = &modelsvc.GenTokenResp{}
	return
}
