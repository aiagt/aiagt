package main

import (
	"context"

	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	ctx := context.Background()

	// logger(SendCaptcha(ctx))
	// logger(Register(ctx))
	logger(Login(ctx))
}

func logger(resp any, err error) {
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Infof("resp: %#v", resp)
}

func SendCaptcha(ctx context.Context) (any, error) {
	return rpc.UserCli.SendCaptcha(ctx, &usersvc.SendCaptchaReq{
		Email: "ahao_study@163.com",
		Type:  usersvc.CaptchaType_AUTH,
	})
}

func Register(ctx context.Context) (any, error) {
	return rpc.UserCli.Register(ctx, &usersvc.RegisterReq{
		Email:    "ahao_study@163.com",
		Captcha:  "14819",
		Username: "ahaostudy",
		Password: "au199108",
	})
}

func Login(ctx context.Context) (any, error) {
	password := "au199108"
	return rpc.UserCli.Login(ctx, &usersvc.LoginReq{
		Email:    "ahao_study@163.com",
		Password: &password,
	})
}
