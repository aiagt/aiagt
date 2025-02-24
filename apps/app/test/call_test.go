package test

import (
	"github.com/aiagt/aiagt/common/tests"
	"testing"

	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/rpc"
)

var ctx = tests.InitTesting()

func TestGetApp(t *testing.T) {
	tests.RpcCallWrap(rpc.AppCli.GetAppByID(ctx, &appsvc.GetAppByIDReq{Id: 1}))
}

func TestListApp(t *testing.T) {
	tests.RpcCallWrap(rpc.AppCli.ListApp(ctx, &appsvc.ListAppReq{}))
}

func TestCreateApp(t *testing.T) {
	tests.RpcCallWrap(rpc.AppCli.CreateApp(ctx, &appsvc.CreateAppReq{
		Name:        "test",
		Description: "test app",
		ToolIds:     []int64{1},
	}))
}
