package handler

import (
	"context"

	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// ListConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ListConversation(ctx context.Context, req *chatsvc.ListConversationReq) (resp *chatsvc.ListConversationResp, err error) {
	return
}
