package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryActivityLikeActivityIds() ([]*model.ActivityLike, error) {
	al := dal.Executor.ActivityLike
	likes, err := al.WithContext(context.Background()).Select(al.ActivityID).Find()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func QueryActivityLikeUserIdsByActivityId(actvityId int64) ([]*model.ActivityLike, error) {
	al := dal.Executor.ActivityLike
	likes, err := al.WithContext(context.Background()).Where(al.ActivityID.Eq(actvityId)).Select(al.UserID).Find()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func InsertActivityLikeByUserIds(activityId int64, userIds []int64) error {
	likes := []*model.ActivityLike{}
	for _, userId := range userIds {
		likes = append(likes, &model.ActivityLike{
			ActivityID: activityId,
			UserID:     userId,
		})
	}
	al := dal.Executor.ActivityLike
	err := al.WithContext(context.Background()).Create(likes...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityLikeByUserIds(activityId int64, userIds []int64) error {
	al := dal.Executor.ActivityLike
	_, err := al.WithContext(context.Background()).Where(al.ActivityID.Eq(activityId), al.UserID.In(userIds...)).Delete()
	if err != nil {
		return err
	}
	return nil
}
