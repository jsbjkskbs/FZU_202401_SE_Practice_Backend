package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryActivityReportExistByIdAndStatus(id int64, status string) (bool, error) {
	ar := dal.Executor.ActivityReport
	count, err := ar.WithContext(context.Background()).Where(ar.ID.Eq(id), ar.Status.Eq(status)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityReportCountByUserIdAndActivityId(userId, activityId int64) (int64, error) {
	ar := dal.Executor.ActivityReport
	count, err := ar.WithContext(context.Background()).Where(ar.UserID.Eq(userId), ar.ActivityID.Eq(activityId)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func InsertActivityReport(reports ...*model.ActivityReport) error {
	ar := dal.Executor.ActivityReport
	err := ar.WithContext(context.Background()).Create(reports...)
	if err != nil {
		return err
	}
	return nil
}

func UpdateActivityReportById(report *model.ActivityReport) error {
	ar := dal.Executor.ActivityReport
	_, err := ar.WithContext(context.Background()).Where(ar.ID.Eq(report.ID)).Updates(report)
	if err != nil {
		return err
	}
	return nil
}
