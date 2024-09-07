package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/plugin/mapping"
	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// CreatePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CreatePlugin(ctx context.Context, req *pluginsvc.CreatePluginReq) (resp *base.Empty, err error) {
	labelIDs, err := s.labelDao.UpdateLabels(ctx, req.LabelIds, req.LabelTexts)
	if err != nil {
		return nil, bizCreatePlugin.NewErr(err)
	}

	user := ctxutil.User(ctx)
	plugin := mapping.NewModelCreatePlugin(req, user, labelIDs)

	if err = s.pluginDao.Create(ctx, plugin); err != nil {
		return nil, bizCreatePlugin.NewErr(err)
	}

	return
}
