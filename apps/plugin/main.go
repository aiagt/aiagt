package main

import (
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	"log"

	"github.com/aiagt/aiagt/apps/plugin/conf"
	"github.com/aiagt/aiagt/apps/plugin/dal/db"
	"github.com/aiagt/aiagt/apps/plugin/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	"github.com/aiagt/aiagt/rpc"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"gorm.io/gorm"

	"github.com/aiagt/aiagt/apps/plugin/handler"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
)

func main() {
	handle := handler.NewPluginService(db.NewPluginDao(), db.NewLabelDao(), db.NewToolDao(), rpc.UserCli)

	svr := pluginsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
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
