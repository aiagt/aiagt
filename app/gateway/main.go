package main

import (
	appcontroller "github.com/aiagt/aiagt/app/app/controller"
	chatcontroller "github.com/aiagt/aiagt/app/chat/controller"
	modelcontroller "github.com/aiagt/aiagt/app/model/controller"
	plugincontroller "github.com/aiagt/aiagt/app/plugin/controller"
	usercontroller "github.com/aiagt/aiagt/app/user/controller"
	"github.com/aiagt/aiagt/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts(":8930"))

	r := h.Group("/api/v1")
	usercontroller.RegisterRouter(r, rpc.UserCli)
	modelcontroller.RegisterRouter(r, rpc.ModelCli)
	plugincontroller.RegisterRouter(r, rpc.PluginCli)
	appcontroller.RegisterRouter(r, rpc.AppCli)
	chatcontroller.RegisterRouter(r, rpc.ChatCli, rpc.ChatStreamCli)

	h.Spin()
}
