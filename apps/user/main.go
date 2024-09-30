package main

import (
	ktcenter "github.com/aiagt/kitextool/conf/center"
	ktlog "github.com/aiagt/kitextool/option/server/log"
	ktregistry "github.com/aiagt/kitextool/option/server/registry"
	"log"

	"github.com/aiagt/aiagt/apps/user/conf"
	"github.com/aiagt/aiagt/apps/user/dal/cache"
	"github.com/aiagt/aiagt/apps/user/dal/db"
	"github.com/aiagt/aiagt/apps/user/handler"
	"github.com/aiagt/aiagt/apps/user/model"
	"github.com/aiagt/aiagt/common/kitex/serversuite"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	ktrdb "github.com/aiagt/kitextool/option/server/redis"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/server"
	"gorm.io/gorm"
)

func main() {
	handle := handler.NewUserService(db.NewUserDao(), db.NewSecretDao(), cache.NewCaptchaCache())

	svr := usersvc.NewServer(handle,
		server.WithSuite(ktserver.NewKitexToolEmptySuite(
			conf.Conf(),
			ktserver.WithLogger(ktlog.WithZap),
			ktserver.WithDynamicConfig(ktcenter.WithConsulConfigCenter(nil)),
			ktregistry.WithRegistry(ktregistry.NewConsulRegistry()),
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
