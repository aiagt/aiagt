package mapper

import (
	"encoding/json"
	"github.com/aiagt/aiagt/pkg/safe"

	"github.com/aiagt/aiagt/app/app/model"
	"github.com/aiagt/aiagt/common/baseutil"
	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
)

func NewGenApp(app *model.App, author *usersvc.User, tools []*pluginsvc.PluginTool, labels []*appsvc.AppLabel) *appsvc.App {
	return &appsvc.App{
		Id:              app.ID,
		Name:            app.Name,
		Description:     app.Description,
		DescriptionMd:   app.DescriptionMd,
		ModelId:         app.ModelID,
		EnableImage:     app.EnableImage,
		EnableFile:      app.EnableFile,
		Version:         app.Version,
		IsPrivate:       app.IsPrivate,
		HomePage:        app.HomePage,
		PresetQuestions: app.PresetQuestions,
		ToolIds:         app.ToolIDs,
		Tools:           tools,
		Logo:            app.Logo,
		AuthorId:        app.AuthorID,
		Author:          author,
		LabelIds:        app.LabelIDs,
		Labels:          labels,
		ModelConfig:     NewGenModelConfig(app.ModelConfig),
		CreatedAt:       baseutil.NewBaseTime(app.CreatedAt),
		UpdatedAt:       baseutil.NewBaseTime(app.UpdatedAt),
		PublishedAt:     baseutil.NewBaseTimeP(app.PublishedAt),
	}
}

func NewGenModelConfig(modelConfig *model.ModelConfig) *appsvc.ModelConfig {
	if modelConfig == nil {
		return nil
	}

	result := &appsvc.ModelConfig{
		MaxTokens:        modelConfig.MaxTokens,
		Temperature:      modelConfig.Temperature,
		TopP:             modelConfig.TopP,
		N:                safe.Value(modelConfig.N),
		Stream:           safe.Value(modelConfig.Stream),
		Stop:             modelConfig.Stop,
		PresencePenalty:  modelConfig.PresencePenalty,
		Seed:             modelConfig.Seed,
		FrequencyPenalty: modelConfig.FrequencyPenalty,
		LogitBias:        modelConfig.LogitBias,
		Logprobs:         modelConfig.Logprobs,
		TopLogprobs:      modelConfig.TopLogprobs,
		User:             modelConfig.User,
		StreamOptions:    (*openai.StreamOptions)(modelConfig.StreamOptions),
	}

	if modelConfig.ResponseFormat != nil {
		var responseFormat openai.ChatCompletionResponseFormat

		err := json.Unmarshal([]byte(*modelConfig.ResponseFormat), &responseFormat)
		if err == nil {
			result.ResponseFormat = &responseFormat
		}
	}

	return result
}

func NewGenListApp(apps []*model.App) []*appsvc.App {
	result := make([]*appsvc.App, len(apps))
	for i, app := range apps {
		result[i] = NewGenApp(app, nil, nil, nil)
	}

	return result
}

func NewModelCreateApp(req *appsvc.CreateAppReq, userID int64) *model.App {
	return &model.App{
		Name:            req.Name,
		Description:     req.Description,
		DescriptionMd:   req.DescriptionMd,
		ModelID:         req.ModelId,
		EnableImage:     req.EnableImage,
		EnableFile:      req.EnableFile,
		Version:         req.Version,
		IsPrivate:       req.IsPrivate,
		HomePage:        req.HomePage,
		PresetQuestions: req.PresetQuestions,
		ToolIDs:         req.ToolIds,
		Logo:            req.Logo,
		AuthorID:        userID,
		LabelIDs:        req.LabelIds,
		ModelConfig:     NewModelModelConfig(req.ModelConfig),
	}
}

func NewModelModelConfig(modelConfig *appsvc.ModelConfig) *model.ModelConfig {
	if modelConfig == nil {
		return nil
	}

	result := &model.ModelConfig{
		MaxTokens:        modelConfig.MaxTokens,
		Temperature:      modelConfig.Temperature,
		TopP:             modelConfig.TopP,
		N:                safe.Pointer(modelConfig.N),
		Stream:           safe.Pointer(modelConfig.Stream),
		Stop:             modelConfig.Stop,
		PresencePenalty:  modelConfig.PresencePenalty,
		Seed:             modelConfig.Seed,
		FrequencyPenalty: modelConfig.FrequencyPenalty,
		LogitBias:        modelConfig.LogitBias,
		Logprobs:         modelConfig.Logprobs,
		TopLogprobs:      modelConfig.TopLogprobs,
		User:             modelConfig.User,
		StreamOptions:    (*model.StreamOptions)(modelConfig.StreamOptions),
	}

	if modelConfig.ResponseFormat != nil {
		responseFormat, err := json.Marshal(modelConfig.ResponseFormat)
		if err == nil {
			result.ResponseFormat = safe.Pointer(string(responseFormat))
		}
	}

	return result
}

func NewGenAppLabel(label *model.AppLabel) *appsvc.AppLabel {
	return &appsvc.AppLabel{
		Id:        label.ID,
		Text:      label.Text,
		CreatedAt: baseutil.NewBaseTime(label.CreatedAt),
	}
}

func NewGenListAppLabel(labels []*model.AppLabel) []*appsvc.AppLabel {
	result := make([]*appsvc.AppLabel, len(labels))
	for i, label := range labels {
		result[i] = NewGenAppLabel(label)
	}

	return result
}

func NewModelUpdateApp(req *appsvc.UpdateAppReq) *model.AppOptional {
	return &model.AppOptional{
		Name:            req.Name,
		Description:     req.Description,
		DescriptionMd:   req.DescriptionMd,
		ModelID:         req.ModelId,
		EnableImage:     req.EnableImage,
		EnableFile:      req.EnableFile,
		IsPrivate:       req.IsPrivate,
		HomePage:        req.HomePage,
		PresetQuestions: req.PresetQuestions,
		ToolIDs:         req.ToolIds,
		Logo:            req.Logo,
		LabelIDs:        req.LabelIds,
		ModelConfig:     NewModelModelConfig(req.ModelConfig),
	}
}
