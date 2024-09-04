package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// UpdateTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) UpdateTool(ctx context.Context, req *pluginsvc.UpdatePluginToolReq) (resp *base.Empty, err error) {
	return
}
