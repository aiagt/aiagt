package handler

import (
	"github.com/aiagt/aiagt/apps/model/dal/cache"
	"github.com/aiagt/aiagt/apps/model/dal/db"
)

// ModelServiceImpl implements the last service interface defined in the IDL.
type ModelServiceImpl struct {
	modelDao       *db.ModelDao
	apiKeyDao      *db.ApiKeyDao
	callTokenCache *cache.CallTokenCache
}

func NewModelService(modelDao *db.ModelDao, apiKyeDao *db.ApiKeyDao, callTokenCache *cache.CallTokenCache) *ModelServiceImpl {
	initServiceBusiness(5)

	return &ModelServiceImpl{modelDao: modelDao, apiKeyDao: apiKyeDao, callTokenCache: callTokenCache}
}
