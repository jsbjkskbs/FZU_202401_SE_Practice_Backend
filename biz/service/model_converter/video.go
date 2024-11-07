package model_converter

import (
	"fmt"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/biz/mw/redis"
	"sfw/pkg/oss"
	"sfw/pkg/utils/checker"
)

func VideoDal2Resp(v *model.Video, fromUser *string) (*base.Video, error) {
	owner, err := exquery.QueryUserByID(v.UserID)
	if err != nil {
		return nil, err
	}

	user, err := UserDal2Resp(owner, fromUser)

	category := "null"
	for k, i := range checker.CategoryMap {
		if i == v.CategoryID {
			category = k
			break
		}
	}

	labels, err := exquery.QueryVideoLabels(v.ID)

	likeCount, err := redis.GetVideoLikeCount(fmt.Sprint(v.ID))
	if err != nil {
		return nil, err
	}

	commentCount, err := exquery.QueryVideoCommentCountById(v.ID)
	if err != nil {
		return nil, err
	}

	isLiked := false
	if fromUser != nil {
		isLiked, err = redis.IsVideoLikedByUser(fmt.Sprint(v.ID), *fromUser)
		if err != nil {
			return nil, err
		}
	}

	return &base.Video{
		ID:           fmt.Sprint(v.ID),
		User:         user,
		VideoURL:     oss.Key2Url(v.VideoURL),
		CoverURL:     oss.Key2Url(v.CoverURL),
		Title:        v.Title,
		Description:  v.Description,
		Category:     category,
		VisitCount:   v.VisitCount,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		Labels:       labels,
		Status:       v.Status,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		DeletedAt:    v.DeletedAt,
		IsLiked:      isLiked,
	}, nil
}

func VideoListDal2Resp(list *[]*model.Video, fromUser *string) ([]*base.Video, error) {
	resp := []*base.Video{}
	for _, v := range *list {
		owner, err := exquery.QueryUserByID(v.UserID)
		if err != nil {
			return nil, err
		}

		user, err := UserDal2Resp(owner, fromUser)
		if err != nil {
			return nil, err
		}

		category := "null"
		for k, i := range checker.CategoryMap {
			if i == v.CategoryID {
				category = k
				break
			}
		}

		labels, err := exquery.QueryVideoLabels(v.ID)

		likeCount, err := redis.GetVideoLikeCount(fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}

		commentCount, err := exquery.QueryVideoCommentCountById(v.ID)
		if err != nil {
			return nil, err
		}

		isLiked := false
		if fromUser != nil {
			isLiked, err = redis.IsVideoLikedByUser(fmt.Sprint(v.ID), *fromUser)
			if err != nil {
				return nil, err
			}
		}

		resp = append(resp, &base.Video{
			ID:           fmt.Sprint(v.ID),
			User:         user,
			VideoURL:     oss.Key2Url(v.VideoURL),
			CoverURL:     oss.Key2Url(v.CoverURL),
			Title:        v.Title,
			Description:  v.Description,
			Category:     category,
			VisitCount:   v.VisitCount,
			LikeCount:    likeCount,
			CommentCount: commentCount,
			Labels:       labels,
			Status:       v.Status,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			DeletedAt:    v.DeletedAt,
			IsLiked:      isLiked,
		})
	}
	return resp, nil
}
