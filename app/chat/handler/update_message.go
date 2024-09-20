package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/chat/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// UpdateMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) UpdateMessage(ctx context.Context, req *chatsvc.UpdateMessageReq) (resp *base.Empty, err error) {
	message, err := s.messageDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizUpdateMessage.NewErr(err)
	}

	conversation, err := s.conversationDao.GetByID(ctx, message.ConversationID)
	if err != nil {
		return nil, bizUpdateMessage.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, conversation.UserID) {
		return nil, bizUpdateMessage.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.messageDao.Update(ctx, req.Id, mapper.NewModelUpdateMessage(req))

	return
}
