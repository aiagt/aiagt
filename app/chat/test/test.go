package main

import (
	"context"
	"encoding/json"
	"io"

	"github.com/aiagt/aiagt/pkg/safe"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/chatsvc"
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

	logger(Chat(ctx))
	// logger(ListMessage(ctx))
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

func Chat(ctx context.Context) (any, error) {
	stream, err := rpc.ChatStreamCli.Chat(ctx, &chatsvc.ChatReq{
		AppId:          1,
		ConversationId: safe.Pointer(int64(31)),
		Messages: []*chatsvc.MessageContent{
			{
				Type: chatsvc.MessageType_TEXT,
				Content: &chatsvc.MessageContentValue{
					Text: &chatsvc.MessageContentValueText{
						Text: "乌鲁木齐呢",
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		for _, m := range msg.Messages {
			if m.Content.Type == chatsvc.MessageType_TEXT {
				print(m.Content.Content.Text.Text)
			} else {
				logger(m, nil)
			}
		}
	}

	return nil, nil
}

func ListMessage(ctx context.Context) (any, error) {
	return rpc.ChatCli.ListMessage(ctx, &chatsvc.ListMessageReq{
		ConversationId: 31,
	})
}
