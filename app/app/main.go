package main

import (
	"log"

	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
)

func main() {
	svr := appsvc.NewServer(new(AppServiceImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
