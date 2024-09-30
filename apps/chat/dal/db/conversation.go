package db

import (
	"context"
	"math"

	"github.com/aiagt/aiagt/apps/chat/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/chatsvc"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type ConversationDao struct {
	m *model.Conversation
}

// NewConversationDao make Conversation dao
func NewConversationDao() *ConversationDao {
	return &ConversationDao{m: new(model.Conversation)}
}

func (d *ConversationDao) db(ctx context.Context) *gorm.DB {
	return ctxutil.Tx(ctx)
}

// GetByID get conversation by id
func (d *ConversationDao) GetByID(ctx context.Context, id int64) (*model.Conversation, error) {
	var result model.Conversation

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "conversation dao get by id error")
	}

	return &result, nil
}

// GetByIDs get conversation list by ids
func (d *ConversationDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Conversation, error) {
	var result []*model.Conversation

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "conversation dao get by ids error")
	}

	return result, nil
}

// List get conversation list
func (d *ConversationDao) List(ctx context.Context, req *chatsvc.ListConversationReq, userID int64) ([]*model.Conversation, *base.PaginationResp, error) {
	var (
		list   []*model.Conversation
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		db = db.Where("app_id = ? AND user_id = ?", req.AppId, userID)
		return db
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "conversation dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a conversation record
func (d *ConversationDao) Create(ctx context.Context, m *model.Conversation) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "conversation dao create error")
	}

	return nil
}

// Update conversation by id
func (d *ConversationDao) Update(ctx context.Context, id int64, m *model.ConversationOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "conversation dao update error")
	}

	return nil
}

// Delete delete conversation by id
func (d *ConversationDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "conversation dao delete error")
	}

	return nil
}
