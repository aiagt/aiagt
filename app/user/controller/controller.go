package controller

import (
	"github.com/aiagt/aiagt/common/hertz/handler"
	"github.com/aiagt/aiagt/common/hertz/router"
	"github.com/aiagt/aiagt/kitex_gen/usersvc/userservice"
	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterRouter(r *route.RouterGroup, cli userservice.Client) {
	r = r.Group("/user")

	router.POST(r, "/register", cli.Register)
	router.POST(r, "/login", cli.Login)
	router.POST(r, "/captcha", cli.SendCaptcha)
	router.POST(r, "/reset-password", cli.ResetPassword)

	router.GET(r, "/current", handler.NoReqPinPongHandler(cli.GetUser))
	router.GET(r, "/:id", cli.GetUserByID)
	router.POST(r, "/", cli.GetUserByIds)
	router.PUT(r, "/", cli.UpdateUser)

	router.POST(r, "/secret", cli.CreateSecret)
	router.PUT(r, "/secret/:id", cli.UpdateSecret)
	router.DELETE(r, "/secret/:id", cli.DeleteSecret)
	router.GET(r, "/secret", cli.ListSecret)
}
