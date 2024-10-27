package service_converter

import (
	"fmt"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
)

func UserListDal2Resp(list *[]*model.User) (resp *[]*base.User) {
	resp = &[]*base.User{}
	for _, v := range *list {
		*resp = append(*resp, &base.User{
			ID:        fmt.Sprint(v.ID),
			Username:  v.Username,
			AvatarURL: v.AvatarURL,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		})
	}
	return
}
