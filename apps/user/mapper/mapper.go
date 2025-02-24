package mapper

import (
	"github.com/aiagt/aiagt/apps/user/model"
	"github.com/aiagt/aiagt/common/baseutil"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
)

func NewGenUser(user *model.User) *usersvc.User {
	return &usersvc.User{
		Id:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		Signature:     user.Signature,
		Homepage:      user.Homepage,
		DescriptionMd: user.DescriptionMd,
		Github:        user.Github,
		Avatar:        user.Avatar,
		CreatedAt:     baseutil.NewBaseTime(user.CreatedAt),
		UpdatedAt:     baseutil.NewBaseTime(user.UpdatedAt),
	}
}

func NewGenListUser(users []*model.User) []*usersvc.User {
	result := make([]*usersvc.User, len(users))
	for i, user := range users {
		result[i] = NewGenUser(user)
	}

	return result
}

func NewGenSecret(secret *model.Secret, plugin *pluginsvc.Plugin) *usersvc.Secret {
	result := &usersvc.Secret{
		Id:        secret.ID,
		UserId:    secret.UserID,
		PluginId:  secret.PluginID,
		Name:      secret.Name,
		Value:     secret.Value,
		CreatedAt: baseutil.NewBaseTime(secret.CreatedAt),
		UpdatedAt: baseutil.NewBaseTime(secret.UpdatedAt),
	}

	if plugin != nil {
		result.PluginName = &plugin.Name
		result.PluginLogo = &plugin.Logo
	}

	return result
}

func NewGenListSecret(secrets []*model.Secret, pluginMap hmap.Map[int64, *pluginsvc.Plugin]) []*usersvc.Secret {
	result := make([]*usersvc.Secret, len(secrets))
	for i, secret := range secrets {
		plugin := pluginMap[secret.PluginID]

		result[i] = NewGenSecret(secret, plugin)
	}

	return result
}

func NewModelUpdateUser(user *usersvc.UpdateUserReq) *model.UserOptional {
	return &model.UserOptional{
		Username:      user.Username,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		Signature:     user.Signature,
		Homepage:      user.Homepage,
		DescriptionMd: user.DescriptionMd,
		Github:        user.Github,
		Avatar:        user.Avatar,
	}
}
