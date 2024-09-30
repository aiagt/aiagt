package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/chat/mapper"
	"github.com/aiagt/aiagt/common/ctxutil"

	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// ListConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ListConversation(ctx context.Context, req *chatsvc.ListConversationReq) (resp *chatsvc.ListConversationResp, err error) {
	userID := ctxutil.UserID(ctx)

	conversations, pageResp, err := s.conversationDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListConversation.NewErr(err)
	}

	resp = &chatsvc.ListConversationResp{
		Conversations: mapper.NewGenListConversation(conversations),
		Pagination:    pageResp,
	}

	return
}
