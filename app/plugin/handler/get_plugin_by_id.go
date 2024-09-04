package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetPluginByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetPluginByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.Plugin, err error) {
	return
}
