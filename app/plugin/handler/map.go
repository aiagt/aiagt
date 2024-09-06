package handler

import (
	"github.com/aiagt/aiagt/app/plugin/model"
	"github.com/aiagt/aiagt/common/baseutil"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

func MapListPlugin(list []*model.Plugin) []*pluginsvc.Plugin {
	result := make([]*pluginsvc.Plugin, len(list))
	for i, plugin := range list {
		result[i] = &pluginsvc.Plugin{
			Id:            plugin.ID,
			Key:           plugin.Key,
			Name:          plugin.Name,
			Description:   plugin.Description,
			DescriptionMd: plugin.DescriptionMd,
			// TODO: Author:        &usersvc.User{Id: plugin.AuthorID},
			IsPrivate:    plugin.IsPrivate,
			HomePage:     plugin.HomePage,
			EnableSecret: plugin.EnableSecret,
			Secrets:      MapListPluginSecret(plugin.Secrets),
			Labels:       plugin.LabelIDs,
			// TODO: Tools:        plugin.ToolIDs,
			Logo:        plugin.Logo,
			CreatedAt:   baseutil.NewBaseTime(plugin.CreatedAt),
			UpdatedAt:   baseutil.NewBaseTime(plugin.UpdatedAt),
			PublishedAt: baseutil.NewBaseTime(plugin.PublishedAt),
		}
	}
	return result
}

func MapListPluginSecret(list []*model.PluginSecret) []*pluginsvc.PluginSecret {
	result := make([]*pluginsvc.PluginSecret, len(list))
	for i, secret := range list {
		result[i] = &pluginsvc.PluginSecret{
			Name:          secret.Name,
			Description:   secret.Description,
			AcquireMethod: secret.AcquireMethod,
			Link:          secret.Link,
		}
	}
	return result
}
