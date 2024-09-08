package middleware

import (
	"context"
	"github.com/aiagt/aiagt/common/ctxutil"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/pkg/errors"
)

func (m *Middleware) Transaction(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		tx := ktdb.DBCtx(ctx).Begin()
		ctx = ctxutil.WithTx(ctx, tx)

		err = next(ctx, req, resp)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Commit().Error
		if err != nil {
			return ReturnBizErr(ctx, errors.Wrap(err, "transaction commit failed"))
		}

		return
	}
}
