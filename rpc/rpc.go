package rpc

import (
	"github.com/aiagt/aiagt/common/confutil"
	"github.com/aiagt/aiagt/common/kitex/clientsuite"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc/chatservice"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc/workflowservice"
	ktconf "github.com/aiagt/kitextool/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/streamclient"
)

var (
	UserCli usersvc.Client

	PluginCli pluginsvc.Client

	AppCli appsvc.Client

	ChatCli       chatsvc.Client
	ChatStreamCli chatsvc.StreamClient

	ModelCli       modelsvc.Client
	ModelStreamCli modelsvc.StreamClient

	WorkflowCli workflowsvc.Client

	conf = new(ktconf.MultiClientConf)
)

func Conf() *ktconf.MultiClientConf {
	return conf
}

func init() {
	confutil.LoadConf(conf, ".", "rpc")
	confutil.LoadConf(conf, "../../../rpc")

	UserCli = usersvc.MustNewClient("user", client.WithSuite(clientsuite.NewClientSuite(conf, "user")))

	PluginCli = pluginsvc.MustNewClient("plugin", client.WithSuite(clientsuite.NewClientSuite(conf, "plugin")))

	AppCli = appsvc.MustNewClient("app", client.WithSuite(clientsuite.NewClientSuite(conf, "app")))

	ChatCli = chatsvc.MustNewClient("chat", client.WithSuite(clientsuite.NewClientSuite(conf, "chat")))
	ChatStreamCli = chatsvc.MustNewStreamClient("chat", streamclient.WithSuite(clientsuite.NewClientSuite(conf, "chat")))

	ModelCli = modelsvc.MustNewClient("model", client.WithSuite(clientsuite.NewClientSuite(conf, "model")))
	ModelStreamCli = modelsvc.MustNewStreamClient("model", streamclient.WithSuite(clientsuite.NewClientSuite(conf, "model")))

	WorkflowCli = workflowsvc.MustNewClient("workflow", client.WithSuite(clientsuite.NewClientSuite(conf, "workflow")))
}
