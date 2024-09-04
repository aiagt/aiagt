package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// UpdatePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) UpdatePlugin(ctx context.Context, req *pluginsvc.UpdatePluginReq) (resp *base.Empty, err error) {
	return
}
