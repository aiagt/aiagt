package db

import (
	"context"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

func db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

type AppDao[T any] struct{}

func (d *AppDao[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	var m T

	if err := db(ctx).Model(&m).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

