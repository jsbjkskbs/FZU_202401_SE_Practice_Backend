package exquery

import (
	"context"
	"fmt"
	"strconv"

	"sfw/biz/dal"
	"sfw/biz/dal/model"

	"gorm.io/gen"
)

func QueryVideoCommentAllIdAndVideoId() ([]*model.VideoComment, error) {
	vc := dal.Executor.VideoComment
	comments, err := vc.WithContext(context.Background()).Select(vc.ID, vc.VideoID).Find()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

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

func QueryVideoCommentReportByBasicInfoPaged(status, keyword, userId, label *string, pageNum, pageSize int64) ([]*model.VideoCommentReport, int64, error) {
	vcr := dal.Executor.VideoCommentReport
	conditions := []gen.Condition{}
	if status != nil {
		conditions = append(conditions, vcr.Status.Eq(*status))
	}
	if keyword != nil {
		conditions = append(conditions, vcr.Reason.Like(fmt.Sprint("%", *keyword, "%")))
	}
	if userId != nil {
		userId, err := strconv.ParseInt(*userId, 10, 64)
		if err != nil {
			return nil, 0, err
		}
		conditions = append(conditions, vcr.UserID.Eq(userId))
	}
	if label != nil {
		conditions = append(conditions, vcr.Label.Eq(*label))
	}
	// 此处代码不必提取至exquery
	items, count, err := vcr.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(pageNum * pageSize)), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func QueryVideoCommentIdAndVidByCommentId(commentId int64) ([]model.VideoComment, error) {
	vc := dal.Executor.VideoComment
	list := []model.VideoComment{}
	err := vc.WithContext(context.Background()).
		Where(vc.ID.Eq(commentId)).
		Or(vc.RootID.Eq(commentId)).
		Select(vc.ID, vc.VideoID).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func QueryActivityCommentAllIdAndActivityId() ([]*model.ActivityComment, error) {
	ac := dal.Executor.ActivityComment
	comments, err := ac.WithContext(context.Background()).Select(ac.ID, ac.ActivityID).Find()
	if err != nil {
		return nil, err
	}
	return comments, nil
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

func QueryActivityCommentReportByBasicInfoPaged(status, keyword, userId, label *string, pageNum, pageSize int64) ([]*model.ActivityCommentReport, int64, error) {
	ac := dal.Executor.ActivityCommentReport
	conditions := []gen.Condition{}
	if status != nil {
		conditions = append(conditions, ac.Status.Eq(*status))
	}
	if keyword != nil {
		conditions = append(conditions, ac.Reason.Like(fmt.Sprint("%", *keyword, "%")))
	}
	if userId != nil {
		userId, err := strconv.ParseInt(*userId, 10, 64)
		if err != nil {
			return nil, 0, err
		}
		conditions = append(conditions, ac.UserID.Eq(userId))
	}
	if label != nil {
		conditions = append(conditions, ac.Label.Eq(*label))
	}
	items, count, err := ac.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(pageNum * pageSize)), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func QueryActivityCommentIdAndVidByCommentId(commentId int64) ([]model.ActivityComment, error) {
	ac := dal.Executor.ActivityComment
	list := []model.ActivityComment{}
	err := ac.WithContext(context.Background()).
		Where(ac.ID.Eq(commentId)).
		Or(ac.RootID.Eq(commentId)).
		Select(ac.ID, ac.ActivityID).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
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

func DeleteVideoCommentCascadeById(id int64) error {
	vc := dal.Executor.VideoComment
	_, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(id)).Or(vc.RootID.Eq(id)).Delete()
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

func DeleteActivityCommentCascadeById(commentId int64) error {
	ac := dal.Executor.ActivityComment
	_, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(commentId)).Or(ac.RootID.Eq(commentId)).Delete()
	if err != nil {
		return err
	}
	return nil
}
