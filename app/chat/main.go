package main

import (
	"github.com/aiagt/aiagt/app/chat/dal/db"
	"github.com/aiagt/aiagt/app/chat/handler"
	"github.com/aiagt/aiagt/app/chat/model"
	"github.com/aiagt/aiagt/app/user/conf"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc/chatservice"
	"github.com/aiagt/aiagt/rpc"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"gorm.io/gorm"
	"log"
)

func main() {
	handle := handler.NewChatService(db.NewConversationDao(), db.NewMessageDao(), rpc.AppCli, rpc.ModelCli, rpc.ModelStreamCli)

	svr := chatsvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
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
