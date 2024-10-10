package main

import (
	"log"

	"github.com/aiagt/aiagt/common/logger"
	ktlog "github.com/aiagt/kitextool/option/server/log"

	"github.com/aiagt/aiagt/common/observability"
	"github.com/aiagt/aiagt/pkg/logerr"
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	"gorm.io/plugin/opentelemetry/tracing"

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

	config := conf.Conf()
	observability.InitMetrics(config.Server.Name, config.Metrics.Addr, config.Registry.Address[0])
	observability.InitTracing(config.Server.Name)

	svr := pluginsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			config,
			ktlog.WithLogger(logger.Logger()),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
		)),
		server.WithSuite(serversuite.NewServerSuite(config.GetServerConf(), rpc.UserCli)))

	logerr.Fatal(ktdb.DB().AutoMigrate(new(model.PluginTool), new(model.PluginLabel), new(model.PluginTool)))
	logerr.Fatal(ktdb.DB().Use(tracing.NewPlugin(tracing.WithoutMetrics())))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
