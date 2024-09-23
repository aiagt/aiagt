package db

import (
	"context"
	"math"

	"github.com/aiagt/aiagt/app/model/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type ModelDao struct {
	m *model.Models
}

// NewModelDao make Models dao
func NewModelDao() *ModelDao {
	return &ModelDao{m: new(model.Models)}
}

func (d *ModelDao) db(ctx context.Context) *gorm.DB {
	return ctxutil.Tx(ctx)
}

// GetByID get model by id
func (d *ModelDao) GetByID(ctx context.Context, id int64) (*model.Models, error) {
	var result model.Models

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "model dao get by id error")
	}

	return &result, nil
}

// GetByIDs get model list by ids
func (d *ModelDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Models, error) {
	var result []*model.Models

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "model dao get by ids error")
	}

	return result, nil
}

// List get model list
func (d *ModelDao) List(ctx context.Context, page *base.PaginationReq) ([]*model.Models, *base.PaginationResp, error) {
	var (
		list   []*model.Models
		total  int64
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "model dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a model record
func (d *ModelDao) Create(ctx context.Context, m *model.Models) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "model dao create error")
	}

	return nil
}

// Update model by id
func (d *ModelDao) Update(ctx context.Context, id int64, m *model.ModelsOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "model dao update error")
	}

	return nil
}

// Delete delete model by id
func (d *ModelDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "model dao delete error")
	}

	return nil
}
