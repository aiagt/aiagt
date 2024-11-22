package handler

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/aiagt/aiagt/apps/model/conf"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/require"
)

func TestOpenAIChat(t *testing.T) {
	conf.Conf().OpenAI = conf.OpenAI{
		APIKey:  "sk-yjyY92xs9ivKwPd4B549C9B80eF74809BfF6887159944321",
		BaseURL: "https://api.lqqq.cc/v1",
	}

	config := openai.DefaultConfig(conf.Conf().OpenAI.APIKey)
	config.BaseURL = conf.Conf().OpenAI.BaseURL

	openaiCli := openai.NewClientWithConfig(config)

	chatStream, err := openaiCli.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo-0125",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "你好，吃饭了吗",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "你好，我是一个人工智能助手，不需要进食。您需要什么帮助吗？",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "今天开心吗",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "作为人工智能助手，我没有情绪，但我很高兴能为您提供帮助。您有什么问题或需求可以告诉我吗？",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "好的咧",
			},
		},
	})

	require.NoError(t, err)

	for {
		r, err := chatStream.Recv()
		if errors.Is(err, io.EOF) {
			t.Log("\nDONE")
			return
		}

		require.NoError(t, err)

		fmt.Print(utils.First(r.Choices).Delta.Content)
	}
}
