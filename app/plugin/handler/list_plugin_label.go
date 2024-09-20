package handler

import (
	"context"

	"github.com/aiagt/aiagt/app/plugin/mapper"

	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// ListPluginLabel implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListPluginLabel(ctx context.Context, req *pluginsvc.ListPluginLabelReq) (resp *pluginsvc.ListPluginLabelResp, err error) {
	list, page, err := s.labelDao.List(ctx, req)
	if err != nil {
		return nil, bizListPluginLabel.NewErr(err)
	}

	resp = &pluginsvc.ListPluginLabelResp{
		Labels:     mapper.NewGenListPluginLabel(list),
		Pagination: page,
	}

	return
}
