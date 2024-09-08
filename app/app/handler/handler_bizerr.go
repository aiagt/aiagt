package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "app"

	bizCodeCreateApp    = 0
	bizCodeDeleteApp    = 1
	bizCodeGetAppByID   = 2
	bizCodeListApp      = 3
	bizCodeListAppLabel = 4
	bizCodePublishApp   = 5
	bizCodeUpdateApp    = 6
)

var (
	bizCreateApp    *bizerr.Biz
	bizDeleteApp    *bizerr.Biz
	bizGetAppByID   *bizerr.Biz
	bizListApp      *bizerr.Biz
	bizListAppLabel *bizerr.Biz
	bizPublishApp   *bizerr.Biz
	bizUpdateApp    *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizCreateApp = bizerr.NewBiz(ServiceName, "CreateApp", baseCode+bizCodeCreateApp)
	bizDeleteApp = bizerr.NewBiz(ServiceName, "DeleteApp", baseCode+bizCodeDeleteApp)
	bizGetAppByID = bizerr.NewBiz(ServiceName, "GetAppByID", baseCode+bizCodeGetAppByID)
	bizListApp = bizerr.NewBiz(ServiceName, "ListApp", baseCode+bizCodeListApp)
	bizListAppLabel = bizerr.NewBiz(ServiceName, "ListAppLabel", baseCode+bizCodeListAppLabel)
	bizPublishApp = bizerr.NewBiz(ServiceName, "PublishApp", baseCode+bizCodePublishApp)
	bizUpdateApp = bizerr.NewBiz(ServiceName, "UpdateApp", baseCode+bizCodeUpdateApp)
}
