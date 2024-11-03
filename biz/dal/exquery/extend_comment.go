package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryVideoCommentExistById(id int64) (bool, error) {
	vc := dal.Executor.VideoComment
	count, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(id)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoCommentExistByIdAndUserId(id, userId int64) (bool, error) {
	vc := dal.Executor.VideoComment
	count, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(id), vc.UserID.Eq(userId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoCommentExistByIdAndVideoId(commentId, videoId int64) (bool, error) {
	vc := dal.Executor.VideoComment
	count, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(commentId), vc.VideoID.Eq(videoId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoCommentExistByIdAndRootId(id, rootId int64) (bool, error) {
	vc := dal.Executor.VideoComment
	count, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(id), vc.RootID.Eq(rootId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoCommentExistByIdParentIdAndRootId(id, parentId, rootId int64) (bool, error) {
	vc := dal.Executor.VideoComment
	count, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(id), vc.ParentID.Eq(parentId), vc.RootID.Eq(rootId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryVideoRootCommentByVideoIdPaged(videoId int64, pageNum, pageSize int) ([]*model.VideoComment, int64, error) {
	vc := dal.Executor.VideoComment
	comments, count, err := vc.WithContext(context.Background()).
		Where(vc.VideoID.Eq(videoId), vc.RootID.Eq(0), vc.ParentID.Eq(0)).
		Order(vc.CreatedAt.Desc()).
		FindByPage(pageNum*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}

func QueryVideoChildCommentByRootIdPaged(rootID int64, pageNum, pageSize int) ([]*model.VideoComment, int64, error) {
	vc := dal.Executor.VideoComment
	comments, count, err := vc.WithContext(context.Background()).
		Where(vc.RootID.Eq(rootID)).
		Order(vc.CreatedAt.Desc()).
		FindByPage(pageNum*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}

func QueryActivityCommentExistByIdAndUserId(id, userId int64) (bool, error) {
	ac := dal.Executor.ActivityComment
	count, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(id), ac.UserID.Eq(userId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityCommentExistByIdAndActivityId(commentId, activityId int64) (bool, error) {
	ac := dal.Executor.ActivityComment
	count, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(commentId), ac.ActivityID.Eq(activityId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityCommentExistById(id int64) (bool, error) {
	ac := dal.Executor.ActivityComment
	count, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(id)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityCommentExistByIdAndRootId(id, rootId int64) (bool, error) {
	ac := dal.Executor.ActivityComment
	count, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(id), ac.RootID.Eq(rootId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityCommentExistByIdParentIdAndRootId(id, parentId, rootId int64) (bool, error) {
	ac := dal.Executor.ActivityComment
	count, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(id), ac.ParentID.Eq(parentId), ac.RootID.Eq(rootId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityRootCommentByActivityIdPaged(activityId int64, pageNum, pageSize int) ([]*model.ActivityComment, int64, error) {
	ac := dal.Executor.ActivityComment
	comments, count, err := ac.WithContext(context.Background()).
		Where(ac.ActivityID.Eq(activityId), ac.RootID.Eq(0), ac.ParentID.Eq(0)).
		Order(ac.CreatedAt.Desc()).
		FindByPage(pageNum*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}

func QueryActivityChildCommentByRootIdPaged(rootID int64, pageNum, pageSize int) ([]*model.ActivityComment, int64, error) {
	ac := dal.Executor.ActivityComment
	comments, count, err := ac.WithContext(context.Background()).
		Where(ac.RootID.Eq(rootID)).
		Order(ac.CreatedAt.Desc()).
		FindByPage(pageNum*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}

func InsertVideoComment(items ...*model.VideoComment) error {
	vc := dal.Executor.VideoComment
	err := vc.WithContext(context.Background()).Create(items...)
	if err != nil {
		return err
	}
	return nil
}

func InsertActivityComment(items ...*model.ActivityComment) error {
	ac := dal.Executor.ActivityComment
	err := ac.WithContext(context.Background()).Create(items...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentById(id int64) error {
	vc := dal.Executor.VideoComment
	_, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentByVideoId(videoId int64) error {
	vc := dal.Executor.VideoComment
	_, err := vc.WithContext(context.Background()).Where(vc.VideoID.Eq(videoId)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentById(id int64) error {
	ac := dal.Executor.ActivityComment
	_, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentByActivityId(activityId int64) error {
	ac := dal.Executor.ActivityComment
	_, err := ac.WithContext(context.Background()).Where(ac.ActivityID.Eq(activityId)).Delete()
	if err != nil {
		return err
	}
	return nil
}
