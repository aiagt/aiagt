package mapping

import (
	"github.com/aiagt/aiagt/app/user/model"
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
