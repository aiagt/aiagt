package main

import (
	"context"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// ModelServiceImpl implements the last service interface defined in the IDL.
type ModelServiceImpl struct{}

// GenToken implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) GenToken(ctx context.Context, req *modelsvc.GenTokenReq) (resp *modelsvc.GenTokenResp, err error) {
	// TODO: Your code here...
	return
}

func (s *ModelServiceImpl) Chat(req *modelsvc.ChatReq, stream modelsvc.ModelService_ChatServer) (err error) {
	println("Chat called")
	return
}
