package model_converter

import (
	"fmt"
	"strconv"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/pkg/oss"
)

func UserDal2Resp(user *model.User, fromUser *string) (*base.User, error) {
	followed := false
	if fromUser != nil {
		userId, err := strconv.ParseInt(*fromUser, 10, 64)
		if err != nil {
			return nil, err
		}
		followed, err = exquery.QueryFollowExistByFollowerIDAndFollowedID(userId, user.ID)
		if err != nil {
			return nil, err
		}
	}

	resp := &base.User{
		ID:         fmt.Sprint(user.ID),
		Username:   user.Username,
		AvatarURL:  oss.Key2Url(user.AvatarURL),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt,
		IsFollowed: followed,
	}
	return resp, nil
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

func UserListDal2Resp(list *[]*model.User, fromUser *string) (*[]*base.User, error) {
	resp := &[]*base.User{}
	for _, v := range *list {
		followed := false
		if fromUser != nil {
			userID, err := strconv.ParseInt(*fromUser, 10, 64)
			if err != nil {
				return nil, err
			}
			followed, err = exquery.QueryFollowExistByFollowerIDAndFollowedID(userID, v.ID)
			if err != nil {
				return nil, err
			}
		}
		*resp = append(*resp, &base.User{
			ID:         fmt.Sprint(v.ID),
			Username:   v.Username,
			AvatarURL:  oss.Key2Url(v.AvatarURL),
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			DeletedAt:  v.DeletedAt,
			IsFollowed: followed,
		})
	}
	return resp, nil
}
