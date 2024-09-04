package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// UpdateConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) UpdateConversation(ctx context.Context, req *chatsvc.UpdateConversationReq) (resp *base.Empty, err error) {
	return
}
