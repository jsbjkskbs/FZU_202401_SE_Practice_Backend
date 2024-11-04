package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/service/common"
)

func QueryVideoExistById(id int64) (bool, error) {
	v := dal.Executor.Video
	count, err := v.WithContext(context.Background()).Where(v.ID.Eq(id)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoExistByIdAndStatus(id int64, status string) (bool, error) {
	v := dal.Executor.Video
	count, err := v.WithContext(context.Background()).Where(v.ID.Eq(id), v.Status.Eq(status)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoExistByIdAndUserId(id, userId int64) (bool, error) {
	v := dal.Executor.Video
	count, err := v.WithContext(context.Background()).Where(v.ID.Eq(id), v.UserID.Eq(userId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoById(id int64) (*model.Video, error) {
	v := dal.Executor.Video
	video, _, err := v.WithContext(context.Background()).Where(v.ID.Eq(id)).FindByPage(0, 1)
	if err != nil {
		return nil, err
	}
	if len(video) == 0 {
		return nil, nil
	}
	return video[0], nil
}

func QueryVideoByUserIdPaged(userId int64, pageNum, pageSize int) ([]*model.Video, int64, error) {
	v := dal.Executor.Video
	result, count, err := v.WithContext(context.Background()).Where(v.UserID.Eq(userId), v.Status.Neq(common.VideoStatusSubmit)).FindByPage(pageNum*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return result, count, nil
}

func QueryVideoByUserIdAndStatusPaged(userId int64, pageNum, pageSize int, status string) ([]*model.Video, int64, error) {
	v := dal.Executor.Video
	result, count, err := v.WithContext(context.Background()).
		Where(v.UserID.Eq(userId), v.Status.Eq(status)).
		FindByPage(pageNum*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return result, count, nil
}

func InsertVideo(videos ...*model.Video) error {
	v := dal.Executor.Video
	err := v.WithContext(context.Background()).Create(videos...)
	if err != nil {
		return err
	}
	return nil
}

func UpdateVideoWithId(video *model.Video) error {
	v := dal.Executor.Video
	_, err := v.WithContext(context.Background()).Where(v.ID.Eq(video.ID)).Updates(video)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoById(id int64) error {
	v := dal.Executor.Video
	_, err := v.WithContext(context.Background()).Where(v.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}
