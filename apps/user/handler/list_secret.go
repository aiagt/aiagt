package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/user/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/hash/hset"

	"github.com/aiagt/aiagt/apps/user/mapper"
	"github.com/aiagt/aiagt/common/ctxutil"

	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// ListSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) ListSecret(ctx context.Context, req *usersvc.ListSecretReq) (resp *usersvc.ListSecretResp, err error) {
	userID := ctxutil.UserID(ctx)

	list, page, err := s.secretDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListSecret.NewErr(err).Log(ctx, "get secrets error")
	}

	pluginIDs := hset.FromSliceEntries(list, func(t *model.Secret) int64 { return t.PluginID }).List()

	plugins, err := s.pluginCli.GetPluginByIDs(ctx, &base.IDsReq{Ids: pluginIDs})
	if err != nil {
		return nil, bizListSecret.CallErr(err).Log(ctx, "get plugins by ids error")
	}

	pluginMap := hmap.FromSliceEntries(plugins, func(t *pluginsvc.Plugin) (int64, *pluginsvc.Plugin, bool) {
		if t.IsPrivate && t.AuthorId != userID {
			return t.Id, nil, false
		}

		return t.Id, t, true
	})

	resp = &usersvc.ListSecretResp{
		Pagination: page,
		Secrets:    mapper.NewGenListSecret(list, pluginMap),
	}

	return
}
