package main

import (
	"github.com/aiagt/aiagt/apps/workflow/conf"
	"github.com/aiagt/aiagt/apps/workflow/dal/db"
	"github.com/aiagt/aiagt/apps/workflow/handler"
	"github.com/aiagt/aiagt/apps/workflow/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	"github.com/aiagt/aiagt/common/logger"
	"github.com/aiagt/aiagt/common/observability"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc/workflowservice"
	"github.com/aiagt/aiagt/pkg/logerr"
	"github.com/aiagt/aiagt/rpc"
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktlog "github.com/aiagt/kitextool/option/server/log"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"log"
)

func main() {
	handle := handler.NewWorkflowServiceImpl(
		db.NewWorkflowDao(),
		db.NewWorkflowNodeDao(),
		rpc.ModelCli,
		rpc.PluginCli,
		rpc.UserCli,
	)

	config := conf.Conf()
	observability.InitMetrics(config.Server.Name, config.Metrics.Addr, config.Registry.Address[0])
	observability.InitTracing(config.Server.Name, config.Tracing.ExportAddr)

	svr := workflowsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			config,
			ktlog.WithLogger(logger.Logger()),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
		)),
		server.WithSuite(serversuite.NewServerSuite(config.GetServerConf(), rpc.UserCli)))

	logerr.Fatal(ktdb.DB().AutoMigrate(new(model.Workflow), new(model.WorkflowNode)))
	logerr.Fatal(ktdb.DB().Use(tracing.NewPlugin(tracing.WithoutMetrics())))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
