package main

import (
	"context"
	"log"

	"github.com/aiagt/aiagt/app/model/conf"
	"github.com/aiagt/aiagt/app/model/handler"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	ktserver "github.com/aiagt/kitextool/suite/server"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
)

type suite struct{}

func (s *suite) Options() []server.Option {
	opts := []server.Option{
		server.WithMiddleware(func(next endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, req, resp interface{}) (err error) {
				return kerrors.NewBizStatusError(1, "hello")
			}
		}),
	}
	return opts
}

func main() {
	svr := modelsvc.NewServer(new(handler.ModelServiceImpl),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithSuite(ktserver.NewKitexToolSuite(conf.Conf())),
		server.WithSuite(new(suite)),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
