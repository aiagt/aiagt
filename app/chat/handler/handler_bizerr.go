package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "chat"

	bizCodeChat                = 0
	bizCodeDeleteConversation  = 1
	bizCodeDeleteMessage       = 2
	bizCodeGetConversationByID = 3
	bizCodeListConversation    = 4
	bizCodeListMessage         = 5
	bizCodeUpdateConversation  = 6
	bizCodeUpdateMessage       = 7
)

var (
	bizChat                *bizerr.Biz
	bizDeleteConversation  *bizerr.Biz
	bizDeleteMessage       *bizerr.Biz
	bizGetConversationByID *bizerr.Biz
	bizListConversation    *bizerr.Biz
	bizListMessage         *bizerr.Biz
	bizUpdateConversation  *bizerr.Biz
	bizUpdateMessage       *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizChat = bizerr.NewBiz(ServiceName, "Chat", baseCode+bizCodeChat)
	bizDeleteConversation = bizerr.NewBiz(ServiceName, "DeleteConversation", baseCode+bizCodeDeleteConversation)
	bizDeleteMessage = bizerr.NewBiz(ServiceName, "DeleteMessage", baseCode+bizCodeDeleteMessage)
	bizGetConversationByID = bizerr.NewBiz(ServiceName, "GetConversationByID", baseCode+bizCodeGetConversationByID)
	bizListConversation = bizerr.NewBiz(ServiceName, "ListConversation", baseCode+bizCodeListConversation)
	bizListMessage = bizerr.NewBiz(ServiceName, "ListMessage", baseCode+bizCodeListMessage)
	bizUpdateConversation = bizerr.NewBiz(ServiceName, "UpdateConversation", baseCode+bizCodeUpdateConversation)
	bizUpdateMessage = bizerr.NewBiz(ServiceName, "UpdateMessage", baseCode+bizCodeUpdateMessage)
}
