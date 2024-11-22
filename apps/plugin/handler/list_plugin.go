package handler

import (
	"context"
	"github.com/aiagt/aiagt/pkg/lists"

	"github.com/aiagt/aiagt/apps/plugin/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/pkg/hash/hmap"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// ListPlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListPlugin(ctx context.Context, req *pluginsvc.ListPluginReq) (resp *pluginsvc.ListPluginResp, err error) {
	userID := ctxutil.UserID(ctx)

	plugins, page, err := s.pluginDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListPlugin.NewErr(err)
	}

	labelIDs := lists.FlatMap(plugins, func(t *model.Plugin) []int64 {
		return t.LabelIDs
	})
	labels, err := s.labelDao.GetByIDs(ctx, labelIDs)

	labelMap := hmap.FromSliceEntries(labels, func(t *model.PluginLabel) (int64, *pluginsvc.PluginLabel, bool) {
		return t.ID, mapper.NewGenPluginLabel(t), true
	})

	resp = &pluginsvc.ListPluginResp{
		Plugins:    mapper.NewGenListPlugin(plugins, labelMap),
		Pagination: page,
	}

	return
}
