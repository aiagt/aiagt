package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/plugin/mapper"
	"github.com/aiagt/aiagt/common/bizerr"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetPluginByKey implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetPluginByKey(ctx context.Context, req *pluginsvc.GetPluginByKeyReq) (resp *pluginsvc.Plugin, err error) {
	plugin, err := s.pluginDao.GetByKey(ctx, req.Key)
	if err != nil {
		return nil, bizGetPluginByID.NewErr(err).Log(ctx, "get plugin by key failed")
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
