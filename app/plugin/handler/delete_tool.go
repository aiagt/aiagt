package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) DeleteTool(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	tool, err := s.toolDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteTool.NewErr(err)
	}

	plugin, err := s.pluginDao.GetByID(ctx, tool.PluginID)
	if err != nil {
		return nil, bizDeleteTool.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, plugin.AuthorID) {
		return nil, bizDeleteTool.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.toolDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, bizDeleteTool.NewErr(err)
	}

	return
}
