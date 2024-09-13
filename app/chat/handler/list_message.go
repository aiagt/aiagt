package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/chat/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// ListMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ListMessage(ctx context.Context, req *chatsvc.ListMessageReq) (resp *chatsvc.ListMessageResp, err error) {
	conversation, err := s.conversationDao.GetByID(ctx, req.ConversationId)
	if err != nil {
		return nil, bizListMessage.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, conversation.UserID) {
		return nil, bizListMessage.CodeErr(bizerr.ErrCodeForbidden)
	}

	messages, page, err := s.messageDao.List(ctx, req)
	if err != nil {
		return nil, bizListConversation.NewErr(err)
	}

	resp = &chatsvc.ListMessageResp{
		Messages: mapper.NewGenListMessage(messages),
		Pagination: page,
	}

	return
}
