package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// CreateTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CreateTool(ctx context.Context, req *pluginsvc.CreatePluginToolReq) (resp *base.Empty, err error) {
	return
}
