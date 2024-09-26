package main

import (
	"log"

	"github.com/cloudwego/kitex/transport"
	"gorm.io/gorm"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktrdb "github.com/aiagt/kitextool/option/server/redis"

	"github.com/aiagt/aiagt/app/model/conf"
	"github.com/aiagt/aiagt/app/model/dal/cache"
	"github.com/aiagt/aiagt/app/model/dal/db"
	"github.com/aiagt/aiagt/app/model/handler"
	"github.com/aiagt/aiagt/app/model/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	"github.com/aiagt/aiagt/rpc"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
)

func main() {
	handle := handler.NewModelService(db.NewModelDao(), cache.NewCallTokenCache())

	svr := modelsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
			ktrdb.WithRedis(),
		)),
		server.WithSuite(serversuite.NewServerSuite(rpc.UserCli)))

	if err := ktdb.DB().AutoMigrate(new(model.Models)); err != nil {
		panic(err)
	}

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
