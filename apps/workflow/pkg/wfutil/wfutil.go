package wfutil

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	pluginservice "github.com/aiagt/aiagt/kitex_gen/pluginsvc/pluginservice"
	"github.com/aiagt/aiagt/pkg/workflow"
)

type PluginNodeRunner struct {
	pluginID  int64
	toolID    int64
	secrets   map[string]string
	pluginCli pluginservice.Client
}

func NewPluginNodeRunner(pluginID, toolID int64, secrets map[string]string, pluginCli pluginservice.Client) *PluginNodeRunner {
	return &PluginNodeRunner{
		pluginID:  pluginID,
		toolID:    toolID,
		secrets:   secrets,
		pluginCli: pluginCli,
	}
}

func (r *PluginNodeRunner) Run(ctx context.Context, input workflow.Object) (workflow.Object, error) {
	req, err := input.JSON()
	if err != nil {
		return nil, fmt.Errorf("encode input object error: %w", err)
	}

	resp, err := r.pluginCli.CallPluginTool(ctx, &pluginsvc.CallPluginToolReq{
		PluginId: r.pluginID,
		ToolId:   r.toolID,
		Secrets:  r.secrets,
		Request:  req,
	})
	if err != nil {
		return nil, fmt.Errorf("plugin call error: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("plugin call error, code: %d, http_code: %d, message: %s", resp.Code, resp.HttpCode, resp.Msg)
	}

	output, err := workflow.NewJSONObject(resp.Response)
	if err != nil {
		return nil, fmt.Errorf("decode output object error: %w", err)
	}

	return output, nil
}
