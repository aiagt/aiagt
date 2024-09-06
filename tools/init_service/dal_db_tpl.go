package main

const dalDBTpl = `package db

import (
	"context"

	"github.com/aiagt/aiagt/app/{{ .ServiceName }}/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/pkg/errors"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type {{ .CamelServiceName }}Dao struct {
	m *model.{{ .CamelServiceName }}
}

// New{{ .CamelServiceName }}Dao make {{ .SnakeServiceName }} dao
func New{{ .CamelServiceName }}Dao() *{{ .CamelServiceName }}Dao {
	return &{{ .CamelServiceName }}Dao{m: new(model.{{ .CamelServiceName }})}
}

func (d *{{ .CamelServiceName }}Dao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get {{ .SnakeServiceName }} by id
func (d *{{ .CamelServiceName }}Dao) GetByID(ctx context.Context, id int64) (*model.{{ .CamelServiceName }}, error) {
	var result model.{{ .CamelServiceName }}

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "{{ .SnakeServiceName }} dao get by id error")
	}

	return &result, nil
}

// List get {{ .SnakeServiceName }} list
func (d *{{ .CamelServiceName }}Dao) List(ctx context.Context, page *base.PaginationReq) ([]*model.{{ .CamelServiceName }}, *base.PaginationResp, error) {
	var (
		list   []*model.{{ .CamelServiceName }}
		total  int64
		offset = int((page.Page-1)*page.PageSize)
		limit  = int(page.PageSize)
	)


	err := d.db(ctx).Model(d.m).Offset(offset).Limit(limit).Find(&list).Count(&total).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "{{ .SnakeServiceName }} dao get page error")
	}

	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: int32(total) / page.PageSize}
	return list, pageResp, nil
}

// Create insert a {{ .SnakeServiceName }} record
func (d *{{ .CamelServiceName }}Dao) Create(ctx context.Context, m *model.{{ .CamelServiceName }}) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "{{ .SnakeServiceName }} dao create error")
	}

	return nil
}

// Update update {{ .SnakeServiceName }} by id
func (d *{{ .CamelServiceName }}Dao) Update(ctx context.Context, id int64, m *model.{{ .CamelServiceName }}) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "{{ .SnakeServiceName }} dao update error")
	}

	return nil
}

// Delete delete {{ .SnakeServiceName }} by id
func (d *{{ .CamelServiceName }}Dao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "{{ .SnakeServiceName }} dao delete error")
	}

	return nil
}
`

var DalDBTpl = NewTemplate("dal.db", dalDBTpl, false)
