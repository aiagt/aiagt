package router

import (
	"github.com/aiagt/aiagt/common/hertz/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func GET[Q, P any](r *route.RouterGroup, relativePath string, handle handler.PinPongHandle[Q, P]) {
	r.GET(relativePath, handler.PinPongHandler(handle))
}

func POST[Q, P any](r *route.RouterGroup, relativePath string, handle handler.PinPongHandle[Q, P]) {
	r.POST(relativePath, handler.PinPongHandler(handle))
}

func PUT[Q, P any](r *route.RouterGroup, relativePath string, handle handler.PinPongHandle[Q, P]) {
	r.PUT(relativePath, handler.PinPongHandler(handle))
}

func DELETE[Q, P any](r *route.RouterGroup, relativePath string, handle handler.PinPongHandle[Q, P]) {
	r.DELETE(relativePath, handler.PinPongHandler(handle))
}

func SSE[Q, P any, S handler.ServerStreamingClient[P]](r *route.RouterGroup, relativePath string, handle handler.ServerStreamingHandle[Q, P, S]) {
	r.POST(relativePath, handler.SererStreamingHandler(handle))
}
