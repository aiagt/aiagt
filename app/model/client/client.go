package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/sashabaranov/go-openai/jsonschema"

	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"github.com/aiagt/aiagt/rpc"
)

func main() {
	pd := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "The city and state, e.g. San Francisco, CA",
			},
			"unit": {
				Type: jsonschema.String,
				Enum: []string{"celsius", "fahrenheit"},
			},
		},
		Required: []string{"location"},
	}

	p, _ := pd.MarshalJSON()

	stream, err := rpc.ModelStreamCli.Chat(context.Background(), &modelsvc.ChatReq{
		Token:   "",
		ModelId: 0,
		OpenaiReq: &openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []*openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: utils.Pointer("北京天气怎么样？"),
				},
			},
			// MaxTokens:         nil,
			// Temperature:       nil,
			// TopP:              nil,
			// N:                 nil,
			// Stream:            nil,
			// Stop:              nil,
			// PresencePenalty:   nil,
			// ResponseFormat:    nil,
			// Seed:              nil,
			// FrequencyPenalty:  nil,
			// LogitBias:         nil,
			// Logprobs:          nil,
			// TopLogprobs:       nil,
			// User:              nil,
			Functions: []*openai.FunctionDefinition{
				{
					Name:        "get_current_weather",
					Description: utils.Pointer("Get the current weather in a given location"),
					Parameters:  p,
				},
			},
			// FunctionCall:      nil,
			// Tools:             nil,
			// ToolChoice:        nil,
			// StreamOptions:     nil,
			// ParallelToolCalls: nil,
		},
	})
	if err != nil {
		klog.Errorf("err: %#v", err)
		return
	}

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			return
		}

		s, _ := json.Marshal(res)
		klog.Infof("[RECV] %v", string(s))
	}
}
