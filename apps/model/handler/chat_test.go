package handler

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/apps/model/model"
	"io"
	"testing"

	"github.com/aiagt/aiagt/pkg/schema"

	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/require"
)

func TestOpenAIChat(t *testing.T) {
	openaiCli := newOpenaiCli(&model.ApiKey{
		APIKey:  "",
		BaseURL: "",
	})

	chatStream, err := openaiCli.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model: "deepseek-chat",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "今天北京天气怎么样",
			},
		},
		Tools: []openai.Tool{
			{
				Type: openai.ToolTypeFunction,
				Function: &openai.FunctionDefinition{
					Name:        "GetWeather",
					Description: "查询今天的天气",
					Strict:      true,
					Parameters: schema.Definition{
						Type:                 "object",
						Required:             []string{"location"},
						AdditionalProperties: false,
						Properties: map[string]schema.Definition{
							"location": {
								Type:        "string",
								Description: "地点名称",
							},
						},
						Items: nil,
					},
				},
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

		if utils.First(r.Choices).Delta.Content != "" {
			fmt.Print(utils.First(r.Choices).Delta.Content)
		} else {
			fmt.Println(utils.Pretty(utils.First(r.Choices).Delta, 1<<16))
		}
	}
}
