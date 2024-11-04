package model_converter

import (
	"context"
	"fmt"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/biz/mw/redis"
)

func VideoCommentDal2Resp(list *[]*model.VideoComment, fromUser *string) (*[]*base.Comment, error) {
	vc := dal.Executor.VideoComment
	resp := &[]*base.Comment{}
	for _, v := range *list {
		likeCount, err := redis.GetVideoCommentLikeCount(fmt.Sprint(v.VideoID), fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}
		childCount, err := vc.WithContext(context.Background()).Where(vc.RootID.Eq(v.ID)).Count()
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
			UserID:     fmt.Sprint(v.UserID),
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
	ac := dal.Executor.ActivityComment
	resp := &[]*base.Comment{}
	for _, v := range *list {
		likeCount, err := redis.GetActivityCommentLikeCount(fmt.Sprint(v.ActivityID), fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}
		childCount, err := ac.WithContext(context.Background()).Where(ac.RootID.Eq(v.ID)).Count()
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
			UserID:     fmt.Sprint(v.UserID),
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
