package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/chat/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/kitex_gen/appsvc"

	"github.com/aiagt/aiagt/apps/chat/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
)

// InitDevelop implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) InitDevelop(ctx context.Context, req *chatsvc.InitDevelopReq) (resp *chatsvc.InitDevelopResp, err error) {
	userID := ctxutil.UserID(ctx)

	getAppResp, err := s.appCli.GetAppByID(ctx, &appsvc.GetAppByIDReq{Id: req.AppId})
	if err != nil {
		return nil, bizInitDevelop.CallErr(err).Log(ctx, "get app by id error")
	}

	if getAppResp.App.AuthorId != userID {
		return nil, bizInitDevelop.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "user is not the app author")
	}

	var messages []*model.Message

	conversation, err := s.conversationDao.GetDevelop(ctx, userID, req.AppId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		const newConversationTitle = "Develop Conversation"
		conversation = &model.Conversation{
			Title:   newConversationTitle,
			UserID:  userID,
			AppID:   req.AppId,
			Develop: true,
		}

		err = s.conversationDao.Create(ctx, conversation)
		if err != nil {
			return nil, bizInitDevelop.NewErr(err).Log(ctx, "create conversation error")
		}
	} else if err != nil {
		return nil, bizInitDevelop.NewErr(err).Log(ctx, "get develop conversation error")
	} else {
		messages, err = s.messageDao.GetByConversationID(ctx, conversation.ID)
		if err != nil {
			return nil, bizInitDevelop.NewErr(err).Log(ctx, "get message by conversation id error")
		}
	}

	resp = &chatsvc.InitDevelopResp{
		Conversation: mapper.NewGenConversation(conversation),
		Messages:     mapper.NewGenListMessage(messages),
	}

	return
}
