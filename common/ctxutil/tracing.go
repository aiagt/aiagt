package ctxutil

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

const (
	currentSpanKey CtxKey = "CURRENT_SPAN_KEY"
)

func WithSpan(ctx context.Context, span trace.Span) context.Context {
	return WithMapValue(ctx, currentSpanKey, span)
}

func ApplySpan(ctx context.Context) context.Context {
	span, ok := GetMapValue[trace.Span](ctx, currentSpanKey)
	if ok {
		ctx = trace.ContextWithSpan(ctx, span)
	}

	return ctx
}
