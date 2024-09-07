package main

import (
	"context"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	ctx := context.Background()

	getPlugin(ctx)
	//updatePlugin(ctx)
}

func listPlugin(ctx context.Context) {
	labels := []int64{1, 2, 3}
	plugin, err := rpc.PluginCli.ListPlugin(ctx, &pluginsvc.ListPluginReq{Labels: labels})
	if err != nil {
		klog.Warnf("ListPlugin err: %v", err)
	}
	klog.Infof("ListPlugin: %v", plugin)
}

func createPlugin(ctx context.Context) {
	resp, err := rpc.PluginCli.CreatePlugin(ctx, &pluginsvc.CreatePluginReq{
		Name:          "plugin1",
		Description:   "plugin test",
		DescriptionMd: "#Plugin1\nplugin test",
		IsPrivate:     true,
		HomePage:      "https://github.com/aiagt/aiagt",
		EnableSecret:  false,
		Secrets: []*pluginsvc.PluginSecret{
			{
				Name:          "secret1",
				Description:   "secret test",
				AcquireMethod: "from github settings",
				Link:          "https://github.com",
			},
		},
		LabelTexts: []string{"label1", "label2"},
		Logo:       "https://github.com",
	})
	if err != nil {
		klog.Warnf("create plugin err: %v", err)
	}
	klog.Infof("CreatePluginResp: %+v", resp)
}

func getPlugin(ctx context.Context) {
	plugin, err := rpc.PluginCli.GetPluginByID(ctx, &base.IDReq{Id: 1})
	if err != nil {
		klog.Warnf("GetPlugin err: %v", err)
	}
	klog.Infof("GetPlugin: %#v", plugin)
}

func updatePlugin(ctx context.Context) {
	resp, err := rpc.PluginCli.UpdatePlugin(ctx, &pluginsvc.UpdatePluginReq{
		Id:            1,
		Name:          "plugin1",
		Description:   "plugin test",
		DescriptionMd: "#Plugin1\nplugin test",
		IsPrivate:     true,
		HomePage:      "https://github.com/aiagt/aiagt",
		EnableSecret:  false,
		Secrets: []*pluginsvc.PluginSecret{
			{
				Name:          "secret1",
				Description:   "secret test",
				AcquireMethod: "from github settings",
				Link:          "https://github.com",
			},
		},
		LabelIds:   []int64{3, 5},
		LabelTexts: []string{"image"},
		Logo:       "https://github.com",
	})
	if err != nil {
		klog.Warnf("update plugin err: %v", err)
	}
	klog.Infof("UpdatePluginResp: %+v", resp)

}
