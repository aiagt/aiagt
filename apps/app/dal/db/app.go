package db

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"math"

	ktdb "github.com/aiagt/kitextool/option/server/db"

	"github.com/aiagt/aiagt/apps/app/model"
	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type AppDao struct {
	m *model.App
}

// NewAppDao make App dao
func NewAppDao() *AppDao {
	return &AppDao{m: new(model.App)}
}

func (d *AppDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get app by id
func (d *AppDao) GetByID(ctx context.Context, id int64) (*model.App, error) {
	var result model.App

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "app dao get by id error")
	}

	return &result, nil
}

// GetByIDs get app list by ids
func (d *AppDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.App, error) {
	var result []*model.App

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "app dao get by ids error")
	}

	return result, nil
}

// List get app list
func (d *AppDao) List(ctx context.Context, req *appsvc.ListAppReq, userID int64) ([]*model.App, *base.PaginationResp, error) {
	var (
		list   []*model.App
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		db = db.Where("author_id = ? OR is_private = ?", userID, false)
		if req.AuthorId != nil {
			db = db.Where("author_id = ?", req.AuthorId)
		}
		if req.Name != nil {
			db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *req.Name))
		}
		if req.Description != nil {
			db = db.Where("description LIKE ?", fmt.Sprintf("%%%s%%", *req.Description))
		}
		if req.Labels != nil {
			db = db.Where("label_ids IN ?", req.Labels)
		}
		return db
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "app dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a app record
func (d *AppDao) Create(ctx context.Context, m *model.App) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "app dao create error")
	}

	return nil
}

// Update app by id
func (d *AppDao) Update(ctx context.Context, id int64, m *model.AppOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "app dao update error")
	}

	return nil
}

// Delete delete app by id
func (d *AppDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "app dao delete error")
	}

	return nil
}
