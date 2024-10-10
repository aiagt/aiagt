package ctxutil

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"go.opentelemetry.io/otel/trace"
	"unsafe"
)

const currentSpanKey CtxKey = "currentSpanKey"
const parentSpanIDKey = "PARENT_SPAN_ID"
const spanIDHeader = "OT_TRACER_SPANID"
const smallSpanIDHeader = "ot-tracer-spanid"

func WithSpan(ctx context.Context, span trace.Span) context.Context {
	return WithMapValue(ctx, currentSpanKey, span)
}

func Span(ctx context.Context) trace.Span {
	span, _ := GetMapValue[trace.Span](ctx, currentSpanKey)
	return span
}

func WithParentSpanID(ctx context.Context) context.Context {
	spanCtx := trace.SpanContextFromContext(ctx)
	return metainfo.WithPersistentValue(ctx, parentSpanIDKey, bytes2Hex(spanCtx.SpanID()))
}

func ResetParentSpanID(ctx context.Context) context.Context {
	ctx = metainfo.WithValue(ctx, spanIDHeader, ParentSpanID(ctx))
	ctx = metainfo.WithValue(ctx, smallSpanIDHeader, ParentSpanID(ctx))
	return ctx
}

func ParentSpanID(ctx context.Context) string {
	spanID, _ := metainfo.GetPersistentValue(ctx, parentSpanIDKey)
	return spanID
}

func bytes2Hex(byteArray [8]byte) string {
	const hexCharset = "0123456789abcdef"

	hexArray := make([]byte, 2*len(byteArray))

	for i, b := range byteArray {
		hexArray[i*2] = hexCharset[b>>4]     // High 4 bits
		hexArray[i*2+1] = hexCharset[b&0x0F] // Lower 4 bits
	}

	return *(*string)(unsafe.Pointer(&hexArray))
}
