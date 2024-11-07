package model_converter

import (
	"fmt"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/biz/mw/redis"
)

func VideoCommentDal2Resp(list *[]*model.VideoComment, fromUser *string) (*[]*base.Comment, error) {
	resp := &[]*base.Comment{}
	for _, v := range *list {
		owner, err := exquery.QueryUserByID(v.UserID)
		if err != nil {
			return nil, err
		}

		user, err := UserDal2Resp(owner, fromUser)
		if err != nil {
			return nil, err
		}

		likeCount, err := redis.GetVideoCommentLikeCount(fmt.Sprint(v.VideoID), fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}
		childCount, err := exquery.QueryVideoCommentChildCommentCountById(v.ID)
		if err != nil {
			return nil, err
		}

		isLiked := false
		if fromUser != nil {
			isLiked, err = redis.IsVideoCommentLikedByUser(fmt.Sprint(v.VideoID), fmt.Sprint(v.ID), *fromUser)
			if err != nil {
				return nil, err
			}
		}

		*resp = append(*resp, &base.Comment{
			ID:         fmt.Sprint(v.ID),
			User:       user,
			Otype:      "video",
			Oid:        fmt.Sprint(v.VideoID),
			Content:    v.Content,
			RootID:     fmt.Sprint(v.RootID),
			ParentID:   fmt.Sprint(v.ParentID),
			LikeCount:  likeCount,
			ChildCount: childCount,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			DeletedAt:  v.DeletedAt,
			IsLiked:    isLiked,
		})
	}
	return resp, nil
}

func ActivityCommentDal2Resp(list *[]*model.ActivityComment, fromUser *string) (*[]*base.Comment, error) {
	resp := &[]*base.Comment{}
	for _, v := range *list {
		owner, err := exquery.QueryUserByID(v.UserID)
		if err != nil {
			return nil, err
		}

		user, err := UserDal2Resp(owner, fromUser)
		if err != nil {
			return nil, err
		}

		likeCount, err := redis.GetActivityCommentLikeCount(fmt.Sprint(v.ActivityID), fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}
		childCount, err := exquery.QueryActivityCommentChildCommentCountById(v.ID)
		if err != nil {
			return nil, err
		}

		isLiked := false
		if fromUser != nil {
			isLiked, err = redis.IsActivityCommentLikedByUser(fmt.Sprint(v.ActivityID), fmt.Sprint(v.ID), *fromUser)
			if err != nil {
				return nil, err
			}
		}

		*resp = append(*resp, &base.Comment{
			ID:         fmt.Sprint(v.ID),
			User:       user,
			Otype:      "activity",
			Oid:        fmt.Sprint(v.ActivityID),
			Content:    v.Content,
			RootID:     fmt.Sprint(v.RootID),
			ParentID:   fmt.Sprint(v.ParentID),
			LikeCount:  likeCount,
			ChildCount: childCount,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			DeletedAt:  v.DeletedAt,
			IsLiked:    isLiked,
		})
	}
	return resp, nil
}
