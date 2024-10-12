package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/chat/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// UpdateConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) UpdateConversation(ctx context.Context, req *chatsvc.UpdateConversationReq) (resp *base.Empty, err error) {
	conversation, err := s.conversationDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizUpdateConversation.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, conversation.UserID) {
		return nil, bizUpdateConversation.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.conversationDao.Update(ctx, req.Id, mapper.NewModelUpdateConversation(req))
	if err != nil {
		return nil, bizUpdateConversation.NewErr(err)
	}

	return
}
