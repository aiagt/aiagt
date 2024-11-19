package db

import (
	"context"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	"math"

	"github.com/aiagt/aiagt/apps/chat/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/kitex_gen/chatsvc"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type MessageDao struct {
	m *model.Message
}

// NewMessageDao make Message dao
func NewMessageDao() *MessageDao {
	return &MessageDao{m: new(model.Message)}
}

func (d *MessageDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get message by id
func (d *MessageDao) GetByID(ctx context.Context, id int64) (*model.Message, error) {
	var result model.Message

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "message dao get by id error")
	}

	return &result, nil
}

// GetByIDs get message list by ids
func (d *MessageDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Message, error) {
	var result []*model.Message

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "message dao get by ids error")
	}

	return result, nil
}

// List get message list
func (d *MessageDao) List(ctx context.Context, req *chatsvc.ListMessageReq) ([]*model.Message, *base.PaginationResp, error) {
	var (
		list   []*model.Message
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Where("conversation_id = ?", req.ConversationId).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "message dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a message record
func (d *MessageDao) Create(ctx context.Context, m *model.Message) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "message dao create error")
	}

	return nil
}

// Update message by id
func (d *MessageDao) Update(ctx context.Context, id int64, m *model.MessageOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "message dao update error")
	}

	return nil
}

// Delete delete message by id
func (d *MessageDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "message dao delete error")
	}

	return nil
}

func (d *MessageDao) DeleteByConversationID(ctx context.Context, conversationID int64) error {
	err := d.db(ctx).Model(d.m).Where("conversation_id = ?", conversationID).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "message dao delete by conversation id error")
	}

	return nil
}

func (d *MessageDao) GetByConversationID(ctx context.Context, id int64) ([]*model.Message, error) {
	var result []*model.Message

	err := d.db(ctx).Model(d.m).Where("conversation_id = ?", id).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "message dao get by conversation id error")
	}

	return result, nil
}

func (d *MessageDao) CreateBatch(ctx context.Context, ms []*model.Message) error {
	err := d.db(ctx).Model(d.m).CreateInBatches(ms, 100).Error
	if err != nil {
		return errors.Wrap(err, "message dao create batch error")
	}

	return nil
}

func (d *MessageDao) DeleteGtID(ctx context.Context, id int64, conversationID int64) error {
	err := d.db(ctx).Model(d.m).Where("id > ? AND conversation_id = ?", id, conversationID).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "message dao delete by created at error")
	}

	return nil
}

func (d *MessageDao) DeleteGteID(ctx context.Context, id int64, conversationID int64) error {
	err := d.db(ctx).Model(d.m).Where("id >= ? AND conversation_id = ?", id, conversationID).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "message dao delete by created at error")
	}

	return nil
}
