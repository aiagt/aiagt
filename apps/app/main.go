package main

import (
	"github.com/aiagt/aiagt/common/logger"
	ktlog "github.com/aiagt/kitextool/option/server/log"
	"log"

	"github.com/aiagt/aiagt/common/observability"
	"github.com/aiagt/aiagt/pkg/logerr"
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	"gorm.io/plugin/opentelemetry/tracing"

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

	config := conf.Conf()
	observability.InitMetrics(config.Server.Name, config.Metrics.Addr, config.Registry.Address[0])
	observability.InitTracing(config.Server.Name)

	svr := appsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			config,
			ktlog.WithLogger(logger.Logger()),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
		)),
		server.WithSuite(serversuite.NewServerSuite(config.GetServerConf(), rpc.UserCli)),
	)

	logerr.Fatal(ktdb.DB().AutoMigrate(new(model.App), new(model.AppLabel)))
	logerr.Fatal(ktdb.DB().Use(tracing.NewPlugin(tracing.WithoutMetrics())))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
