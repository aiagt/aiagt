package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/plugin/mapping"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// UpdateTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) UpdateTool(ctx context.Context, req *pluginsvc.UpdatePluginToolReq) (resp *base.Empty, err error) {
	tool, err := s.toolDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizUpdateTool.NewErr(err)
	}

	plugin, err := s.pluginDao.GetByID(ctx, tool.PluginID)
	if err != nil {
		return nil, bizUpdateTool.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, plugin.AuthorID) {
		return nil, bizUpdateTool.CodeErr(bizerr.ErrCodeForbidden)
	}

	newTool := mapping.NewModelUpdatePluginTool(req)

	err = s.toolDao.Update(ctx, req.Id, newTool)
	if err != nil {
		return nil, bizUpdateTool.NewErr(err)
	}

	return
}
