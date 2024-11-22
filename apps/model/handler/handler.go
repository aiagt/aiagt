package handler

import (
	"github.com/aiagt/aiagt/apps/model/conf"
	"github.com/aiagt/aiagt/apps/model/dal/cache"
	"github.com/aiagt/aiagt/apps/model/dal/db"
	"github.com/sashabaranov/go-openai"
)

// ModelServiceImpl implements the last service interface defined in the IDL.
type ModelServiceImpl struct {
	modelDao       *db.ModelDao
	callTokenCache *cache.CallTokenCache
}

func NewModelService(modelDao *db.ModelDao, callTokenCache *cache.CallTokenCache) *ModelServiceImpl {
	initServiceBusiness(5)

	return &ModelServiceImpl{modelDao: modelDao, callTokenCache: callTokenCache}
}

func (*ModelServiceImpl) openaiCli() *openai.Client {
	config := openai.DefaultConfig(conf.Conf().OpenAI.APIKey)
	config.BaseURL = conf.Conf().OpenAI.BaseURL

	return openai.NewClientWithConfig(config)
}
