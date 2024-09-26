package controller

import (
	"github.com/aiagt/aiagt/common/hertz/router"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterRouter(r *route.RouterGroup, cli appsvc.Client) {
	r = r.Group("/app")

	router.POST(r, "/", cli.CreateApp)
	router.PUT(r, "/:id", cli.UpdateApp)
	router.DELETE(r, "/:id", cli.DeleteApp)
	router.GET(r, "/:id", cli.GetAppByID)
	router.POST(r, "/list", cli.ListApp)

	router.POST(r, "/publish", cli.PublishApp)

	router.GET(r, "/label", cli.ListAppLabel)
}
