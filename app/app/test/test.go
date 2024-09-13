package main

import (
	"context"
	"encoding/json"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	ctx := context.Background()
	ctx, err := login(ctx)
	if err != nil {
		logger(nil, err)
	}

	logger(GetApp(ctx))
	//logger(ListApp(ctx))
	//logger(CreateApp(ctx))
}

func login(ctx context.Context) (context.Context, error) {
	password := "au199108"
	resp, err := rpc.UserCli.Login(ctx, &usersvc.LoginReq{Email: "ahao_study@163.com", Password: &password})
	if err != nil {
		return nil, err
	}
	return ctxutil.WithToken(ctx, resp.Token), nil
}

func logger(resp any, err error) {
	if err != nil {
		klog.Error(err)
		return
	}
	result, _ := json.MarshalIndent(resp, "", "  ")
	klog.Infof("result: %v", string(result))
}

func GetApp(ctx context.Context) (any, error) {
	return rpc.AppCli.GetAppByID(ctx, &base.IDReq{Id: 1})
}

func ListApp(ctx context.Context) (any, error) {
	return rpc.AppCli.ListApp(ctx, &appsvc.ListAppReq{})
}

func CreateApp(ctx context.Context) (any, error) {
	return rpc.AppCli.CreateApp(ctx, &appsvc.CreateAppReq{
		Name:        "test",
		Description: "test app",
		ToolIds:     []int64{1},
	})
}
