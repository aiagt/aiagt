package main

import (
	"context"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// PluginServiceImpl implements the last service interface defined in the IDL.
type PluginServiceImpl struct{}

// CreatePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CreatePlugin(ctx context.Context, req *pluginsvc.CreatePluginReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// UpdatePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) UpdatePlugin(ctx context.Context, req *pluginsvc.UpdatePluginReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// DeletePlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) DeletePlugin(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// GetPluginByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetPluginByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.Plugin, err error) {
	// TODO: Your code here...
	return
}

// ListPlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListPlugin(ctx context.Context, req *pluginsvc.ListPluginReq) (resp *pluginsvc.ListPluginResp, err error) {
	// TODO: Your code here...
	return
}

// CreateTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CreateTool(ctx context.Context, req *pluginsvc.CreatePluginToolReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// UpdateTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) UpdateTool(ctx context.Context, req *pluginsvc.UpdatePluginToolReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// DeleteTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) DeleteTool(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// GetToolByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetToolByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.PluginTool, err error) {
	// TODO: Your code here...
	return
}

// ListTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListTool(ctx context.Context, req *pluginsvc.ListPluginToolReq) (resp *pluginsvc.ListPluginToolResp, err error) {
	// TODO: Your code here...
	return
}

// CallPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) CallPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (resp *pluginsvc.CallPluginToolResp, err error) {
	// TODO: Your code here...
	return
}

// TestPluginTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) TestPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq) (resp *pluginsvc.TestPluginToolResp, err error) {
	// TODO: Your code here...
	return
}
