package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "user"

	bizCodeCreateSecret   = 0
	bizCodeDeleteSecret   = 1
	bizCodeForgotPassword = 2
	bizCodeGetUser        = 3
	bizCodeGetUserByIDs   = 4
	bizCodeGetUserByID    = 5
	bizCodeListSecret     = 6
	bizCodeLogin          = 7
	bizCodeRegister       = 8
	bizCodeUpdateSecret   = 9
	bizCodeUpdateUser     = 10
)

var (
	bizCreateSecret   *bizerr.Biz
	bizDeleteSecret   *bizerr.Biz
	bizForgotPassword *bizerr.Biz
	bizGetUser        *bizerr.Biz
	bizGetUserByIDs   *bizerr.Biz
	bizGetUserByID    *bizerr.Biz
	bizListSecret     *bizerr.Biz
	bizLogin          *bizerr.Biz
	bizRegister       *bizerr.Biz
	bizUpdateSecret   *bizerr.Biz
	bizUpdateUser     *bizerr.Biz
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100

	bizCreateSecret = bizerr.NewBiz(ServiceName, "CreateSecret", baseCode+bizCodeCreateSecret)
	bizDeleteSecret = bizerr.NewBiz(ServiceName, "DeleteSecret", baseCode+bizCodeDeleteSecret)
	bizForgotPassword = bizerr.NewBiz(ServiceName, "ForgotPassword", baseCode+bizCodeForgotPassword)
	bizGetUser = bizerr.NewBiz(ServiceName, "GetUser", baseCode+bizCodeGetUser)
	bizGetUserByIDs = bizerr.NewBiz(ServiceName, "GetUserByIDs", baseCode+bizCodeGetUserByIDs)
	bizGetUserByID = bizerr.NewBiz(ServiceName, "GetUserByID", baseCode+bizCodeGetUserByID)
	bizListSecret = bizerr.NewBiz(ServiceName, "ListSecret", baseCode+bizCodeListSecret)
	bizLogin = bizerr.NewBiz(ServiceName, "Login", baseCode+bizCodeLogin)
	bizRegister = bizerr.NewBiz(ServiceName, "Register", baseCode+bizCodeRegister)
	bizUpdateSecret = bizerr.NewBiz(ServiceName, "UpdateSecret", baseCode+bizCodeUpdateSecret)
	bizUpdateUser = bizerr.NewBiz(ServiceName, "UpdateUser", baseCode+bizCodeUpdateUser)
}
