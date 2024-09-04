package main

import (
	"github.com/aiagt/aiagt/app/model/handler"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	"log"
)

func main() {
	svr := modelsvc.NewServer(new(handler.ModelServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
