package main

import (
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc/chatservice"
	"log"
)

func main() {
	svr := chatsvc.NewServer(new(ChatServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
