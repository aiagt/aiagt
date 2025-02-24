package db

import (
	"context"
	"math"

	"github.com/aiagt/aiagt/apps/workflow/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"github.com/pkg/errors"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type WorkflowDao struct {
	m *model.Workflow
}

// NewWorkflowDao make Workflow dao
func NewWorkflowDao() *WorkflowDao {
	return &WorkflowDao{m: new(model.Workflow)}
}

func (d *WorkflowDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get workflow by id
func (d *WorkflowDao) GetByID(ctx context.Context, id int64) (*model.Workflow, error) {
	var result model.Workflow

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "workflow dao get by id error")
	}

	return &result, nil
}

// GetByIDs get workflow list by ids
func (d *WorkflowDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Workflow, error) {
	var result []*model.Workflow

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "workflow dao get by ids error")
	}

	return result, nil
}

// List get workflow list
func (d *WorkflowDao) List(ctx context.Context, page *base.PaginationReq) ([]*model.Workflow, *base.PaginationResp, error) {
	var (
		list   []*model.Workflow
		total  int64
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "workflow dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a workflow record
func (d *WorkflowDao) Create(ctx context.Context, m *model.Workflow) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "workflow dao create error")
	}

	return nil
}

// Update workflow by id
func (d *WorkflowDao) Update(ctx context.Context, id int64, m *model.WorkflowOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "workflow dao update error")
	}

	return nil
}

// Delete delete workflow by id
func (d *WorkflowDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "workflow dao delete error")
	}

	return nil
}
