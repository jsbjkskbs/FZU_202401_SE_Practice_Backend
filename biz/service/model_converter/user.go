package model_converter

import (
	"fmt"

	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/pkg/oss"
)

func UserDal2Resp(user *model.User) (resp *base.User) {
	resp = &base.User{
		ID:        fmt.Sprint(user.ID),
		Username:  user.Username,
		AvatarURL: oss.Key2Url(user.AvatarURL),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return
}

func UserWithTokenDal2Resp(user *model.User) (resp *base.UserWithToken) {
	resp = &base.UserWithToken{
		ID:        fmt.Sprint(user.ID),
		Username:  user.Username,
		AvatarURL: oss.Key2Url(user.AvatarURL),
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return
}

func UserListDal2Resp(list *[]*model.User) (resp *[]*base.User) {
	resp = &[]*base.User{}
	for _, v := range *list {
		*resp = append(*resp, &base.User{
			ID:        fmt.Sprint(v.ID),
			Username:  v.Username,
			AvatarURL: oss.Key2Url(v.AvatarURL),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		})
	}
	return
}
