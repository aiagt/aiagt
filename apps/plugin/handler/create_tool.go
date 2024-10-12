package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/apps/plugin/mapper"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// CreateTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CreateTool(ctx context.Context, req *pluginsvc.CreatePluginToolReq) (resp *base.Empty, err error) {
	plugin, err := s.pluginDao.GetByID(ctx, req.PluginId)
	if err != nil {
		return nil, bizCreateTool.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, plugin.AuthorID) {
		return nil, bizCreateTool.NewErr(bizerr.ErrForbidden)
	}

	tool := mapper.NewModelCreatePluginTool(req)

	err = s.toolDao.Create(ctx, tool)
	if err != nil {
		return nil, bizCreateTool.NewErr(err)
	}

	return
}
