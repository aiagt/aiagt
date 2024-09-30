package main

import (
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	"log"

	"github.com/aiagt/aiagt/apps/chat/conf"
	"github.com/aiagt/aiagt/apps/chat/dal/db"
	"github.com/aiagt/aiagt/apps/chat/handler"
	"github.com/aiagt/aiagt/apps/chat/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc/chatservice"
	"github.com/aiagt/aiagt/rpc"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"gorm.io/gorm"
)

func main() {
	handle := handler.NewChatService(db.NewConversationDao(), db.NewMessageDao(), rpc.UserCli, rpc.AppCli, rpc.PluginCli, rpc.ModelCli, rpc.ModelStreamCli)

	svr := chatsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
		)),
		server.WithSuite(serversuite.NewServerSuite(rpc.UserCli)))

	if err := ktdb.DB().AutoMigrate(new(model.Conversation), new(model.Message)); err != nil {
		panic(err)
	}

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
