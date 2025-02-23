package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetPluginByIDs implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetPluginByIDs(ctx context.Context, req *base.IDsReq) (resp []*pluginsvc.Plugin, err error) {
	plugins, err := s.pluginDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		return nil, bizGetPluginByIds.NewErr(err).Log(ctx, "get plugin by ids")
	}

	resp = mapper.NewGenListPlugin(plugins, nil)

	return
}
