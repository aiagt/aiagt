package ctxutil

import (
	"context"

	"gorm.io/gorm"
)

const (
	TxKey = "TX"
)

func WithTx(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, TxKey, db)
}

func Tx(ctx context.Context) *gorm.DB {
	tx, _ := ctx.Value(TxKey).(*gorm.DB)
	return tx
}
