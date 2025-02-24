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

type WorkflowNodeDao struct {
	m *model.WorkflowNode
}

// NewWorkflowNodeDao make WorkflowNode dao
func NewWorkflowNodeDao() *WorkflowNodeDao {
	return &WorkflowNodeDao{m: new(model.WorkflowNode)}
}

func (d *WorkflowNodeDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get workflow_node by id
func (d *WorkflowNodeDao) GetByID(ctx context.Context, id int64) (*model.WorkflowNode, error) {
	var result model.WorkflowNode

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "workflow_node dao get by id error")
	}

	return &result, nil
}

// GetByIDs get workflow_node list by ids
func (d *WorkflowNodeDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.WorkflowNode, error) {
	var result []*model.WorkflowNode

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "workflow_node dao get by ids error")
	}

	return result, nil
}

// List get workflow_node list
func (d *WorkflowNodeDao) List(ctx context.Context, page *base.PaginationReq) ([]*model.WorkflowNode, *base.PaginationResp, error) {
	var (
		list   []*model.WorkflowNode
		total  int64
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "workflow_node dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a workflow_node record
func (d *WorkflowNodeDao) Create(ctx context.Context, m *model.WorkflowNode) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "workflow_node dao create error")
	}

	return nil
}

// Update workflow_node by id
func (d *WorkflowNodeDao) Update(ctx context.Context, id int64, m *model.WorkflowNodeOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "workflow_node dao update error")
	}

	return nil
}

// Delete delete workflow_node by id
func (d *WorkflowNodeDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "workflow_node dao delete error")
	}

	return nil
}

// GetByWorkflowID get workflow_node by workflow id
func (d *WorkflowNodeDao) GetByWorkflowID(ctx context.Context, workflowID int64) ([]*model.WorkflowNode, error) {
	var result []*model.WorkflowNode

	err := d.db(ctx).Model(d.m).Where("workflow_id = ?", workflowID).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "workflow_node dao get by workflow id error")
	}

	return result, nil
}
