package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryVideoCommentReportExistByIdAndStatus(id int64, status string) (bool, error) {
	vcr := dal.Executor.VideoCommentReport
	count, err := vcr.WithContext(context.Background()).Where(vcr.ID.Eq(id), vcr.Status.Eq(status)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoCommentReportCountByUserIdAndCommentId(userId, commentId int64) (int64, error) {
	vcr := dal.Executor.VideoCommentReport
	count, err := vcr.WithContext(context.Background()).Where(vcr.UserID.Eq(userId), vcr.CommentID.Eq(commentId)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func InsertVideoCommentReport(reports ...*model.VideoCommentReport) error {
	vcr := dal.Executor.VideoCommentReport
	err := vcr.WithContext(context.Background()).Create(reports...)
	if err != nil {
		return err
	}
	return nil
}

func UpdateVideoCommentReportById(report *model.VideoCommentReport) error {
	vcr := dal.Executor.VideoCommentReport
	_, err := vcr.WithContext(context.Background()).Where(vcr.ID.Eq(report.ID)).Updates(report)
	if err != nil {
		return err
	}
	return nil
}
