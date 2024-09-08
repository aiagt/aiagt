package main

import (
	"context"
	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	modelservice "github.com/aiagt/aiagt/kitex_gen/modelsvc/modelservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli := modelservice.MustNewClient("model",
		client.WithHostPorts(":8889"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler))

	resp, err := cli.GenToken(context.Background(), &modelsvc.GenTokenReq{})
	if err != nil {
		bizErr, ok := kerrors.FromBizStatusError(err)
		if ok {
			klog.Errorf("gen token biz err: %+v", bizErr)
		} else {
			klog.Errorf("gen token err: %+v", err)
		}
	}
	klog.Info(resp)
}
