package model_converter

import (
	"fmt"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/biz/mw/redis"
	"sfw/pkg/oss"
)

func ActivityListDal2Resp(list *[]*model.Activity, fromUser *string) (*[]*base.Activity, error) {
	resp := &[]*base.Activity{}
	for _, v := range *list {
		owner, err := exquery.QueryUserByID(v.UserID)
		if err != nil {
			return nil, err
		}

		user, err := UserDal2Resp(owner, fromUser)

		images, err := exquery.QueryImageUrlsByActivityId(v.ID)
		for i, image := range images {
			images[i] = oss.Key2Url(image)
		}

		refVideoID := ""
		refActivityID := ""
		if v.RefVideoID != 0 {
			refVideoID = fmt.Sprint(v.RefVideoID)
		}
		if v.RefActivityID != 0 {
			refActivityID = fmt.Sprint(v.RefActivityID)
		}

		likeCount, err := redis.GetActivityLikeCount(fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}

		commentCount, err := exquery.QueryActivityCommentCountById(v.ID)
		if err != nil {
			return nil, err
		}

		isLiked := false
		if fromUser != nil {
			isLiked, err = redis.IsActivityLikedByUser(fmt.Sprint(v.ID), *fromUser)
			if err != nil {
				return nil, err
			}
		}

		*resp = append(*resp, &base.Activity{
			ID:           fmt.Sprint(v.ID),
			User:         user,
			Content:      v.Content,
			Image:        images,
			RefVideo:     refVideoID,
			RefActivity:  refActivityID,
			LikeCount:    likeCount,
			CommentCount: commentCount,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			DeletedAt:    v.DeletedAt,
			IsLiked:      isLiked,
		})
	}
	return resp, nil
}
