package handler

import (
	"github.com/aiagt/aiagt/app/chat/dal/db"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct {
	conversationDao *db.ConversationDao
	messageDao      *db.MessageDao

	appCli         appsvc.Client
	modelCli       modelsvc.Client
	modelStreamCli modelsvc.StreamClient
}

func NewChatService(conversationDao *db.ConversationDao, messageDao *db.MessageDao, appCli appsvc.Client, modelCli modelsvc.Client, modelStreamCli modelsvc.StreamClient) *ChatServiceImpl {
	return &ChatServiceImpl{conversationDao: conversationDao, messageDao: messageDao, appCli: appCli, modelCli: modelCli, modelStreamCli: modelStreamCli}
}
