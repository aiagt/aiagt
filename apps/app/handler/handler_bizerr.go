// Code generated by tools/init_service. DO NOT EDIT.

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

	bizCreateApp = bizerr.NewBiz(ServiceName, "create_app", baseCode+bizCodeCreateApp)
	bizDeleteApp = bizerr.NewBiz(ServiceName, "delete_app", baseCode+bizCodeDeleteApp)
	bizGetAppByID = bizerr.NewBiz(ServiceName, "get_app_by_id", baseCode+bizCodeGetAppByID)
	bizListApp = bizerr.NewBiz(ServiceName, "list_app", baseCode+bizCodeListApp)
	bizListAppLabel = bizerr.NewBiz(ServiceName, "list_app_label", baseCode+bizCodeListAppLabel)
	bizPublishApp = bizerr.NewBiz(ServiceName, "publish_app", baseCode+bizCodePublishApp)
	bizUpdateApp = bizerr.NewBiz(ServiceName, "update_app", baseCode+bizCodeUpdateApp)
}
