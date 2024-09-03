package main

import (
	"context"
	appsvc "github.com/aiagt/aiagt/kitex_gen/appsvc"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// AppServiceImpl implements the last service interface defined in the IDL.
type AppServiceImpl struct{}

// CreateApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) CreateApp(ctx context.Context, req *appsvc.CreateAppReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// UpdateApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) UpdateApp(ctx context.Context, req *appsvc.UpdateAppReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// DeleteApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) DeleteApp(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// GetAppByID implements the AppServiceImpl interface.
func (s *AppServiceImpl) GetAppByID(ctx context.Context, req *base.IDReq) (resp *appsvc.App, err error) {
	// TODO: Your code here...
	return
}

// ListApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) ListApp(ctx context.Context, req *appsvc.ListAppReq) (resp *appsvc.ListAppResp, err error) {
	// TODO: Your code here...
	return
}

// PublishApp implements the AppServiceImpl interface.
func (s *AppServiceImpl) PublishApp(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// ListAppLabel implements the AppServiceImpl interface.
func (s *AppServiceImpl) ListAppLabel(ctx context.Context, req *appsvc.ListAppLabelReq) (resp *appsvc.ListAppLabelResp, err error) {
	// TODO: Your code here...
	return
}
