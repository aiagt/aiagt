package handler

import (
	"context"
	"github.com/aiagt/aiagt/apps/plugin/model"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/hash/hset"

	"github.com/aiagt/aiagt/apps/plugin/mapper"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// AllPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) AllPluginTool(ctx context.Context, req *pluginsvc.AllPluginToolReq) (resp []*pluginsvc.PluginTool, err error) {
	tools, err := s.toolDao.GetByIDs(ctx, req.ToolIds)
	if err != nil {
		return nil, bizAllPluginTool.NewErr(err).Log(ctx, "get plugin tools error")
	}

	pluginIDs := hset.FromSlice(tools, func(t *model.PluginTool) int64 { return t.PluginID })

	plugins, err := s.pluginDao.GetByIDs(ctx, pluginIDs.List())
	if err != nil {
		return nil, bizAllPluginTool.NewErr(err).Log(ctx, "get plugins error")
	}

	pluginMap := hmap.FromSliceEntries(plugins, func(t *model.Plugin) (int64, *model.Plugin, bool) {
		return t.ID, t, true
	})

	resp = mapper.NewGenListPluginToolWithPlugin(tools, pluginMap)

	return
}
