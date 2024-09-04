package handler

import (
	"context"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// TestPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) TestPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (resp *pluginsvc.TestPluginToolResp, err error) {
	return
}
