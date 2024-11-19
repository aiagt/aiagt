package controller

import (
	"github.com/aiagt/aiagt/common/hertz/router"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc/chatservice"
	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterRouter(r *route.RouterGroup, cli chatsvc.Client, streamCli chatsvc.StreamClient) {
	chatRouter := r.Group("/chat")

	messageRouter := chatRouter.Group("/message")
	router.GET(messageRouter, "/", cli.ListMessage)
	router.PUT(messageRouter, "/:id", cli.UpdateMessage)
	router.DELETE(messageRouter, "/:id", cli.DeleteMessage)

	conversationRouter := chatRouter.Group("/conversation")
	router.GET(conversationRouter, "/", cli.ListConversation)
	router.GET(conversationRouter, "/:id", cli.GetConversationByID)
	router.PUT(conversationRouter, "/:id", cli.UpdateConversation)
	router.DELETE(conversationRouter, "/:id", cli.DeleteConversation)
	router.POST(conversationRouter, "/develop", cli.InitDevelop)

	router.SSE(chatRouter, "/chat", streamCli.Chat)
}
