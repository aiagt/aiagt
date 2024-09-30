package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/app/mapper"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// GetAppByID implements the AppServiceImpl interface.
func (s *AppServiceImpl) GetAppByID(ctx context.Context, req *base.IDReq) (resp *appsvc.App, err error) {
	app, err := s.appDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizGetAppByID.NewErr(err)
	}

	userID := ctxutil.UserID(ctx)
	if app.AuthorID != userID {
		return nil, bizGetAppByID.CodeErr(bizerr.ErrCodeForbidden)
	}

	author, err := s.userCli.GetUserByID(ctx, &base.IDReq{Id: userID})
	if err != nil {
		return nil, bizGetAppByID.NewErr(err)
	}

	listToolResp, err := s.pluginCli.ListPluginTool(ctx, &pluginsvc.ListPluginToolReq{
		ToolIds: app.ToolIDs,
	})
	if err != nil {
		return nil, bizGetAppByID.NewErr(err)
	}

	labels, err := s.labelDao.GetByIDs(ctx, app.LabelIDs)
	if err != nil {
		return nil, bizGetAppByID.NewErr(err)
	}

	resp = mapper.NewGenApp(app, author, listToolResp.Tools, mapper.NewGenListAppLabel(labels))

	return
}
