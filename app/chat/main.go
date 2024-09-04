package main

import (
	"github.com/aiagt/aiagt/app/chat/handler"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc/chatservice"
	"log"
)

func main() {
	svr := chatsvc.NewServer(new(handler.ChatServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
