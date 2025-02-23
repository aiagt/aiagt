package handler

import (
	"context"

	"github.com/aiagt/aiagt/pkg/utils"

	"github.com/aiagt/aiagt/apps/app/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetAppByID implements the AppServiceImpl interface.
func (s *AppServiceImpl) GetAppByID(ctx context.Context, req *appsvc.GetAppByIDReq) (resp *appsvc.GetAppByIDResp, err error) {
	app, err := s.appDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetAppByID.NewErr(err)
	}

	userID := ctxutil.UserID(ctx)
	if app.IsPrivate && app.AuthorID != userID {
		return nil, bizGetAppByID.CodeErr(bizerr.ErrCodeForbidden)
	}

	author, err := s.userCli.GetUserByID(ctx, &base.IDReq{Id: app.AuthorID})
	if err != nil {
		return nil, bizGetAppByID.NewErr(err)
	}

	tools, err := s.pluginCli.AllPluginTool(ctx, &pluginsvc.AllPluginToolReq{ToolIds: app.ToolIDs})
	if err != nil {
		return nil, bizGetAppByID.CallErr(err)
	}

	pluginSecrets, err := s.pluginCli.GetPluginSecretsByTools(ctx, &base.IDsReq{Ids: app.ToolIDs})
	if err != nil {
		return nil, bizGetAppByID.CallErr(err).Log(ctx, "get plugin secrets by tool_ids error")
	}

	var (
		publicTools       []*pluginsvc.PluginTool
		privateToolsCount int32
	)

	if !utils.Value(req.Unfold) {
		for _, tool := range tools {
			if tool.Plugin.IsPrivate && tool.Plugin.AuthorId != userID {
				privateToolsCount++
			} else {
				publicTools = append(publicTools, tool)
			}
		}

		tools = publicTools

		for _, pluginSecret := range pluginSecrets {
			if pluginSecret.IsPrivate && pluginSecret.AuthorId != userID {
				pluginSecret.PluginName = ""
				pluginSecret.PluginLogo = ""
			}
		}
	}

	labels, err := s.labelDao.GetByIDs(ctx, app.LabelIDs)
	if err != nil {
		return nil, bizGetAppByID.NewErr(err)
	}

	resp = &appsvc.GetAppByIDResp{
		App: mapper.NewGenApp(app, author, tools, mapper.NewGenListAppLabel(labels), pluginSecrets),
		Ext: &appsvc.GetAppByIDRespExtend{PrivateToolsCount: privateToolsCount},
	}

	return
}
