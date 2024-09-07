package main

import (
	"github.com/aiagt/aiagt/app/user/conf"
	"github.com/aiagt/aiagt/app/user/dal/db"
	"github.com/aiagt/aiagt/app/user/handler"
	"github.com/aiagt/aiagt/app/user/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"log"
)

func main() {
	handle := handler.NewUserService(db.NewUserDao())

	svr := usersvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
			ktdb.WithDB(ktdb.NewMySQLDial()),
		)),
		server.WithSuite(serversuite.NewServerSuite(handler.NewAuthService(handle))),
	)

	if err := ktdb.DB().AutoMigrate(new(model.User)); err != nil {
		panic(err)
	}

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
