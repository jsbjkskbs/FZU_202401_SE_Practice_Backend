package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryVideoCommentLikeAllUserIdByCommentId(commentId int64) ([]*model.VideoCommentLike, error) {
	vcl := dal.Executor.VideoCommentLike
	likes, err := vcl.WithContext(context.Background()).Where(vcl.CommentID.Eq(commentId)).Select(vcl.UserID).Find()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func InsertVideoCommentLikeByUserIds(commentId int64, userIds []int64) error {
	vcls := []*model.VideoCommentLike{}
	for _, userId := range userIds {
		vcls = append(vcls, &model.VideoCommentLike{
			CommentID: commentId,
			UserID:    userId,
		})
	}
	vcl := dal.Executor.VideoCommentLike
	err := vcl.WithContext(context.Background()).Create(vcls...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentLikeByCommentIdAndUserIds(commentId int64, userIds []int64) error {
	vcl := dal.Executor.VideoCommentLike
	_, err := vcl.WithContext(context.Background()).Where(vcl.CommentID.Eq(commentId), vcl.UserID.In(userIds...)).Delete()
	if err != nil {
		return err
	}
	return nil
}
