package handler

import (
	"context"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// CallPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CallPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (resp *pluginsvc.CallPluginToolResp, err error) {
	return
}
