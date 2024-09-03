package main

import (
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	"log"
)

func main() {
	svr := modelsvc.NewServer(new(ModelServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
