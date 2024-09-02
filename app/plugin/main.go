package main

import (
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	"log"
)

func main() {
	svr := pluginsvc.NewServer(new(PluginServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
