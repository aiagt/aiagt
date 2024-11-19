package controller

import (
	"github.com/aiagt/aiagt/common/hertz/router"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterRouter(r *route.RouterGroup, cli pluginsvc.Client) {
	r = r.Group("/plugin")

	router.POST(r, "/", cli.CreatePlugin)
	router.PUT(r, "/:id", cli.UpdatePlugin)
	router.DELETE(r, "/:id", cli.DeletePlugin)
	router.GET(r, "/", cli.GetPluginByKey)
	router.GET(r, "/:id", cli.GetPluginByID)
	router.POST(r, "/list", cli.ListPlugin)
	router.POST(r, "/list_by_tools", cli.ListPluginByTools)

	router.POST(r, "/publish/:id", cli.PublishPlugin)

	router.POST(r, "/tool", cli.CreateTool)
	router.PUT(r, "/tool/:id", cli.UpdateTool)
	router.DELETE(r, "/tool/:id", cli.DeleteTool)
	router.GET(r, "/tool/:id", cli.GetToolByID)
	router.POST(r, "/tool/list", cli.ListPluginTool)

	router.GET(r, "/label", cli.ListPluginLabel)
	// router.POST(r, "/tool/call/:id", cli.CallPluginTool)
	router.POST(r, "/tool/test", cli.TestPluginTool)
}
