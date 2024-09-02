package main

import (
	"context"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

// PluginServiceImpl implements the last service interface defined in the IDL.
type PluginServiceImpl struct{}

// Create implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) Create(ctx context.Context, req *pluginsvc.CreatePluginReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// Update implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) Update(ctx context.Context, req *pluginsvc.UpdatePluginReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// List implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) List(ctx context.Context, req *base.PaginationReq) (resp *pluginsvc.ListPluginResp, err error) {
	// TODO: Your code here...
	return
}

// GetByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.Plugin, err error) {
	// TODO: Your code here...
	return
}

// Delete implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) Delete(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
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

// ListTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) ListTool(ctx context.Context, req *base.IDReq) (resp *pluginsvc.ListPluginToolResp, err error) {
	// TODO: Your code here...
	return
}

// GetToolByID implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) GetToolByID(ctx context.Context, req *base.IDReq) (resp *pluginsvc.PluginTool, err error) {
	// TODO: Your code here...
	return
}

// DeleteTool implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) DeleteTool(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
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
