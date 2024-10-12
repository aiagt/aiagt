package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) DeleteMessage(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	message, err := s.messageDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteMessage.NewErr(err)
	}

	conversation, err := s.conversationDao.GetByID(ctx, message.ConversationID)
	if err != nil {
		return nil, bizDeleteMessage.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, conversation.UserID) {
		return nil, bizDeleteMessage.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.messageDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteMessage.NewErr(err)
	}

	return
}
