package main

import (
	"github.com/aiagt/aiagt/app/user/handler"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	"log"
)

func main() {
	svr := usersvc.NewServer(new(handler.UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
