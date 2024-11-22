package mapper

import (
	"encoding/json"

	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/lists"
	"github.com/aiagt/aiagt/pkg/utils"

	"github.com/aiagt/aiagt/apps/app/model"
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
		N:                utils.Value(modelConfig.N),
		Stream:           utils.Value(modelConfig.Stream),
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

func NewGenListApp(apps []*model.App, labels hmap.Map[int64, *appsvc.AppLabel]) []*appsvc.App {
	result := make([]*appsvc.App, len(apps))
	for i, app := range apps {
		appLabels := lists.Map(app.LabelIDs, func(t int64) *appsvc.AppLabel { return labels[t] })
		appLabels = lists.Filter(appLabels, func(t *appsvc.AppLabel) bool { return t != nil })

		result[i] = NewGenApp(app, nil, nil, appLabels)
	}

	return result
}

func NewModelCreateApp(req *appsvc.CreateAppReq, userID int64, labelIDs []int64) *model.App {
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
		LabelIDs:        labelIDs,
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
		N:                utils.Pointer(modelConfig.N),
		Stream:           utils.Pointer(modelConfig.Stream),
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
			result.ResponseFormat = utils.Pointer(string(responseFormat))
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

func NewModelUpdateApp(req *appsvc.UpdateAppReq, labelIDs []int64) *model.AppOptional {
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
		LabelIDs:        labelIDs,
		ModelConfig:     NewModelModelConfig(req.ModelConfig),
	}
}
