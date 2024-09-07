package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/app/plugin/mapping"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// ListPlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListPlugin(ctx context.Context, req *pluginsvc.ListPluginReq) (resp *pluginsvc.ListPluginResp, err error) {
	userID := ctxutil.UserID(ctx)

	plugins, page, err := s.pluginDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListPlugin.NewErr(err)
	}

	resp = &pluginsvc.ListPluginResp{
		Plugins:    mapping.NewGenListPlugin(plugins),
		Pagination: page,
	}

	return
}
