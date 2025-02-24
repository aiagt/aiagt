package tests

import (
	"context"
	"github.com/aiagt/aiagt/common/confutil"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/aiagt/aiagt/rpc"
	ktconf "github.com/aiagt/kitextool/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	"regexp"
)

func InitTesting() (ctx context.Context) {
	ctx = context.Background()

	var (
		password = "au199108"
		email    = "ahao_study@163.com"
	)

	ctx, err := login(ctx, email, password)
	if err != nil {
		log.Fatal(err)
	}

	return ctx
}

func RpcCallWrap(resp any, err error) {
	if err != nil {
		klog.Error(err)
		return
	}

	Log(resp)
}

func Log(val any) {
	var m map[string]interface{}
	_ = mapstructure.Decode(val, &m)

	for k, v := range m {
		if bs, ok := v.([]byte); ok {
			m[k] = string(bs)
		}
	}

	klog.Infof("%v", utils.Pretty(m, 0))
}

func LoadConf(conf ktconf.Conf) {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`^(.*?/apps/[^/]+/)[^/]+(/?)$`)
	confDir := re.ReplaceAllString(workDir, "${1}conf")

	confutil.LoadConf(conf, confDir)
}

func login(ctx context.Context, email, password string) (context.Context, error) {
	resp, err := rpc.UserCli.Login(ctx, &usersvc.LoginReq{Email: email, Password: &password})
	if err != nil {
		return nil, err
	}

	return ctxutil.WithToken(ctx, resp.Token), nil
}
