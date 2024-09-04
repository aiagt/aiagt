package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetToolByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetToolByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.PluginTool, err error) {
	return
}
