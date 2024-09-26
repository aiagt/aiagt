package handler

import (
	"github.com/aiagt/aiagt/app/model/conf"
	"github.com/aiagt/aiagt/app/model/dal/cache"
	"github.com/aiagt/aiagt/app/model/dal/db"
	"github.com/sashabaranov/go-openai"
)

// ModelServiceImpl implements the last service interface defined in the IDL.
type ModelServiceImpl struct {
	modelDao       *db.ModelDao
	callTokenCache *cache.CallTokenCache

	openaiCli *openai.Client
}

func NewModelService(modelDao *db.ModelDao, callTokenCache *cache.CallTokenCache) *ModelServiceImpl {
	initServiceBusiness(5)

	config := openai.DefaultConfig(conf.Conf().OpenAI.APIKey)
	config.BaseURL = conf.Conf().OpenAI.BaseURL

	openaiCli := openai.NewClientWithConfig(config)

	return &ModelServiceImpl{modelDao: modelDao, callTokenCache: callTokenCache, openaiCli: openaiCli}
}
