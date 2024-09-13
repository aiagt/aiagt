package main

import (
	"github.com/aiagt/aiagt/app/model/conf"
	"github.com/aiagt/aiagt/app/model/handler"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	"github.com/aiagt/aiagt/rpc"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"log"
)

func main() {
	handle := handler.NewModelService()

	svr := modelsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
		)),
		server.WithSuite(serversuite.NewServerSuite(rpc.UserCli)))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
