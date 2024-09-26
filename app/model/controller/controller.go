package controller

import (
	"github.com/aiagt/aiagt/common/hertz/router"
	"github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterRouter(r *route.RouterGroup, handler modelservice.Client) {
	r = r.Group("/model")

	router.GET(r, "/", handler.ListModel)
	router.GET(r, "/:id", handler.GetModelByID)
	router.POST(r, "/", handler.CreateModel)
	router.PUT(r, "/:id", handler.UpdateModel)
	router.DELETE(r, "/:id", handler.DeleteModel)
}
