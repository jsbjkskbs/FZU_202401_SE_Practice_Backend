package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryVideoLikeVideoIds() ([]*model.VideoLike, error) {
	vl := dal.Executor.VideoLike
	likes, err := vl.WithContext(context.Background()).Select(vl.VideoID).Find()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func QueryVideoLikeUserIdsByVideoId(videoId int64) ([]*model.VideoLike, error) {
	vl := dal.Executor.VideoLike
	likes, err := vl.WithContext(context.Background()).Where(vl.VideoID.Eq(videoId)).Select(vl.UserID).Find()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func InsertVideoLikeByUserIds(videoId int64, userIds []int64) error {
	likes := []*model.VideoLike{}
	for _, userId := range userIds {
		likes = append(likes, &model.VideoLike{
			VideoID: videoId,
			UserID:  userId,
		})
	}
	vl := dal.Executor.VideoLike
	err := vl.WithContext(context.Background()).Create(likes...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoLikeByVideoIdAndUserIds(videoId int64, userIds []int64) error {
	vl := dal.Executor.VideoLike
	_, err := vl.WithContext(context.Background()).Where(vl.VideoID.Eq(videoId), vl.UserID.In(userIds...)).Delete()
	if err != nil {
		return err
	}
	return nil
}
