package db

import (
	"context"
	"github.com/aiagt/aiagt/apps/model/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"github.com/pkg/errors"
	"math"

	ktdb "github.com/aiagt/kitextool/option/server/db"
	"gorm.io/gorm"
)

type ApiKeyDao struct {
	m *model.ApiKey
}

// NewApiKeyDao make ApiKey dao
func NewApiKeyDao() *ApiKeyDao {
	return &ApiKeyDao{m: new(model.ApiKey)}
}

func (d *ApiKeyDao) db(ctx context.Context) *gorm.DB {
	return ktdb.DBCtx(ctx)
}

// GetByID get api_key by id
func (d *ApiKeyDao) GetByID(ctx context.Context, id int64) (*model.ApiKey, error) {
	var result model.ApiKey

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "api_key dao get by id error")
	}

	return &result, nil
}

// GetByIDs get api_key list by ids
func (d *ApiKeyDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.ApiKey, error) {
	var result []*model.ApiKey

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "api_key dao get by ids error")
	}

	return result, nil
}

// List get api_key list
func (d *ApiKeyDao) List(ctx context.Context, page *base.PaginationReq) ([]*model.ApiKey, *base.PaginationResp, error) {
	var (
		list   []*model.ApiKey
		total  int64
		offset = int((page.Page - 1) * page.PageSize)
		limit  = int(page.PageSize)
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "api_key dao get page error")
	}

	pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, pageResp, nil
}

// Create insert a api_key record
func (d *ApiKeyDao) Create(ctx context.Context, m *model.ApiKey) error {
	m.ID = snowflake.Generate().Int64()

	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "api_key dao create error")
	}

	return nil
}

// Update api_key by id
func (d *ApiKeyDao) Update(ctx context.Context, id int64, m *model.ApiKeyOptional) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "api_key dao update error")
	}

	return nil
}

// Delete delete api_key by id
func (d *ApiKeyDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "api_key dao delete error")
	}

	return nil
}

// GetBySourceOrDefault get api_key by source or default
func (d *ApiKeyDao) GetBySourceOrDefault(ctx context.Context, source string) (*model.ApiKey, error) {
	var result model.ApiKey

	err := d.db(ctx).Model(d.m).Where("source = ? OR source = ?", source, model.DefaultSource).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "api_key dao get by source error")
	}

	return &result, nil
}
