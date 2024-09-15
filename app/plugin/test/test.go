package main

import (
	"context"
	"encoding/json"

	"github.com/aiagt/aiagt/pkg/call"
	"github.com/sashabaranov/go-openai/jsonschema"

	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
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

	// logger(GetPlugin(ctx))
	// logger(ListPlugin(ctx))
	// logger(CreatePlugigTool(ctx))
	logger(UpdatePluginTool(ctx))
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
	klog.Infof("resp: %#v", resp)
}

func CreatePlugin(ctx context.Context) (any, error) {
	return rpc.PluginCli.CreatePlugin(ctx, &pluginsvc.CreatePluginReq{
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
	})
}

func UpdatePlugin(ctx context.Context) (any, error) {
	return rpc.PluginCli.UpdatePlugin(ctx, &pluginsvc.UpdatePluginReq{
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
	})
}

func GetPlugin(ctx context.Context) (any, error) {
	return rpc.PluginCli.GetPluginByID(ctx, &base.IDReq{Id: 1})
}

func ListPlugin(ctx context.Context) (any, error) {
	return rpc.PluginCli.ListPlugin(ctx, &pluginsvc.ListPluginReq{})
}

func CreatePlugigTool(ctx context.Context) (any, error) {
	reqType := call.RequestType{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
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
		// ContentType: "application/json",
		// Parameters: call.Object{
		// 	{Name: "location", Description: "The city and state, e.g. San Francisco, CA", Type: "string"},
		// 	{Name: "unit", Description: "The unit of the weather information, e.g. Celsius, Fahrenheit", Type: "string"},
		// },
	}
	requestType, _ := json.Marshal(reqType)

	return rpc.PluginCli.CreateTool(ctx, &pluginsvc.CreatePluginToolReq{
		PluginId:    1,
		Name:        "get_current_weather",
		Description: "Get the current weather in a given location",
		RequestType: requestType,
		ApiUrl:      "https://api.openweathermap.org/data/2.5/weather",
	})
}

func UpdatePluginTool(ctx context.Context) (any, error) {
	reqType := &call.RequestType{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
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

	requestType, _ := reqType.MarshalJSON()

	respType := &call.ResponseType{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"code": {
				Type: jsonschema.Integer,
			},
			"message": {
				Type: jsonschema.String,
			},
			"data": {
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"temperature": {
						Type: jsonschema.Number,
					},
				},
			},
		},
		Required: []string{"code", "message"},
	}
	responseType, _ := respType.MarshalJSON()

	return rpc.PluginCli.UpdateTool(ctx, &pluginsvc.UpdatePluginToolReq{
		Id:           1,
		RequestType:  requestType,
		ResponseType: responseType,
	})
}
