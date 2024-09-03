package main

import (
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	"log"
)

func main() {
	svr := appsvc.NewServer(new(AppServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
