package rpc

import (
	"github.com/aiagt/aiagt/common/kitex/clientsuite"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc/appservice"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc/chatservice"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	ktclient "github.com/aiagt/kitextool/suite/client"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/streamclient"
	"github.com/cloudwego/kitex/transport"
)

var (
	UserCli        usersvc.Client
	PluginCli      pluginsvc.Client
	AppCli         appsvc.Client
	ChatCli        chatsvc.Client
	ChatStreamCli  chatsvc.StreamClient
	ModelCli       modelsvc.Client
	ModelStreamCli modelsvc.StreamClient
)

func init() {
	UserCli = usersvc.MustNewClient("user",
		client.WithHostPorts(":8931"),
		client.WithSuite(clientsuite.NewClientSuite()),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)

	PluginCli = pluginsvc.MustNewClient("plugin",
		client.WithHostPorts(":8932"),
		client.WithSuite(clientsuite.NewClientSuite()),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)

	AppCli = appsvc.MustNewClient("app",
		client.WithHostPorts(":8933"),
		client.WithSuite(clientsuite.NewClientSuite()),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)

	ChatCli = chatsvc.MustNewClient("app",
		client.WithHostPorts(":8934"),
		client.WithSuite(clientsuite.NewClientSuite()),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)

	ChatStreamCli = chatsvc.MustNewStreamClient("app",
		streamclient.WithHostPorts(":8934"),
		streamclient.WithSuite(clientsuite.NewClientSuite()),
		streamclient.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)

	ModelCli = modelsvc.MustNewClient("model",
		client.WithHostPorts(":8935"),
		client.WithSuite(clientsuite.NewClientSuite()),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)

	ModelStreamCli = modelsvc.MustNewStreamClient("model",
		streamclient.WithHostPorts(":8935"),
		streamclient.WithSuite(clientsuite.NewClientSuite()),
		streamclient.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
