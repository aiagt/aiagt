package main

import (
	ktrdb "github.com/aiagt/kitextool/option/server/redis"
	"log"

	"github.com/aiagt/aiagt/app/model/conf"
	"github.com/aiagt/aiagt/app/model/dal/cache"
	"github.com/aiagt/aiagt/app/model/handler"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	"github.com/aiagt/aiagt/rpc"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	handle := handler.NewModelService(cache.NewCallTokenCache())

	svr := modelsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
			ktrdb.WithRedis(),
		)),
		server.WithSuite(serversuite.NewServerSuite(rpc.UserCli)))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
