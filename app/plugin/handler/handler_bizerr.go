package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "plugin"

	bizCodeCallPluginTool  = 0
	bizCodeCreatePlugin    = 1
	bizCodeCreateTool      = 2
	bizCodeDeletePlugin    = 3
	bizCodeDeleteTool      = 4
	bizCodeGetPluginByID   = 5
	bizCodeGetToolByID     = 6
	bizCodeListPlugin      = 7
	bizCodeListPluginLabel = 8
	bizCodeListPluginTool  = 9
	bizCodePublishPlugin   = 10
	bizCodeTestPluginTool  = 11
	bizCodeUpdatePlugin    = 12
	bizCodeUpdateTool      = 13
)

var (
	bizCallPluginTool  *bizerr.Biz
	bizCreatePlugin    *bizerr.Biz
	bizCreateTool      *bizerr.Biz
	bizDeletePlugin    *bizerr.Biz
	bizDeleteTool      *bizerr.Biz
	bizGetPluginByID   *bizerr.Biz
	bizGetToolByID     *bizerr.Biz
	bizListPlugin      *bizerr.Biz
	bizListPluginLabel *bizerr.Biz
	bizListPluginTool  *bizerr.Biz
	bizPublishPlugin   *bizerr.Biz
	bizTestPluginTool  *bizerr.Biz
	bizUpdatePlugin    *bizerr.Biz
	bizUpdateTool      *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizCallPluginTool = bizerr.NewBiz(ServiceName, "CallPluginTool", baseCode+bizCodeCallPluginTool)
	bizCreatePlugin = bizerr.NewBiz(ServiceName, "CreatePlugin", baseCode+bizCodeCreatePlugin)
	bizCreateTool = bizerr.NewBiz(ServiceName, "CreateTool", baseCode+bizCodeCreateTool)
	bizDeletePlugin = bizerr.NewBiz(ServiceName, "DeletePlugin", baseCode+bizCodeDeletePlugin)
	bizDeleteTool = bizerr.NewBiz(ServiceName, "DeleteTool", baseCode+bizCodeDeleteTool)
	bizGetPluginByID = bizerr.NewBiz(ServiceName, "GetPluginByID", baseCode+bizCodeGetPluginByID)
	bizGetToolByID = bizerr.NewBiz(ServiceName, "GetToolByID", baseCode+bizCodeGetToolByID)
	bizListPlugin = bizerr.NewBiz(ServiceName, "ListPlugin", baseCode+bizCodeListPlugin)
	bizListPluginLabel = bizerr.NewBiz(ServiceName, "ListPluginLabel", baseCode+bizCodeListPluginLabel)
	bizListPluginTool = bizerr.NewBiz(ServiceName, "ListPluginTool", baseCode+bizCodeListPluginTool)
	bizPublishPlugin = bizerr.NewBiz(ServiceName, "PublishPlugin", baseCode+bizCodePublishPlugin)
	bizTestPluginTool = bizerr.NewBiz(ServiceName, "TestPluginTool", baseCode+bizCodeTestPluginTool)
	bizUpdatePlugin = bizerr.NewBiz(ServiceName, "UpdatePlugin", baseCode+bizCodeUpdatePlugin)
	bizUpdateTool = bizerr.NewBiz(ServiceName, "UpdateTool", baseCode+bizCodeUpdateTool)
}
