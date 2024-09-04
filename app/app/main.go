package main

import (
	"github.com/aiagt/aiagt/app/app/handler"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	"log"
)

func main() {
	svr := appsvc.NewServer(new(handler.AppServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
