package handler

import (
	"github.com/aiagt/aiagt/app/model/conf"
	"github.com/sashabaranov/go-openai"
)

// ModelServiceImpl implements the last service interface defined in the IDL.
type ModelServiceImpl struct {
	openaiCli *openai.Client
}

func NewModelService() *ModelServiceImpl {
	initServiceBusiness(5)

	config := openai.DefaultConfig(conf.Conf().OpenAI.APIKey)
	config.BaseURL = conf.Conf().OpenAI.BaseURL

	openaiCli := openai.NewClientWithConfig(config)

	return &ModelServiceImpl{openaiCli: openaiCli}
}
