package test

import (
	"github.com/aiagt/aiagt/common/tests"
	"github.com/aiagt/aiagt/pkg/utils"
	"testing"

	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/rpc"
)

var ctx = tests.InitTesting()

func TestLogin(t *testing.T) {
	tests.RpcCallWrap(rpc.UserCli.Login(ctx, &usersvc.LoginReq{
		Email:    "ahao_study@163.com",
		Password: utils.PtrOf("123456"),
	}))
}

func TestSendCaptcha(t *testing.T) {
	tests.RpcCallWrap(rpc.UserCli.SendCaptcha(ctx, &usersvc.SendCaptchaReq{
		Email: "ahao_study@163.com",
		Type:  usersvc.CaptchaType_AUTH,
	}))
}

func TestRegister(t *testing.T) {
	tests.RpcCallWrap(rpc.UserCli.Register(ctx, &usersvc.RegisterReq{
		Email:   "ahao_study@163.com",
		Captcha: "14819",
	}))
}
