package service_converter

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/biz/mw/redis"
	"sfw/biz/service/checker"
	"sfw/pkg/oss"
)

func VideoDal2Resp(v *model.Video) *base.Video {
	category := "null"
	for k, i := range checker.CategoryMap {
		if i == v.CategoryID {
			category = k
			break
		}
	}

	labelItems := []model.VideoLabel{}
	l := dal.Executor.VideoLabel
	err := l.WithContext(context.Background()).Where(l.VideoID.Eq(v.ID)).Scan(&labelItems)
	if err != nil {
		return nil
	}
	labels := []string{}
	for _, item := range labelItems {
		labels = append(labels, item.LabelName)
	}

	likeCount, err := redis.GetVideoLikeCount(fmt.Sprint(v.ID))
	if err != nil {
		return nil
	}

	c := dal.Executor.VideoComment
	commentCount, err := c.WithContext(context.Background()).Where(c.VideoID.Eq(v.ID)).Count()
	if err != nil {
		return nil
	}

	return &base.Video{
		ID:           fmt.Sprint(v.ID),
		UserID:       fmt.Sprint(v.UserID),
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
	}
}

func VideoListDal2Resp(list *[]*model.Video) ([]*base.Video, error) {
	resp := []*base.Video{}
	for _, v := range *list {
		category := "null"
		for k, i := range checker.CategoryMap {
			if i == v.CategoryID {
				category = k
				break
			}
		}

		labelItems := []model.VideoLabel{}
		l := dal.Executor.VideoLabel
		err := l.WithContext(context.Background()).Where(l.VideoID.Eq(v.ID)).Scan(&labelItems)
		if err != nil {
			return nil, err
		}
		labels := []string{}
		for _, item := range labelItems {
			labels = append(labels, item.LabelName)
		}

		likeCount, err := redis.GetVideoLikeCount(fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}

		c := dal.Executor.VideoComment
		commentCount, err := c.WithContext(context.Background()).Where(c.VideoID.Eq(v.ID)).Count()
		if err != nil {
			return nil, err
		}

		resp = append(resp, &base.Video{
			ID:           fmt.Sprint(v.ID),
			UserID:       fmt.Sprint(v.UserID),
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
		})
	}
	return resp, nil
}