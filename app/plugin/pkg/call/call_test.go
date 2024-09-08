package call

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stretchr/testify/assert"
)

const addr = "localhost:34378"

type (
	ReqBodyType struct {
		Message string `json:"message"`
	}
	ReqType struct {
		PluginID       int64             `json:"plugin_id"`
		ToolID         int64             `json:"tool_id"`
		UserID         int64             `json:"user_id"`
		Secrets        map[string]string `json:"secrets"`
		ModelCallToken *string           `json:"model_call_token,omitempty"`
		ModelCallLimit uint              `json:"model_call_limit,omitempty"`
		Body           *ReqBodyType      `json:"body,omitempty" binding:"required"`
	}
)

func runServer() {
	h := server.Default(server.WithHostPorts(addr))

	h.POST("/test", func(ctx context.Context, c *app.RequestContext) {
		var r ReqType
		err := c.BindAndValidate(&r)
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{"code": -1, "msg": err.Error()})
			return
		}

		c.JSON(consts.StatusOK, utils.H{"code": 0, "msg": "ok", "data": r.Body.Message})
	})

	h.Spin()
}

func TestCall(t *testing.T) {
	go runServer()

	var (
		reqType = &RequestType{ContentType: "application/json", Parameters: Object{
			&Field{Name: "message", Type: "string"},
		}}
		respType = &ResponseType{ContentType: "application/json", Parameters: Object{
			&Field{Name: "code", Type: "number"},
			&Field{Name: "msg", Type: "string"},
		}}
		body   = []byte(`{"message":"success"}`)
		apiURL = fmt.Sprintf("http://%s/test", addr)
	)

	resp, err := Call(context.Background(), new(RequestBody), apiURL, reqType, respType, body)
	assert.NoError(t, err)

	t.Log(string(resp))
}
