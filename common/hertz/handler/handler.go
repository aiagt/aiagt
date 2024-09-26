package handler

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/common/hertz/result"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/hertz-contrib/sse"
	"github.com/pkg/errors"
)

// PinPongHandle pin pong api
type PinPongHandle[Q, P any] func(context.Context, *Q, ...callopt.Option) (P, error)

func PinPongHandler[Q, P any](handle PinPongHandle[Q, P]) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req Q

		if err := c.BindAndValidate(&req); err != nil {
			c.JSON(consts.StatusOK, result.Error(bizerr.ErrCodeBadRequest, err))
			return
		}

		authorization := c.Request.Header.Get("Authorization")
		if len(authorization) > 0 {
			token := strings.TrimPrefix(authorization, "Bearer ")
			ctx = ctxutil.WithToken(ctx, token)
		}

		resp, err := handle(ctx, &req)
		if err != nil {
			hlog.CtxErrorf(ctx, err.Error())
			c.JSON(consts.StatusOK, result.BizError(err))

			return
		}

		c.JSON(consts.StatusOK, result.Success(resp))
	}
}

// NoReqPinPongHandle no request parameters pin pong api
type NoReqPinPongHandle[P any] func(context.Context, ...callopt.Option) (P, error)

func NoReqPinPongHandler[P any](h NoReqPinPongHandle[P]) PinPongHandle[struct{}, P] {
	return func(ctx context.Context, _ *struct{}, options ...callopt.Option) (P, error) {
		return h(ctx, options...)
	}
}

// ServerStreamingHandle server streaming api
type ServerStreamingHandle[Q, P any, S ServerStreamingClient[P]] func(ctx context.Context, req *Q, callOptions ...streamcall.Option) (stream S, err error)

func SererStreamingHandler[Q, P any, S ServerStreamingClient[P]](handle ServerStreamingHandle[Q, P, S]) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var (
			req    Q
			resp   P
			stream = sse.NewStream(c)
			chunk  []byte
		)

		if err := c.BindAndValidate(&req); err != nil {
			result.StreamError(ctx, stream, bizerr.ErrCodeBadRequest, err)
			return
		}

		authorization := c.Request.Header.Get("Authorization")
		if len(authorization) > 0 {
			token := strings.TrimPrefix(authorization, "Bearer ")
			ctx = ctxutil.WithToken(ctx, token)
		}

		respStream, err := handle(ctx, &req)
		if err != nil {
			hlog.CtxErrorf(ctx, err.Error())
			result.StreamBizError(ctx, stream, err)

			return
		}

		for {
			resp, err = respStream.Recv()

			if errors.Is(err, io.EOF) {
				result.Stream(ctx, stream, "done", nil)
				return
			}

			if err != nil {
				hlog.CtxErrorf(ctx, err.Error())
				result.StreamBizError(ctx, stream, err)

				return
			}

			chunk, err = json.Marshal(resp)
			if err != nil {
				hlog.CtxWarnf(ctx, err.Error())
			}

			result.Stream(ctx, stream, "chunk", chunk)
		}
	}
}

type ServerStreamingClient[P any] interface {
	Recv() (P, error)
}
