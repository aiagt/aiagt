package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
	"github.com/aiagt/aiagt/apps/plugin/model"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/pkg/hash/hset"
)

// GetPluginSecretsByTools implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetPluginSecretsByTools(ctx context.Context, req *base.IDsReq) (resp []*pluginsvc.PluginSecrets, err error) {
	pluginTools, err := s.toolDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		return nil, bizGetPluginSecretsByTools.NewErr(err).Log(ctx, "get plugin tools by ids error")
	}

	pluginIDs := hset.FromSlice(pluginTools, func(t *model.PluginTool) int64 { return t.PluginID }).List()

	plugins, err := s.pluginDao.GetByIDs(ctx, pluginIDs)
	if err != nil {
		return nil, bizGetPluginSecretsByTools.NewErr(err).Log(ctx, "get plugins by ids error")
	}

	resp = mapper.NewGenListPluginSecrets(plugins)

	return
}
