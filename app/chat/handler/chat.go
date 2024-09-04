package handler

import (
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

func (s *ChatServiceImpl) Chat(req *chatsvc.ChatReq, stream chatsvc.ChatService_ChatServer) (err error) {
	println("Chat called")
	return
}
