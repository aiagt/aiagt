package handler

import (
	"context"
	"time"

	"github.com/aiagt/aiagt/apps/plugin/model"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// TestPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) TestPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (resp *pluginsvc.TestPluginToolResp, err error) {
	callResp, err := s.CallPluginTool(ctx, req)
	if err != nil {
		return nil, bizTestPluginTool.CallErr(err).Log(ctx, "call plugin tool error")
	}

	resp = &pluginsvc.TestPluginToolResp{
		Code:     callResp.Code,
		Msg:      callResp.Msg,
		Response: callResp.Response,
		HttpCode: callResp.HttpCode,
	}

	if callResp.Code != 0 {
		return
	}

	now := time.Now()

	err = s.toolDao.Update(ctx, req.ToolId, &model.PluginToolOptional{TestedAt: &now})
	if err != nil {
		return nil, bizTestPluginTool.NewErr(err).Log(ctx, "update tool tested_at error")
	}

	return
}
