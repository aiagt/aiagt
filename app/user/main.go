package main

import (
	"github.com/aiagt/aiagt/app/user/conf"
	"github.com/aiagt/aiagt/app/user/dal/cache"
	"github.com/aiagt/aiagt/app/user/dal/db"
	"github.com/aiagt/aiagt/app/user/handler"
	"github.com/aiagt/aiagt/app/user/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktrdb "github.com/aiagt/kitextool/option/server/redis"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"gorm.io/gorm"
	"log"
)

func main() {
	handle := handler.NewUserService(db.NewUserDao(), db.NewSecretDao(), cache.NewCaptchaCache())

	svr := usersvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.Conf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
			ktdb.WithDB(ktdb.NewMySQLDial(), ktdb.WithGormConf(&gorm.Config{TranslateError: true})),
			ktrdb.WithRedis(),
		)),
		server.WithSuite(serversuite.NewServerSuite(handler.NewAuthService(handle))),
	)

	if err := ktdb.DB().AutoMigrate(new(model.User), new(model.Secret)); err != nil {
		panic(err)
	}

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
