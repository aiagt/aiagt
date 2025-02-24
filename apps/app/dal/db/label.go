package db

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/pkg/lists"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"github.com/aiagt/aiagt/pkg/utils"
	"math"

	ktdb "github.com/aiagt/kitextool/option/server/db"

	"github.com/aiagt/aiagt/apps/app/model"
	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type LabelDao struct {
	m *model.AppLabel
}

// NewLabelDao make AppLabel dao
func NewLabelDao() *LabelDao {
	return &LabelDao{m: new(model.AppLabel)}
}

func (d *LabelDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get label by id
func (d *LabelDao) GetByID(ctx context.Context, id int64) (*model.AppLabel, error) {
	var result model.AppLabel

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "label dao get by id error")
	}

	return &result, nil
}

// GetByIDs get label list by ids
func (d *LabelDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.AppLabel, error) {
	var result []*model.AppLabel

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "label dao get by ids error")
	}

	return result, nil
}

// List get label list
func (d *LabelDao) List(ctx context.Context, req *appsvc.ListAppLabelReq) ([]*model.AppLabel, *base.PaginationResp, error) {
	var (
		list   []*model.AppLabel
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		if req.Text != nil {
			db = db.Where("text like ?", fmt.Sprintf("%%%s%%", *req.Text))
		}
		if utils.ValOf(req.Pinned) {
			db = db.Where("pinned > ?", 0)
		}
		return db
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "label dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a label record
func (d *LabelDao) Create(ctx context.Context, m *model.AppLabel) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "label dao create error")
	}

	return nil
}

// Delete delete label by id
func (d *LabelDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "label dao delete error")
	}

	return nil
}

func (d *LabelDao) CreateBatch(ctx context.Context, labels []*model.AppLabel) error {
	labels = lists.Map(labels, func(t *model.AppLabel) *model.AppLabel {
		t.ID = snowflake.Generate().Int64()
		return t
	})

	err := d.db(ctx).Model(d.m).CreateInBatches(&labels, 100).Error
	if err != nil {
		return errors.Wrap(err, "app label dao create batch error")
	}

	return nil
}

func (d *LabelDao) UpdateLabels(ctx context.Context, labelIDs []int64, labelTexts []string) ([]int64, error) {
	if labelTexts == nil {
		return labelIDs, nil
	}

	labels := lists.Map(labelTexts, func(t string) *model.AppLabel {
		return &model.AppLabel{Text: t}
	})

	if err := d.CreateBatch(ctx, labels); err != nil {
		return nil, errors.Wrap(err, "app label dao update labels error")
	}

	for _, label := range labels {
		labelIDs = append(labelIDs, label.ID)
	}

	return labelIDs, nil
}
