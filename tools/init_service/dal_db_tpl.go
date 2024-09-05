package main

import "text/template"

// TODO: support transaction

const daoDBTpl = `package db

import (
	"context"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type {{ .CamelServiceName }}Dao[T any] struct {
	m *T
}

func New{{ .CamelServiceName }}Dao[T any](m *T) *{{ .CamelServiceName }}Dao[T] {
	return &{{ .CamelServiceName }}Dao[T]{m: m}
}

func (dao *{{ .CamelServiceName }}Dao[T]) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

func (dao *{{ .CamelServiceName }}Dao[T]) mdb(ctx context.Context) *gorm.DB {
	return dao.db(ctx).Model(dao.m)
}

func (dao *{{ .CamelServiceName }}Dao[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	var result T

	err := dao.mdb(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
`

var DaoDBTpl = template.Must(template.New("db.go").Parse(daoDBTpl))
