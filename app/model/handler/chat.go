package handler

import (
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

func (s *ModelServiceImpl) Chat(req *modelsvc.ChatReq, stream modelsvc.ModelService_ChatServer) (err error) {
	println("Chat called")
	return
}
