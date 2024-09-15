package handler

import (
	"context"

	"github.com/aiagt/aiagt/pkg/call"
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
		return nil, bizDeleteTool.NewErr(err)
	}

	if tool.PluginID != req.PluginId {
		return nil, bizCallPluginTool.CodeErr(bizerr.ErrCodeNotExists)
	}

	plugin, err := s.pluginDao.GetByID(ctx, tool.PluginID)
	if err != nil {
		return nil, bizDeleteTool.NewErr(err)
	}

	userID, ok := ctxutil.GetUserID(ctx)
	if !ok {
		return nil, bizDeleteTool.NewErr(err)
	}

	if plugin.AuthorID != userID {
		return nil, bizDeleteTool.CodeErr(bizerr.ErrCodeForbidden)
	}

	listSecret, err := s.userCli.ListSecret(ctx, &usersvc.ListSecretReq{
		Pagination: &base.PaginationReq{
			PageSize: int32(len(plugin.Secrets)),
		},
		PluginId: &req.PluginId,
	})
	if err != nil {
		return nil, bizDeleteTool.NewErr(err)
	}

	secretSet := hset.NewSetWithKey("secret_name", plugin.Secrets...)

	secretMap := make(map[string]string, len(plugin.Secrets))
	for _, secret := range listSecret.Secrets {
		if secretSet.Has(secret.Name) {
			secretMap[secret.Name] = secret.Value
		}
	}

	body := &call.RequestBody{
		PluginID: tool.PluginID,
		ToolID:   tool.ID,
		UserID:   userID,
		Secrets:  secretMap,
	}

	callResp, err := call.Call(ctx, body, tool.ApiURL, tool.RequestType, tool.ResponseType, req.Request)
	if err != nil {
		return nil, bizDeleteTool.NewErr(err)
	}

	resp = &pluginsvc.CallPluginToolResp{
		Response: callResp,
	}

	return
}
