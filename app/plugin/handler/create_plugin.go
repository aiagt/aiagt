package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// CreatePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CreatePlugin(ctx context.Context, req *pluginsvc.CreatePluginReq) (resp *base.Empty, err error) {
	return
}
