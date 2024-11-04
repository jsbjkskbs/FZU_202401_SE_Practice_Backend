package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
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
