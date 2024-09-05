package db

import (
	"context"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type PluginDao[T any] struct {
	m *T
}

func NewPluginDao[T any](m *T) *PluginDao[T] {
	return &PluginDao[T]{m: m}
}

func (dao *PluginDao[T]) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

func (dao *PluginDao[T]) mdb(ctx context.Context) *gorm.DB {
	return dao.db(ctx).Model(dao.m)
}

func (dao *PluginDao[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	var result T

	err := dao.mdb(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
