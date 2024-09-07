package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/plugin/mapping"
	"github.com/aiagt/aiagt/common/bizerr"

	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// UpdatePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) UpdatePlugin(ctx context.Context, req *pluginsvc.UpdatePluginReq) (resp *base.Empty, err error) {
	plugin, err := s.pluginDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizUpdatePlugin.NewErr(err)
	}

	user := ctxutil.User(ctx)
	if user.Id != plugin.AuthorID {
		return nil, bizUpdatePlugin.CodeErr(bizerr.ErrCodeForbidden)
	}

	labelIDs, err := s.labelDao.UpdateLabels(ctx, req.LabelIds, req.LabelTexts)
	if err != nil {
		return nil, bizUpdatePlugin.NewErr(err)
	}

	newPlugin := mapping.NewModelUpdatePlugin(req, user, labelIDs)

	if err = s.pluginDao.Update(ctx, req.Id, newPlugin); err != nil {
		return nil, bizUpdatePlugin.NewErr(err)
	}

	return
}
