package rpc

import (
	"github.com/aiagt/aiagt/common/kitex/clientsuite"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	ktclient "github.com/aiagt/kitextool/suite/client"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

var (
	UserCli   usersvc.Client
	PluginCli pluginsvc.Client
)

func init() {
	var err error

	UserCli, err = usersvc.NewClient("user",
		client.WithHostPorts(":8931"),
		client.WithSuite(clientsuite.NewClientSuite()),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)
	noError(err)

	PluginCli, err = pluginsvc.NewClient("plugin",
		client.WithHostPorts(":8932"),
		client.WithSuite(clientsuite.NewClientSuite()),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))),
	)
	noError(err)
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
