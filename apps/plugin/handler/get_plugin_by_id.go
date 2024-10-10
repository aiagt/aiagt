package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
	"github.com/aiagt/aiagt/common/bizerr"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetPluginByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetPluginByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.Plugin, err error) {
	plugin, err := s.pluginDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetPluginByID.NewErr(err).Log(ctx, "get plugin by id failed")
	}

	user, err := s.userCli.GetUser(ctx)
	if err != nil {
		return nil, bizGetPluginByID.CallErr(err).Log(ctx, "get user failed")
	}

	if plugin.IsPrivate && plugin.AuthorID != user.Id {
		return nil, bizGetPluginByID.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	labels, err := s.labelDao.GetByIDs(ctx, plugin.LabelIDs)
	if err != nil {
		return nil, bizGetPluginByID.NewErr(err).Log(ctx, "get labels failed")
	}

	tools, err := s.toolDao.GetByPluginID(ctx, plugin.ID)
	if err != nil {
		return nil, bizGetPluginByID.NewErr(err).Log(ctx, "get tools failed")
	}

	resp = mapper.NewGenPlugin(
		plugin,
		user,
		mapper.NewGenListPluginLabel(labels),
		mapper.NewGenListPluginTool(tools),
	)

	return
}
