package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
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

	userID, ok := ctxutil.GetUserID(ctx)
	if !ok {
		return nil, bizCreatePlugin.NewErr(err)
	}

	plugin := mapper.NewModelCreatePlugin(req, userID, labelIDs)

	if err = s.pluginDao.Create(ctx, plugin); err != nil {
		return nil, bizCreatePlugin.NewErr(err)
	}

	return
}
