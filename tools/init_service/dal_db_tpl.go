package main

import "text/template"

const daoDBTpl = `package db

import (
	"context"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

func db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

type {{ .CamelServiceName }}Dao[T any] struct{}

func (d *{{ .CamelServiceName }}Dao[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	var m T

	if err := db(ctx).Model(&m).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}
`

var DaoDBTpl = template.Must(template.New("db.go").Parse(daoDBTpl))
