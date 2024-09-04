package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// GetConversationByID implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetConversationByID(ctx context.Context, req *base.IDReq) (resp *chatsvc.Conversation, err error) {
	return
}
