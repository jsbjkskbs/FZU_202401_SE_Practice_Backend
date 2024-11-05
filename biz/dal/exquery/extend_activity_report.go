package exquery

import (
	"context"
	"fmt"
	"strconv"

	"sfw/biz/dal"
	"sfw/biz/dal/model"

	"gorm.io/gen"
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

func QueryActivityReportByBasicInfoPaged(status, keyword, userId, label *string, pageNum, pageSize int64) ([]*model.ActivityReport, int64, error) {
	ar := dal.Executor.ActivityReport
	conditions := []gen.Condition{}
	if status != nil {
		conditions = append(conditions, ar.Status.Eq(*status))
	}
	if keyword != nil {
		conditions = append(conditions, ar.Reason.Like(fmt.Sprint("%", *keyword, "%")))
	}
	if userId != nil {
		userId, err := strconv.ParseInt(*userId, 10, 64)
		if err != nil {
			return nil, 0, err
		}
		conditions = append(conditions, ar.UserID.Eq(userId))
	}
	if label != nil {
		conditions = append(conditions, ar.Label.Eq(*label))
	}
	// 此处代码不必提取至exquery
	items, count, err := ar.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(pageNum * pageSize)), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
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
