package ctxutil

import (
	"context"

	"gorm.io/gorm"
)

const (
	TxKey CtxKey = "TRANSACTION"
)

func WithTx(ctx context.Context, db *gorm.DB) context.Context {
	if IsStreaming(ctx) {
		return WithMapValue(ctx, TxKey, db)
	}

	return context.WithValue(ctx, TxKey, db)
}

func Tx(ctx context.Context) (tx *gorm.DB) {
	if IsStreaming(ctx) {
		tx, _ = GetMapValue[*gorm.DB](ctx, TxKey)
	} else {
		tx, _ = ctx.Value(TxKey).(*gorm.DB)
	}

	return tx
}
