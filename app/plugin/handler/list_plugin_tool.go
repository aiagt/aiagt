package handler

import (
	"context"
	"github.com/aiagt/aiagt/app/plugin/mapping"
	"github.com/aiagt/aiagt/common/ctxutil"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// ListPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListPluginTool(ctx context.Context, req *pluginsvc.ListPluginToolReq) (resp *pluginsvc.ListPluginToolResp, err error) {
	userID := ctxutil.UserID(ctx)

	list, page, err := s.toolDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListPluginTool.NewErr(err)
	}

	resp = &pluginsvc.ListPluginToolResp{
		Plugins:    mapping.NewGenListPluginTool(list),
		Pagination: page,
	}

	return
}
