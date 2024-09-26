package result

import (
	"context"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hertz-contrib/sse"
)

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

var biz = bizerr.NewBiz("gateway", "handler", 10000)

func Error(code bizerr.ErrCode, err error) *Response {
	bizErr := biz.NewCodeErr(code, err)
	return &Response{Code: bizErr.BizStatusCode(), Msg: bizErr.BizMessage()}
}

func BizError(err error) *Response {
	bizErr, ok := kerrors.FromBizStatusError(err)
	if !ok {
		bizErr = biz.NewErr(err)
	}

	return &Response{Code: bizErr.BizStatusCode(), Msg: bizErr.BizMessage()}
}

func Success(data interface{}) *Response {
	resp := &Response{Code: 0, Msg: "success"}
	if data != nil {
		resp.Data = data
	}
	return resp
}

func Stream(ctx context.Context, stream *sse.Stream, event string, data []byte) {
	err := stream.Publish(&sse.Event{
		Event: event,
		Data:  data,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
	}
}

func StreamError(ctx context.Context, stream *sse.Stream, code bizerr.ErrCode, err error) {
	Stream(ctx, stream, "error", []byte(biz.NewCodeErr(code, err).BizMessage()))
}

func StreamBizError(ctx context.Context, stream *sse.Stream, err error) {
	bizErr, ok := kerrors.FromBizStatusError(err)
	if !ok {
		bizErr = biz.NewCodeErr(bizerr.ErrCodeServerFailure, err)
	}

	Stream(ctx, stream, "error", []byte(bizErr.BizMessage()))
}
