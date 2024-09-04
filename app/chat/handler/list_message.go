package handler

import (
	"context"

	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// ListMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ListMessage(ctx context.Context, req *chatsvc.ListMessageReq) (resp *chatsvc.ListMessageResp, err error) {
	return
}
