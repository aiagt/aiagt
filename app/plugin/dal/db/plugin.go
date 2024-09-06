package db

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/app/plugin/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/pkg/errors"
	"strings"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type PluginDao struct {
	m *model.Plugin
}

// NewPluginDao make plugin dao
func NewPluginDao() *PluginDao {
	return &PluginDao{m: new(model.Plugin)}
}

func (d *PluginDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get plugin by id
func (d *PluginDao) GetByID(ctx context.Context, id int64) (*model.Plugin, error) {
	var result model.Plugin

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin dao get by id error")
	}

	return &result, nil
}

// List get plugin list
func (d *PluginDao) List(ctx context.Context, req *pluginsvc.ListPluginReq) ([]*model.Plugin, *base.PaginationResp, error) {
	var (
		list   []*model.Plugin
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
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
			labels := strings.ReplaceAll(fmt.Sprint(req.Labels), " ", ",")
			db = db.Where("JSON_OVERLAPS(label_ids, ?)", labels)
		}
		return db
	}).Offset(offset).Limit(limit).Find(&list).Count(&total).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "plugin dao get page error")
	}

	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: int32(total) / page.PageSize}
	return list, pageResp, nil
}

// Create insert a plugin record
func (d *PluginDao) Create(ctx context.Context, m *model.Plugin) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "plugin dao create error")
	}

	return nil
}

// Update update plugin by id
func (d *PluginDao) Update(ctx context.Context, id int64, m *model.Plugin) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "plugin dao update error")
	}

	return nil
}

// Delete delete plugin by id
func (d *PluginDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "plugin dao delete error")
	}

	return nil
}
