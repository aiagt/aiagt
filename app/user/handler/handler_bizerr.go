package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "user"

	bizCodeCreateSecret  = 0
	bizCodeDeleteSecret  = 1
	bizCodeGetUser       = 2
	bizCodeGetUserByID   = 3
	bizCodeGetUserByIds  = 4
	bizCodeListSecret    = 5
	bizCodeLogin         = 6
	bizCodeRegister      = 7
	bizCodeResetPassword = 8
	bizCodeSendCaptcha   = 9
	bizCodeUpdateSecret  = 10
	bizCodeUpdateUser    = 11
)

var (
	bizCreateSecret  *bizerr.Biz
	bizDeleteSecret  *bizerr.Biz
	bizGetUser       *bizerr.Biz
	bizGetUserByID   *bizerr.Biz
	bizGetUserByIds  *bizerr.Biz
	bizListSecret    *bizerr.Biz
	bizLogin         *bizerr.Biz
	bizRegister      *bizerr.Biz
	bizResetPassword *bizerr.Biz
	bizSendCaptcha   *bizerr.Biz
	bizUpdateSecret  *bizerr.Biz
	bizUpdateUser    *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizCreateSecret = bizerr.NewBiz(ServiceName, "CreateSecret", baseCode+bizCodeCreateSecret)
	bizDeleteSecret = bizerr.NewBiz(ServiceName, "DeleteSecret", baseCode+bizCodeDeleteSecret)
	bizGetUser = bizerr.NewBiz(ServiceName, "GetUser", baseCode+bizCodeGetUser)
	bizGetUserByID = bizerr.NewBiz(ServiceName, "GetUserByID", baseCode+bizCodeGetUserByID)
	bizGetUserByIds = bizerr.NewBiz(ServiceName, "GetUserByIds", baseCode+bizCodeGetUserByIds)
	bizListSecret = bizerr.NewBiz(ServiceName, "ListSecret", baseCode+bizCodeListSecret)
	bizLogin = bizerr.NewBiz(ServiceName, "Login", baseCode+bizCodeLogin)
	bizRegister = bizerr.NewBiz(ServiceName, "Register", baseCode+bizCodeRegister)
	bizResetPassword = bizerr.NewBiz(ServiceName, "ResetPassword", baseCode+bizCodeResetPassword)
	bizSendCaptcha = bizerr.NewBiz(ServiceName, "SendCaptcha", baseCode+bizCodeSendCaptcha)
	bizUpdateSecret = bizerr.NewBiz(ServiceName, "UpdateSecret", baseCode+bizCodeUpdateSecret)
	bizUpdateUser = bizerr.NewBiz(ServiceName, "UpdateUser", baseCode+bizCodeUpdateUser)
}
