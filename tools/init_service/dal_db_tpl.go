package main

const dalDBTpl = `package db

import (
    "context"
    "math"

    "github.com/aiagt/aiagt/apps/{{ .Service.Name }}/model"
    "github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/pkg/snowflake"
    "github.com/pkg/errors"

    ktdb "github.com/aiagt/kitextool/option/server/db"
    "gorm.io/gorm"
)

type {{ .Model.Camel }}Dao struct {
    m *model.{{ .Model.Camel }}
}

// New{{ .Model.Camel }}Dao make {{ .Model.Camel }} dao
func New{{ .Model.Camel }}Dao() *{{ .Model.Camel }}Dao {
    return &{{ .Model.Camel }}Dao{m: new(model.{{ .Model.Camel }})}
}

func (d *{{ .Model.Camel }}Dao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get {{ .Model.Snake }} by id
func (d *{{ .Model.Camel }}Dao) GetByID(ctx context.Context, id int64) (*model.{{ .Model.Camel }}, error) {
    var result model.{{ .Model.Camel }}

    err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
    if err != nil {
        return nil, errors.Wrap(err, "{{ .Model.Snake }} dao get by id error")
    }

    return &result, nil
}

// GetByIDs get {{ .Model.Snake }} list by ids
func (d *{{ .Model.Camel }}Dao) GetByIDs(ctx context.Context, ids []int64) ([]*model.{{ .Model.Camel }}, error) {
    var result []*model.{{ .Model.Camel }}

    err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
    if err != nil {
        return nil, errors.Wrap(err, "{{ .Model.Snake }} dao get by ids error")
    }

    return result, nil
}

// List get {{ .Model.Snake }} list
func (d *{{ .Model.Camel }}Dao) List(ctx context.Context, page *base.PaginationReq) ([]*model.{{ .Model.Camel }}, *base.PaginationResp, error) {
    var (
        list   []*model.{{ .Model.Camel }}
        total  int64
        offset = int((page.Page-1)*page.PageSize)
        limit  = int(page.PageSize)
    )

    err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
    if err != nil {
        return nil, nil, errors.Wrap(err, "{{ .Model.Snake }} dao get page error")
    }

    pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
    pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

    return list, pageResp, nil
}

// Create insert a {{ .Model.Snake }} record
func (d *{{ .Model.Camel }}Dao) Create(ctx context.Context, m *model.{{ .Model.Camel }}) error {
	m.ID = snowflake.Generate().Int64()

    err := d.db(ctx).Model(d.m).Create(m).Error
    if err != nil {
        return errors.Wrap(err, "{{ .Model.Snake }} dao create error")
    }

    return nil
}

// Update {{ .Model.Snake }} by id
func (d *{{ .Model.Camel }}Dao) Update(ctx context.Context, id int64, m *model.{{ .Model.Camel }}Optional) error {
    err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
    if err != nil {
        return errors.Wrap(err, "{{ .Model.Snake }} dao update error")
    }

    return nil
}

// Delete delete {{ .Model.Snake }} by id
func (d *{{ .Model.Camel }}Dao) Delete(ctx context.Context, id int64) error {
    err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
    if err != nil {
        return errors.Wrap(err, "{{ .Model.Snake }} dao delete error")
    }

    return nil
}
`

var DalDBTpl = NewTemplate("dal.db", dalDBTpl, false)
