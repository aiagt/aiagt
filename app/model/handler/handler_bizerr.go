package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "model"

	bizCodeChat     = 0
	bizCodeGenToken = 1
)

var (
	bizChat     *bizerr.Biz
	bizGenToken *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizChat = bizerr.NewBiz(ServiceName, "Chat", baseCode+bizCodeChat)
	bizGenToken = bizerr.NewBiz(ServiceName, "GenToken", baseCode+bizCodeGenToken)
}
