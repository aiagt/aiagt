package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func Auth(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		return next(ctx, req, resp)
	}
}
