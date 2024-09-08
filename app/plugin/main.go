package main

import (
	"github.com/aiagt/aiagt/app/plugin/conf"
	"github.com/aiagt/aiagt/app/plugin/dal/db"
	"github.com/aiagt/aiagt/app/plugin/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	"github.com/aiagt/aiagt/rpc"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"gorm.io/gorm"
	"log"

	"github.com/aiagt/aiagt/app/plugin/handler"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
)

func main() {
	handle := handler.NewPluginService(db.NewPluginDao(), db.NewLabelDao(), db.NewToolDao(), rpc.UserCli)

	svr := pluginsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
		)),
		server.WithSuite(serversuite.NewServerSuite(rpc.UserCli)))

	if err := ktdb.DB().AutoMigrate(new(model.Plugin), new(model.PluginLabel), new(model.PluginTool)); err != nil {
		panic(err)
	}

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
