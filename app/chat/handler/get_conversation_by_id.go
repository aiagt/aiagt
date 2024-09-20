package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/chat/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// GetConversationByID implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetConversationByID(ctx context.Context, req *base.IDReq) (resp *chatsvc.Conversation, err error) {
	conversation, err := s.conversationDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetConversationByID.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, conversation.UserID) {
		return nil, bizGetConversationByID.CodeErr(bizerr.ErrCodeForbidden)
	}

	resp = mapper.NewGenConversation(conversation)

	return resp, nil
}
