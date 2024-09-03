package main

import (
	"context"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

func (s *ChatServiceImpl) Chat(req *chatsvc.ChatReq, stream chatsvc.ChatService_ChatServer) (err error) {
	println("Chat called")
	return
}

// UpdateMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) UpdateMessage(ctx context.Context, req *chatsvc.UpdateMessageReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// DeleteMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) DeleteMessage(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// ListMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ListMessage(ctx context.Context, req *chatsvc.ListMessageReq) (resp *chatsvc.ListMessageResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) UpdateConversation(ctx context.Context, req *chatsvc.UpdateConversationReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// DeleteConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) DeleteConversation(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// GetConversationByID implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetConversationByID(ctx context.Context, req *base.IDReq) (resp *chatsvc.Conversation, err error) {
	// TODO: Your code here...
	return
}

// ListConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ListConversation(ctx context.Context, req *chatsvc.ListConversationReq) (resp *chatsvc.ListConversationResp, err error) {
	// TODO: Your code here...
	return
}
