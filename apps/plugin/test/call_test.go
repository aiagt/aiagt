package test

import (
	"encoding/json"
	"github.com/aiagt/aiagt/common/tests"
	"github.com/aiagt/aiagt/pkg/caller"
	"github.com/aiagt/aiagt/pkg/schema"
	"testing"

	"github.com/sashabaranov/go-openai/jsonschema"

	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/rpc"
)

var ctx = tests.InitTesting()

func TestGetPlugin(t *testing.T) {
	tests.RpcCallWrap(rpc.PluginCli.GetPluginByID(ctx, &base.IDReq{Id: 1}))
}

func TestCreatePlugin(t *testing.T) {
	tests.RpcCallWrap(rpc.PluginCli.CreatePlugin(ctx, &pluginsvc.CreatePluginReq{
		Name:          "plugin1",
		Description:   "plugin test",
		DescriptionMd: "#Plugin1\nplugin test",
		IsPrivate:     true,
		HomePage:      "https://github.com/aiagt/aiagt",
		EnableSecret:  false,
		Secrets: []*pluginsvc.PluginSecret{
			{
				Name:          "secret1",
				Description:   "secret test",
				AcquireMethod: "from github settings",
				Link:          "https://github.com",
			},
		},
		LabelTexts: []string{"label1", "label2"},
		Logo:       "https://github.com",
	}))
}

func TestUpdatePlugin(t *testing.T) {
	tests.RpcCallWrap(rpc.PluginCli.UpdatePlugin(ctx, &pluginsvc.UpdatePluginReq{
		Id: 1,
		Secrets: []*pluginsvc.PluginSecret{
			{
				Name:          "secret1",
				Description:   "secret test",
				AcquireMethod: "from github settings",
				Link:          "https://github.com",
			},
		},
		LabelIds:   []int64{3, 5},
		LabelTexts: []string{"image"},
	}))
}

func TestListPlugin(t *testing.T) {
	tests.RpcCallWrap(rpc.PluginCli.ListPlugin(ctx, &pluginsvc.ListPluginReq{}))
}

func TestCreatePluginTool(t *testing.T) {
	reqType := caller.RequestType{
		Type: jsonschema.Object,
		Properties: map[string]schema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "The city and state, e.g. San Francisco, CA",
			},
			"unit": {
				Type:        jsonschema.String,
				Description: "The unit of the weather information, e.g. Celsius, Fahrenheit",
			},
		},
		Required: []string{"location", "unit"},
	}
	requestType, _ := json.Marshal(reqType)

	tests.RpcCallWrap(rpc.PluginCli.CreateTool(ctx, &pluginsvc.CreatePluginToolReq{
		PluginId:    1,
		Name:        "get_current_weather",
		Description: "Get the current weather in a given location",
		RequestType: requestType,
		ApiUrl:      "https://api.openweathermap.org/data/2.5/weather",
	}))
}

func TestUpdatePluginTool(t *testing.T) {
	reqType := &caller.RequestType{
		Type: jsonschema.Object,
		Properties: map[string]schema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "The city and state, e.g. San Francisco, CA",
			},
			"unit": {
				Type:        jsonschema.String,
				Description: "The unit of the weather information, e.g. Celsius, Fahrenheit",
			},
		},
		Required: []string{"location", "unit"},
	}

	requestType, _ := reqType.Def().MarshalJSON()

	respType := &caller.ResponseType{
		Type: jsonschema.Object,
		Properties: map[string]schema.Definition{
			"code": {
				Type: jsonschema.Integer,
			},
			"message": {
				Type: jsonschema.String,
			},
			"data": {
				Type: jsonschema.Object,
				Properties: map[string]schema.Definition{
					"temperature": {
						Type: jsonschema.Number,
					},
				},
			},
		},
		Required: []string{"code", "message"},
	}
	responseType, _ := respType.Def().MarshalJSON()

	tests.RpcCallWrap(rpc.PluginCli.UpdateTool(ctx, &pluginsvc.UpdatePluginToolReq{
		Id:           1,
		RequestType:  requestType,
		ResponseType: responseType,
	}))
}
