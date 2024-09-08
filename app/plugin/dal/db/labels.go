package db

import (
	"context"
	"fmt"
	"math"

	"github.com/aiagt/aiagt/app/plugin/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type LabelDao struct {
	m *model.PluginLabel
}

// NewLabelDao make plugin label dao
func NewLabelDao() *LabelDao {
	return &LabelDao{m: new(model.PluginLabel)}
}

func (d *LabelDao) db(ctx context.Context) *gorm.DB {
	return ctxutil.Tx(ctx)
}

// GetByID get plugin label by id
func (d *LabelDao) GetByID(ctx context.Context, id int64) (*model.PluginLabel, error) {
	var result model.PluginLabel

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin label dao get by id error")
	}

	return &result, nil
}

// GetByIDs get plugin label list by ids
func (d *LabelDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.PluginLabel, error) {
	var result []*model.PluginLabel

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "plugin label dao get by ids error")
	}

	return result, nil
}

// List get plugin label list
func (d *LabelDao) List(ctx context.Context, req *pluginsvc.ListPluginLabelReq) ([]*model.PluginLabel, *base.PaginationResp, error) {
	var (
		list   []*model.PluginLabel
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		if req.Text != nil {
			db = db.Where("text LIKE ?", fmt.Sprintf("%%%s%%", *req.Text))
		}
		return db
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "plugin label dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a plugin label record
func (d *LabelDao) Create(ctx context.Context, m *model.PluginLabel) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "plugin label dao create error")
	}

	return nil
}

func (d *LabelDao) CreateBatch(ctx context.Context, labels []*model.PluginLabel) error {
	err := d.db(ctx).Model(d.m).CreateInBatches(&labels, 100).Error
	if err != nil {
		return errors.Wrap(err, "plugin label dao create batch error")
	}
	return nil
}

// Delete delete plugin label by id
func (d *LabelDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "plugin label dao delete error")
	}

	return nil
}

func (d *LabelDao) UpdateLabels(ctx context.Context, labelIDs []int64, labelTexts []string) ([]int64, error) {
	if labelTexts == nil {
		return labelIDs, nil
	}

	labels := make([]*model.PluginLabel, len(labelTexts))
	for i, text := range labelTexts {
		labels[i] = &model.PluginLabel{Text: text}
	}
	if err := d.CreateBatch(ctx, labels); err != nil {
		return nil, errors.Wrap(err, "plugin label dao update labels error")
	}

	for _, label := range labels {
		labelIDs = append(labelIDs, label.ID)
	}

	return labelIDs, nil
}
