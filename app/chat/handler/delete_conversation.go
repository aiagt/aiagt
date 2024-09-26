package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteConversation implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) DeleteConversation(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	conversation, err := s.conversationDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteConversation.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, conversation.UserID) {
		return nil, bizDeleteConversation.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.conversationDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteConversation.NewErr(err)
	}

	err = s.messageDao.DeleteByConversationID(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteConversation.NewErr(err)
	}

	return
}
