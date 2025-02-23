package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/plugin/model"
	"github.com/aiagt/aiagt/pkg/hash/hmap"

	"github.com/aiagt/aiagt/pkg/caller"
	"github.com/aiagt/aiagt/pkg/hash/hset"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// CallPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CallPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (resp *pluginsvc.CallPluginToolResp, err error) {
	tool, err := s.toolDao.GetByID(ctx, req.ToolId)
	if err != nil {
		return nil, bizCallPluginTool.NewErr(err)
	}

	if tool.PluginID != req.PluginId {
		return nil, bizCallPluginTool.CodeErr(bizerr.ErrCodeNotExists)
	}

	plugin, err := s.pluginDao.GetByID(ctx, tool.PluginID)
	if err != nil {
		return nil, bizCallPluginTool.NewErr(err)
	}

	userID, ok := ctxutil.GetUserID(ctx)
	if !ok {
		return nil, bizCallPluginTool.NewErr(err)
	}

	// plugin call do not require permission verification
	//if plugin.AuthorID != userID {
	//	return nil, bizCallPluginTool.CodeErr(bizerr.ErrCodeForbidden)
	//}

	listSecret, err := s.userCli.ListSecret(ctx, &usersvc.ListSecretReq{
		Pagination: &base.PaginationReq{
			PageSize: int32(len(plugin.Secrets)),
		},
		PluginId: &req.PluginId,
	})
	if err != nil {
		return nil, bizCallPluginTool.NewErr(err)
	}

	secretDefs := hset.FromSliceEntries(plugin.Secrets, func(t *model.PluginSecret) string { return t.Name })

	userSecretMap := hmap.FromSliceEntries(listSecret.Secrets, func(t *usersvc.Secret) (string, string, bool) {
		return t.Name, t.Value, t.PluginId == plugin.ID && secretDefs.Has(t.Name)
	})

	body := &caller.RequestBody{
		PluginID: tool.PluginID,
		ToolID:   tool.ID,
		UserID:   userID,
		Secrets:  userSecretMap,
	}

	callResp, callRespCode, err := caller.Call(ctx, body, tool.ApiURL, tool.RequestType, tool.ResponseType, req.Request)

	resp = &pluginsvc.CallPluginToolResp{Response: callResp, HttpCode: int64(callRespCode)}
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}

	return resp, nil
}
