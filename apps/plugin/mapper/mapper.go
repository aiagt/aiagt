package mapper

import (
	"encoding/json"
	"github.com/aiagt/aiagt/pkg/caller"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/snowflake"
	"github.com/aiagt/aiagt/tools/utils/lists"

	"github.com/aiagt/aiagt/apps/plugin/model"
	"github.com/aiagt/aiagt/common/baseutil"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/tools/utils/logger"
)

func NewGenPlugin(plugin *model.Plugin, author *usersvc.User, labels []*pluginsvc.PluginLabel, tools []*pluginsvc.PluginTool) *pluginsvc.Plugin {
	result := &pluginsvc.Plugin{
		Id:            plugin.ID,
		Key:           plugin.Key,
		Name:          plugin.Name,
		Description:   plugin.Description,
		DescriptionMd: plugin.DescriptionMd,
		AuthorId:      plugin.AuthorID,
		IsPrivate:     plugin.IsPrivate,
		HomePage:      plugin.HomePage,
		EnableSecret:  plugin.EnableSecret,
		Secrets:       NewGenListPluginSecret(plugin.Secrets),
		LabelIds:      plugin.LabelIDs,
		ToolIds:       plugin.ToolIDs,
		Logo:          plugin.Logo,
		CreatedAt:     baseutil.NewBaseTime(plugin.CreatedAt),
		UpdatedAt:     baseutil.NewBaseTime(plugin.UpdatedAt),
		PublishedAt:   baseutil.NewBaseTimeP(plugin.PublishedAt),
	}
	if author != nil {
		result.Author = author
	}

	if labels != nil {
		result.Labels = labels
	}

	if tools != nil {
		result.Tools = tools
	}

	return result
}

func NewGenListPlugin(list []*model.Plugin, labels hmap.Map[int64, *pluginsvc.PluginLabel]) []*pluginsvc.Plugin {
	result := make([]*pluginsvc.Plugin, len(list))
	for i, plugin := range list {
		pluginLabels := lists.Map(plugin.LabelIDs, func(t int64) *pluginsvc.PluginLabel { return labels[t] })
		pluginLabels = lists.Filter(pluginLabels, func(t *pluginsvc.PluginLabel) bool { return t != nil })

		result[i] = NewGenPlugin(plugin, nil, pluginLabels, nil)
	}

	return result
}

func NewGenPluginSecret(secret *model.PluginSecret) *pluginsvc.PluginSecret {
	return &pluginsvc.PluginSecret{
		Name:          secret.Name,
		Description:   secret.Description,
		AcquireMethod: secret.AcquireMethod,
		Link:          secret.Link,
	}
}

func NewGenListPluginSecret(list []*model.PluginSecret) []*pluginsvc.PluginSecret {
	result := make([]*pluginsvc.PluginSecret, len(list))
	for i, secret := range list {
		result[i] = NewGenPluginSecret(secret)
	}

	return result
}

func NewGenPluginTool(tool *model.PluginTool) *pluginsvc.PluginTool {
	requestType, err := json.Marshal(tool.RequestType)
	if err != nil {
		logger.Warnf("plugin tool request type marshaling error: %s", err.Error())
	}

	responseType, err := json.Marshal(tool.ResponseType)
	if err != nil {
		logger.Warnf("plugin tool response type marshaling error: %s", err.Error())
	}

	return &pluginsvc.PluginTool{
		Id:            tool.ID,
		Name:          tool.Name,
		Description:   tool.Description,
		PluginId:      tool.PluginID,
		RequestType:   requestType,
		ResponseType:  responseType,
		ApiUrl:        tool.ApiURL,
		ImportModelId: tool.ImportModelID,
		CreatedAt:     baseutil.NewBaseTime(tool.CreatedAt),
		UpdatedAt:     baseutil.NewBaseTime(tool.UpdatedAt),
		TestedAt:      baseutil.NewBaseTimeP(tool.TestedAt),
	}
}

func NewGenListPluginTool(list []*model.PluginTool) []*pluginsvc.PluginTool {
	result := make([]*pluginsvc.PluginTool, len(list))
	for i, tool := range list {
		result[i] = NewGenPluginTool(tool)
	}

	return result
}

func NewGenListPluginToolWithPlugin(list []*model.PluginTool, pluginMap map[int64]*model.Plugin) []*pluginsvc.PluginTool {
	result := make([]*pluginsvc.PluginTool, len(list))
	for i, tool := range list {
		result[i] = NewGenPluginTool(tool)
		result[i].Plugin = NewGenPlugin(pluginMap[tool.PluginID], nil, nil, nil)
	}

	return result
}

