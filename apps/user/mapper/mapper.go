package mapper

import (
	"github.com/aiagt/aiagt/apps/user/model"
	"github.com/aiagt/aiagt/common/baseutil"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
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

func NewGenSecret(secret *model.Secret) *usersvc.Secret {
	return &usersvc.Secret{
		Id:        secret.ID,
		UserId:    secret.UserID,
		PluginId:  secret.PluginID,
		Name:      secret.Name,
		Value:     secret.Value,
		CreatedAt: baseutil.NewBaseTime(secret.CreatedAt),
		UpdatedAt: baseutil.NewBaseTime(secret.UpdatedAt),
	}
}

func NewGenListSecret(secrets []*model.Secret) []*usersvc.Secret {
	result := make([]*usersvc.Secret, len(secrets))
	for i, secret := range secrets {
		result[i] = NewGenSecret(secret)
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

func NewModelCreateSecret(secret *usersvc.CreateSecretReq) *model.Secret {
	return &model.Secret{
		PluginID: secret.PluginId,
		Name:     secret.Name,
		Value:    secret.Value,
	}
}

func NewModelUpdateSecret(secret *usersvc.UpdateSecretReq) *model.SecretOptional {
	return &model.SecretOptional{
		PluginID: secret.PluginId,
		Name:     secret.Name,
		Value:    secret.Value,
	}
}
