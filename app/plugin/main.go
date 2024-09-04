package main

import (
	"github.com/aiagt/aiagt/app/plugin/handler"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	"log"
)

func main() {
	svr := pluginsvc.NewServer(new(handler.PluginServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
