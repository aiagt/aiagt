package db

import (
	"context"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	"math"

	"github.com/aiagt/aiagt/kitex_gen/usersvc"

	"github.com/aiagt/aiagt/apps/user/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type SecretDao struct {
	m *model.Secret
}

// NewSecretDao make Secret dao
func NewSecretDao() *SecretDao {
	return &SecretDao{m: new(model.Secret)}
}

func (d *SecretDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get secret by id
func (d *SecretDao) GetByID(ctx context.Context, id int64) (*model.Secret, error) {
	var result model.Secret

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "secret dao get by id error")
	}

	return &result, nil
}

// GetByIDs get secret list by ids
func (d *SecretDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.Secret, error) {
	var result []*model.Secret

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "secret dao get by ids error")
	}

	return result, nil
}

// List get secret list
func (d *SecretDao) List(ctx context.Context, req *usersvc.ListSecretReq, userID int64) ([]*model.Secret, *base.PaginationResp, error) {
	var (
		list   []*model.Secret
		total  int64
		page   = req.Pagination
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Scopes(func(db *gorm.DB) *gorm.DB {
		db = db.Where("user_id = ?", userID)

		if req.PluginId != nil {
			db = db.Where("plugin_id = ?", req.PluginId)
		}
		if req.Name != nil {
			db = db.Where("name = ?", req.Name)
		}

		return db
	}).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "secret dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a secret record
func (d *SecretDao) Create(ctx context.Context, m *model.Secret) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "secret dao create error")
	}

	return nil
}

// Update secret by id
func (d *SecretDao) Update(ctx context.Context, id int64, m *model.SecretOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "secret dao update error")
	}

	return nil
}

// Delete delete secret by id
func (d *SecretDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "secret dao delete error")
	}

	return nil
}
