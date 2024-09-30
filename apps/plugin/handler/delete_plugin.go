package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeletePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) DeletePlugin(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	plugin, err := s.pluginDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizDeletePlugin.NewErr(err)
	}

	if ctxutil.Forbidden(ctx, plugin.AuthorID) {
		return nil, bizDeletePlugin.CodeErr(bizerr.ErrCodeForbidden)
	}

	err = s.pluginDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, bizDeletePlugin.NewErr(err)
	}

	return
}
