package exquery

import (
	"context"
	"fmt"
	"strconv"

	"sfw/biz/dal"
	"sfw/biz/dal/model"

	"gorm.io/gen"
)

func QueryVideoReportExistByIdAndStatus(id int64, status string) (bool, error) {
	vr := dal.Executor.VideoReport
	count, err := vr.WithContext(context.Background()).Where(vr.ID.Eq(id), vr.Status.Eq(status)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoReportCountByUserIdAndVideoId(userId, videoId int64) (int64, error) {
	vr := dal.Executor.VideoReport
	count, err := vr.WithContext(context.Background()).Where(vr.UserID.Eq(userId), vr.VideoID.Eq(videoId)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryVideoReportByBasicInfoPaged(status, keyword, userId, label *string, pageNum, pageSize int64) ([]*model.VideoReport, int64, error) {
	vr := dal.Executor.VideoReport
	conditions := []gen.Condition{}
	if status != nil {
		conditions = append(conditions, vr.Status.Eq(*status))
	}
	if keyword != nil {
		conditions = append(conditions, vr.Reason.Like(fmt.Sprint("%", *keyword, "%")))
	}
	if userId != nil {
		userId, err := strconv.ParseInt(*userId, 10, 64)
		if err != nil {
			return nil, 0, err
		}
		conditions = append(conditions, vr.UserID.Eq(userId))
	}
	if label != nil {
		conditions = append(conditions, vr.Label.Eq(*label))
	}
	// 此处代码不必提取至exquery
	items, count, err := vr.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(pageNum * pageSize)), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func InsertVideoReport(reports ...*model.VideoReport) error {
	vr := dal.Executor.VideoReport
	err := vr.WithContext(context.Background()).Create(reports...)
	if err != nil {
		return err
	}
	return nil
}

func UpdateVideoReportById(report *model.VideoReport) error {
	vr := dal.Executor.VideoReport
	_, err := vr.WithContext(context.Background()).Where(vr.ID.Eq(report.ID)).Updates(report)
	if err != nil {
		return err
	}
	return nil
}
