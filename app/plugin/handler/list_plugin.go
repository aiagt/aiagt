package handler

import (
	"context"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// ListPlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListPlugin(ctx context.Context, req *pluginsvc.ListPluginReq) (resp *pluginsvc.ListPluginResp, err error) {
	plugins, page, err := s.pluginDao.List(ctx, req)
	if err != nil {
		return nil, bizListPlugin.NewErr(1, err)
	}

	resp = &pluginsvc.ListPluginResp{
		Plugins:    MapListPlugin(plugins),
		Pagination: page,
	}
	return
}
