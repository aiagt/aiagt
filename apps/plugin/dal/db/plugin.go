package db

import (
	"context"
	"fmt"
	"math"
	"strings"

	ktdb "github.com/aiagt/kitextool/option/server/db"

	"github.com/aiagt/aiagt/apps/plugin/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/pkg/errors"

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

// GetByIDs get plugin list by ids
func (d *PluginDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Plugin, error) {
	var result []*model.Plugin

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin dao get by ids error")
	}

	return result, nil
}

// List get plugin list
func (d *PluginDao) List(ctx context.Context, req *pluginsvc.ListPluginReq, userID int64) ([]*model.Plugin, *base.PaginationResp, error) {
	var (
		list   []*model.Plugin
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		if req.AuthorId == nil {
			db = db.Where("author_id = ? OR is_private = ?", userID, false)
		} else {
			if *req.AuthorId != userID {
				db = db.Where("author_id = ? AND is_private = ?", *req.AuthorId, false)
			} else {
				db = db.Where("author_id = ?", *req.AuthorId)
			}
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
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "plugin dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

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

// Update plugin by id
func (d *PluginDao) Update(ctx context.Context, id int64, m *model.PluginOptional) error {
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

// GetByKey get plugin by key
func (d *PluginDao) GetByKey(ctx context.Context, key int64) (*model.Plugin, error) {
	var result model.Plugin

	err := d.db(ctx).Model(d.m).Where("`key` = ?", key).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin dao get by key error")
	}

	return &result, nil
}