func NewGenPluginLabel(label *model.PluginLabel) *pluginsvc.PluginLabel {
	return &pluginsvc.PluginLabel{
		Id:        label.ID,
		Text:      label.Text,
		CreatedAt: baseutil.NewBaseTime(label.CreatedAt),
	}
}

func NewGenListPluginLabel(list []*model.PluginLabel) []*pluginsvc.PluginLabel {
	result := make([]*pluginsvc.PluginLabel, len(list))
	for i, label := range list {
		result[i] = NewGenPluginLabel(label)
	}

	return result
}

func NewModelPluginSecret(secret *pluginsvc.PluginSecret) *model.PluginSecret {
	return &model.PluginSecret{
		Name:          secret.Name,
		Description:   secret.Description,
		AcquireMethod: secret.AcquireMethod,
		Link:          secret.Link,
	}
}

func NewModelListPluginSecret(list []*pluginsvc.PluginSecret) []*model.PluginSecret {
	result := make([]*model.PluginSecret, len(list))
	for i, secret := range list {
		result[i] = NewModelPluginSecret(secret)
	}

	return result
}

func NewModelCreatePlugin(plugin *pluginsvc.CreatePluginReq, userID int64, labelIDs []int64) *model.Plugin {
	var key = plugin.Key
	if key == 0 {
		key = snowflake.Generate().Int64()
	}

	return &model.Plugin{
		Key:           key,
		Name:          plugin.Name,
		Description:   plugin.Description,
		DescriptionMd: plugin.DescriptionMd,
		AuthorID:      userID,
		IsPrivate:     plugin.IsPrivate,
		HomePage:      plugin.HomePage,
		EnableSecret:  plugin.EnableSecret,
		Secrets:       NewModelListPluginSecret(plugin.Secrets),
		LabelIDs:      labelIDs,
		ToolIDs:       plugin.ToolIds,
		Logo:          plugin.Logo,
	}
}

func NewModelUpdatePlugin(plugin *pluginsvc.UpdatePluginReq, labelIDs []int64) *model.PluginOptional {
	return &model.PluginOptional{
		Name:          plugin.Name,
		Description:   plugin.Description,
		DescriptionMd: plugin.DescriptionMd,
		IsPrivate:     plugin.IsPrivate,
		HomePage:      plugin.HomePage,
		EnableSecret:  plugin.EnableSecret,
		Secrets:       NewModelListPluginSecret(plugin.Secrets),
		LabelIDs:      labelIDs,
		ToolIDs:       plugin.ToolIds,
		Logo:          plugin.Logo,
	}
}

func NewModelCreatePluginTool(tool *pluginsvc.CreatePluginToolReq) *model.PluginTool {
	var (
		requestType  caller.RequestType
		responseType caller.ResponseType
	)

	err := json.Unmarshal(tool.RequestType, &requestType)
	if err != nil {
		logger.Warnf("plugin tool request type unmarshaling error: %s", err.Error())
	}

	err = json.Unmarshal(tool.ResponseType, &responseType)
	if err != nil {
		logger.Warnf("plugin tool response type unmarshaling error: %s", err.Error())
	}

	return &model.PluginTool{
		Name:          tool.Name,
		Description:   tool.Description,
		PluginID:      tool.PluginId,
		RequestType:   &requestType,
		ResponseType:  &responseType,
		ApiURL:        tool.ApiUrl,
		ImportModelID: tool.ImportModelId,
	}
}

func NewModelUpdatePluginTool(tool *pluginsvc.UpdatePluginToolReq) *model.PluginToolOptional {
	var (
		requestType  *caller.RequestType
		responseType *caller.ResponseType
	)

	if tool.RequestType != nil {
		requestType = new(caller.RequestType)

		err := json.Unmarshal(tool.RequestType, requestType)
		if err != nil {
			logger.Warnf("plugin tool request type unmarshaling error: %s", err.Error())
		}
	}

	if tool.ResponseType != nil {
		responseType = new(caller.ResponseType)

		err := json.Unmarshal(tool.ResponseType, &responseType)
		if err != nil {
			logger.Warnf("plugin tool response type unmarshaling error: %s", err.Error())
		}
	}

	return &model.PluginToolOptional{
		Name:          tool.Name,
		Description:   tool.Description,
		PluginID:      tool.PluginId,
		RequestType:   requestType,
		ResponseType:  responseType,
		ApiURL:        tool.ApiUrl,
		ImportModelID: tool.ImportModelId,
		TestedAt:      nil,
	}
}
