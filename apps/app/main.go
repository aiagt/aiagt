package main

import (
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	"log"

	"github.com/aiagt/aiagt/apps/app/conf"
	"github.com/aiagt/aiagt/apps/app/dal/db"
	"github.com/aiagt/aiagt/apps/app/handler"
	"github.com/aiagt/aiagt/apps/app/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	"github.com/aiagt/aiagt/rpc"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"gorm.io/gorm"
)

func main() {
	handle := handler.NewAppService(db.NewAppDao(), db.NewLabelDao(), rpc.UserCli, rpc.PluginCli)

	svr := appsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
		)),
		server.WithSuite(serversuite.NewServerSuite(rpc.UserCli)),
	)

	if err := ktdb.DB().AutoMigrate(new(model.App), new(model.AppLabel)); err != nil {
		panic(err)
	}

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
