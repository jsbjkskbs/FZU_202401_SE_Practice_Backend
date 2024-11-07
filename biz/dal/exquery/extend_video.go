package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/service/common"

	"gorm.io/gen"
)

func QueryVideoIdAll() ([]*model.Video, error) {
	v := dal.Executor.Video
	videos, err := v.WithContext(context.Background()).Select(v.ID).Find()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func QueryVideoVisitCountAll() ([]*model.Video, error) {
	v := dal.Executor.Video
	videos, err := v.WithContext(context.Background()).Select(v.ID, v.VisitCount).Find()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

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

func QueryVideoFuzzyByKeywordPaged(keyword string, pageNum, pageSize int, fromDate, toDate *int64) ([]*model.Video, int64, error) {
	v := dal.Executor.Video
	vd := v.WithContext(context.Background())
	conditions := []gen.Condition{}
	if fromDate != nil {
		conditions = append(conditions, v.CreatedAt.Gte(*fromDate))
	}
	if toDate != nil {
		conditions = append(conditions, v.CreatedAt.Lte(*toDate))
	}
	conditions = append(conditions, v.Status.Eq(common.VideoStatusPassed))
	result, count, err := vd.Where(conditions...).
		Where(vd.Where(v.Title.Like("%"+keyword+"%")).Or(v.Description.Like("%"+keyword+"%"))).
		FindByPage(int(pageNum*pageSize), int(pageSize))
		/*
			SELECT *
				FROM video
				WHERE
					(video.title LIKE '%keyword%' OR video.description LIKE '%keyword%')
					AND
					video.status = 'passed'
					AND
					video.created_at >= fromDate
					AND
					video.created_at <= toDate
					LIMIT pageSize OFFSET pageNum
		*/
	if err != nil {
		return nil, 0, err
	}
	return result, count, nil
}

func QueryVideoByCategoryPaged(categoryId *int64, pageNum, pageSize int) ([]*model.Video, int64, error) {
	v := dal.Executor.Video
	conditions := []gen.Condition{v.Status.Eq(common.VideoStatusReview)}

	if categoryId != nil {
		conditions = append(conditions, v.CategoryID.Eq(*categoryId))
	}

	videos, count, err := v.WithContext(context.Background()).
		Where(conditions...).FindByPage((int(pageNum * pageSize)), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return videos, count, nil
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
