package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryActivityCommentReportExistByIdAndStatus(id int64, status string) (bool, error) {
	ac := dal.Executor.ActivityCommentReport
	count, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(id), ac.Status.Eq(status)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityCommentReportCountByUserIdAndCommentId(userId, commentId int64) (int64, error) {
	ac := dal.Executor.ActivityCommentReport
	count, err := ac.WithContext(context.Background()).Where(ac.UserID.Eq(userId), ac.CommentID.Eq(commentId)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func InsertActivityCommentReport(reports ...*model.ActivityCommentReport) error {
	ac := dal.Executor.ActivityCommentReport
	err := ac.WithContext(context.Background()).Create(reports...)
	if err != nil {
		return err
	}
	return nil
}

func UpdateActivityCommentReportById(report *model.ActivityCommentReport) error {
	ac := dal.Executor.ActivityCommentReport
	_, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(report.ID)).Updates(report)
	if err != nil {
		return err
	}
	return nil
}
