package db

import (
	"context"
	"math"

	"github.com/aiagt/aiagt/app/chat/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/pkg/errors"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type ChatDao struct {
	m *model.Chat
}

// NewChatDao make Chat dao
func NewChatDao() *ChatDao {
	return &ChatDao{m: new(model.Chat)}
}

func (d *ChatDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get chat by id
func (d *ChatDao) GetByID(ctx context.Context, id int64) (*model.Chat, error) {
	var result model.Chat

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "chat dao get by id error")
	}

	return &result, nil
}

// GetByIDs get chat list by ids
func (d *ChatDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Chat, error) {
	var result []*model.Chat

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "chat dao get by ids error")
	}

	return result, nil
}

// List get chat list
func (d *ChatDao) List(ctx context.Context, page *base.PaginationReq) ([]*model.Chat, *base.PaginationResp, error) {
	var (
		list   []*model.Chat
		total  int64
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "chat dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a chat record
func (d *ChatDao) Create(ctx context.Context, m *model.Chat) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "chat dao create error")
	}

	return nil
}

// Update chat by id
func (d *ChatDao) Update(ctx context.Context, id int64, m *model.ChatOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "chat dao update error")
	}

	return nil
}

// Delete delete chat by id
func (d *ChatDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "chat dao delete error")
	}

	return nil
}
