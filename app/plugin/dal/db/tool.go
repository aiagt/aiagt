package db

import (
	"context"
	"math"

	"github.com/aiagt/aiagt/app/plugin/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ToolDao struct {
	m *model.PluginTool
}

// NewToolDao make plugin tool dao
func NewToolDao() *ToolDao {
	return &ToolDao{m: new(model.PluginTool)}
}

func (d *ToolDao) db(ctx context.Context) *gorm.DB {
	return ctxutil.Tx(ctx)
}

// GetByID get plugin tool by id
func (d *ToolDao) GetByID(ctx context.Context, id int64) (*model.PluginTool, error) {
	var result model.PluginTool

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin tool dao get by id error")
	}

	return &result, nil
}

// GetByIDs get plugin tool list by ids
func (d *ToolDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.PluginTool, error) {
	var result []*model.PluginTool

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin tool dao get by ids error")
	}

	return result, nil
}

// GetByPluginID get plugin tools by plugin_id
func (d *ToolDao) GetByPluginID(ctx context.Context, pluginID int64) ([]*model.PluginTool, error) {
	var result []*model.PluginTool

	err := d.db(ctx).Model(d.m).Where("plugin_id = ?", pluginID).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin tool dao get by plugin id error")
	}

	return result, nil
}

// List get plugin tool list
func (d *ToolDao) List(ctx context.Context, req *pluginsvc.ListPluginToolReq, userID int64) ([]*model.PluginTool, *base.PaginationResp, error) {
	var (
		list   []*model.PluginTool
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		db = db.Joins("LEFT JOIN plugins ON plugin_id = plugins.id").
			Where("plugins.author_id = ? OR plugins.is_private = ?", userID, false)
		if req.PluginId != nil {
			db = db.Where("plugin_id = ?", req.PluginId)
		}
		if req.ToolIds != nil {
			db = db.Where("plugin_tools.id in ?", req.ToolIds)
		}
		return db
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "plugin tool dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a plugin tool record
func (d *ToolDao) Create(ctx context.Context, m *model.PluginTool) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "plugin tool dao create error")
	}

	return nil
}

// Update plugin tool by id
func (d *ToolDao) Update(ctx context.Context, id int64, m *model.PluginToolOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "plugin tool dao update error")
	}

	return nil
}

// Delete delete plugin tool by id
func (d *ToolDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "plugin tool dao delete error")
	}

	return nil
}
