package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryActivityCommentLikeAllUserIdByCommentId(commentId int64) ([]*model.ActivityCommentLike, error) {
	acl := dal.Executor.ActivityCommentLike
	likes, err := acl.WithContext(context.Background()).Where(acl.CommentID.Eq(commentId)).Select(acl.UserID).Find()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func InsertActivityCommentLikeByUserIds(commentId int64, userIds []int64) error {
	acls := []*model.ActivityCommentLike{}
	for _, userId := range userIds {
		acls = append(acls, &model.ActivityCommentLike{
			CommentID: commentId,
			UserID:    userId,
		})
	}
	acl := dal.Executor.ActivityCommentLike
	err := acl.WithContext(context.Background()).Create(acls...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentLikeByCommentIdAndUserIds(commentId int64, userIds []int64) error {
	acl := dal.Executor.ActivityCommentLike
	_, err := acl.WithContext(context.Background()).Where(acl.CommentID.Eq(commentId), acl.UserID.In(userIds...)).Delete()
	if err != nil {
		return err
	}
	return nil
}
