package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/apps/plugin/mapper"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetToolByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetToolByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.PluginTool, err error) {
	tool, err := s.toolDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetToolByID.NewErr(err)
	}

	plugin, err := s.pluginDao.GetByID(ctx, tool.PluginID)
	if err != nil {
		return nil, bizGetToolByID.NewErr(err)
	}

	if plugin.IsPrivate && ctxutil.Forbidden(ctx, plugin.AuthorID) {
		return nil, bizGetToolByID.CodeErr(bizerr.ErrCodeForbidden)
	}

	resp = mapper.NewGenPluginTool(tool)

	return
}
