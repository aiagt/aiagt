package main

import (
	"context"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	pluginservice "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	ktclient "github.com/aiagt/kitextool/suite/client"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli, err := pluginservice.NewClient("plugin",
		client.WithHostPorts(":8888"),
		client.WithSuite(ktclient.NewKitexToolSuite(nil, ktclient.WithTransport(transport.TTHeaderFramed))))
	if err != nil {
		panic(err)
	}

	//id := int64(1)
	labels := []int64{1, 2, 3}
	plugin, err := cli.ListPlugin(context.Background(), &pluginsvc.ListPluginReq{Labels: labels})
	if err != nil {
		klog.Warnf("ListPlugin err: %v", err)
	}
	klog.Infof("ListPlugin: %v", plugin)
}
