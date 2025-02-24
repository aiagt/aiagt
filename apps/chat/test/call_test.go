package test

import (
	"github.com/aiagt/aiagt/common/tests"
	"github.com/stretchr/testify/require"
	"io"
	"testing"

	"github.com/aiagt/aiagt/pkg/utils"

	"github.com/aiagt/aiagt/kitex_gen/chatsvc"
	"github.com/aiagt/aiagt/rpc"
)

var ctx = tests.InitTesting()

func TestChat(t *testing.T) {
	stream, err := rpc.ChatStreamCli.Chat(ctx, &chatsvc.ChatReq{
		AppId:          1,
		ConversationId: utils.PtrOf(int64(31)),
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
	require.NoError(t, err)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		require.NoError(t, err)

		for _, m := range msg.Messages {
			if m.Content.Type == chatsvc.MessageType_TEXT {
				print(m.Content.Content.Text.Text)
			} else {
				tests.Log(m)
			}
		}
	}
}

func TestListMessage(t *testing.T) {
	tests.RpcCallWrap(rpc.ChatCli.ListMessage(ctx, &chatsvc.ListMessageReq{
		ConversationId: 31,
	}))
}
