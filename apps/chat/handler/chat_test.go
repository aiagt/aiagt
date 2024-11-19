package handler

import (
	"context"
	"testing"

	"github.com/aiagt/aiagt/apps/chat/conf"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/rpc"
	"github.com/stretchr/testify/require"
)

func init() {
	_ = conf.Conf()
}

func TestChatServiceImpl_buildNewTitle(t *testing.T) {
	svc := NewChatService(nil, nil, nil, nil, nil, rpc.ModelCli, rpc.ModelStreamCli)
	ctx := context.Background()

	ctx, err := login(ctx)
	require.NoError(t, err)

	svc.generateNewTitle(ctx, nil, `https://doc.tryfastgpt.ai/docs/guide/knowledge_base/rag/

我想在参考文献中写这个网页中的内容，参考文献的部分应该怎么写，APA格式`, 0, 1)
}

func login(ctx context.Context) (context.Context, error) {
	password := "au199108"

	resp, err := rpc.UserCli.Login(ctx, &usersvc.LoginReq{Email: "ahao_study@163.com", Password: &password})
	if err != nil {
		return nil, err
	}

	return ctxutil.WithToken(ctx, resp.Token), nil
}
