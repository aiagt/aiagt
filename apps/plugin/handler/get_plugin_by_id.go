package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
	"github.com/aiagt/aiagt/common/bizerr"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetPluginByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetPluginByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.Plugin, err error) {
	plugin, err := s.pluginDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetPluginById.NewErr(err).Log(ctx, "get plugin by id failed")
	}

	userID := ctxutil.UserID(ctx)
	if plugin.IsPrivate && plugin.AuthorID != userID {
		return nil, bizGetPluginById.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	author, err := s.userCli.GetUserByID(ctx, &base.IDReq{Id: plugin.AuthorID})
	if err != nil {
		return nil, bizGetPluginById.CallErr(err).Log(ctx, "get user failed")
	}

	labels, err := s.labelDao.GetByIDs(ctx, plugin.LabelIDs)
	if err != nil {
		return nil, bizGetPluginById.NewErr(err).Log(ctx, "get labels failed")
	}

	tools, err := s.toolDao.GetByPluginID(ctx, plugin.ID)
	if err != nil {
		return nil, bizGetPluginById.NewErr(err).Log(ctx, "get tools failed")
	}

	resp = mapper.NewGenPlugin(
		plugin,
		author,
		mapper.NewGenListPluginLabel(labels),
		mapper.NewGenListPluginTool(tools),
	)

	return
}
