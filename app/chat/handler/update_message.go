package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// UpdateMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) UpdateMessage(ctx context.Context, req *chatsvc.UpdateMessageReq) (resp *base.Empty, err error) {
	return
}
