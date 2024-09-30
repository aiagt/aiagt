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
		const defaultDBName = ""

		// start transaction
		tx := ktdb.DB().Begin()
		ctx = ctxutil.WithTx(ctx, tx)

		err = next(ctx, req, resp)
		if err != nil {
			// if error occurs, rollback transaction
			tx.Rollback()
			return err
		}

		// commit transaction
		err = tx.Commit().Error
		if err != nil {
			return ReturnBizErr(ctx, errors.Wrap(err, "transaction commit failed"))
		}

		return
	}
}
