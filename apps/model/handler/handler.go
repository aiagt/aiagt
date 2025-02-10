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

func (*ModelServiceImpl) openaiCli(source string) *openai.Client {
	return newOpenaiCli(conf.Conf().APIKeys.GetOrDefault(source))
}

func newOpenaiCli(api *conf.APIKey) *openai.Client {
	config := openai.DefaultConfig(api.APIKey)
	config.BaseURL = api.BaseURL

	return openai.NewClientWithConfig(config)
}
