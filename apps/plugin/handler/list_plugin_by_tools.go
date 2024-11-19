package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
	"github.com/aiagt/aiagt/apps/plugin/model"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/pkg/hash/hset"
)

// ListPluginByTools implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListPluginByTools(ctx context.Context, req *pluginsvc.ListPluginByToolsReq) (resp *pluginsvc.ListPluginByToolsResp, err error) {
	tools, err := s.toolDao.GetByIDs(ctx, req.ToolIds)
	if err != nil {
		return nil, bizListPluginByTools.NewErr(err).Log(ctx, "get tools by ids error")
	}

	pluginIDs := hset.FromSlice(tools, func(t *model.PluginTool) int64 { return t.PluginID })

	plugins, err := s.pluginDao.GetByIDs(ctx, pluginIDs.List())
	if err != nil {
		return nil, bizListPluginByTools.NewErr(err).Log(ctx, "get plugins by ids error")
	}

	resp = &pluginsvc.ListPluginByToolsResp{
		Plugins: mapper.NewGenListPlugin(plugins, nil),
	}

	return
}
