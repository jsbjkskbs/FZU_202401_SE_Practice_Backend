package synchronizer

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/mw/redis"
	"strconv"
	"sync"
)

func SynchronizeVideoCommentLikeFromRedis2DB(vid, cid string) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		if err := synchronizeNewInsertVideoCommentLikeFromDB2Redis(vid, cid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		if err := synchronizeNewDeleteVideoCommentLikeFromDB2Redis(vid, cid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return err
	default:
		return redis.DeleteVideoCommentLikeListFromDynamicSpace(vid, cid)
	}
}

func synchronizeNewInsertVideoCommentLikeFromDB2Redis(vid, cid string) error {
	commentId, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	vc := dal.Executor.VideoComment
	exist, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(commentId)).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return nil
	}

	list, err := redis.GetNewUpdateVideoCommentLikeList(vid, cid)
	if err != nil {
		return err
	}
	clikes := []*model.VideoCommentLike{}
	uids := []int64{}
	for _, uid := range *list {
		userId, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			continue
		}
		uids = append(uids, userId)
		clikes = append(clikes, &model.VideoCommentLike{CommentID: commentId, UserID: userId})
	}

	cl := dal.Executor.VideoCommentLike
	if _, err = cl.WithContext(context.Background()).Where(cl.CommentID.Eq(commentId), cl.UserID.In(uids...)).Delete(); err != nil {
		return err
	}
	if err = cl.WithContext(context.Background()).Create(clikes...); err != nil {
		return err
	}

	if err = redis.AppendVideoCommentLikeListToStaticSpace(vid, cid, *list); err != nil {
		return err
	}

	return nil
}

func synchronizeNewDeleteVideoCommentLikeFromDB2Redis(vid, cid string) error {
	commentId, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	list, err := redis.GetNewDeleteVideoCommentLikeList(vid, cid)
	if err != nil {
		return err
	}
	uids := []int64{}
	for _, uid := range *list {
		userId, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			continue
		}
		uids = append(uids, userId)
	}

	cl := dal.Executor.VideoCommentLike
	if _, err = cl.WithContext(context.Background()).Where(cl.CommentID.Eq(commentId), cl.UserID.In(uids...)).Delete(); err != nil {
		return err
	}
	return nil
}

func SynchronizeVideoCommentLikeFromDB2Redis() error {
	vc := dal.Executor.VideoComment
	vcomments, err := vc.WithContext(context.Background()).Select(vc.ID, vc.VideoID).Find()
	if err != nil {
		return err
	}

	for _, v := range vcomments {
		cl := dal.Executor.VideoCommentLike
		clikes, err := cl.WithContext(context.Background()).Find()
		if err != nil {
			return err
		}

		list := []string{}
		for _, vcl := range clikes {
			list = append(list, fmt.Sprint(vcl.UserID))
		}
		if err = redis.PutVideoCommentLikeInfo(fmt.Sprint(v.VideoID), fmt.Sprint(v.ID), &list); err != nil {
			return err
		}
	}
	return nil
}

func SynchronizeActivityCommentLikeFromRedis2DB(aid, cid string) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		if err := synchronizeNewInsertActivityCommentLikeFromDB2Redis(aid, cid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		if err := synchronizeNewDeleteActivityCommentLikeFromDB2Redis(aid, cid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return err
	default:
		return redis.DeleteActivityCommentLikeListFromDynamicSpace(aid, cid)
	}
}

func synchronizeNewInsertActivityCommentLikeFromDB2Redis(aid, cid string) error {
	commentId, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	ac := dal.Executor.ActivityComment
	exist, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(commentId)).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return nil
	}

	list, err := redis.GetNewUpdateActivityCommentLikeList(aid, cid)
	if err != nil {
		return err
	}
	clikes := []*model.ActivityCommentLike{}
	uids := []int64{}
	for _, uid := range *list {
		userId, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			continue
		}
		uids = append(uids, userId)
		clikes = append(clikes, &model.ActivityCommentLike{CommentID: commentId, UserID: userId})
	}

	cl := dal.Executor.ActivityCommentLike
	if _, err = cl.WithContext(context.Background()).Where(cl.CommentID.Eq(commentId), cl.UserID.In(uids...)).Delete(); err != nil {
		return err
	}
	if err = cl.WithContext(context.Background()).Create(clikes...); err != nil {
		return err
	}

	if err = redis.AppendActivityCommentLikeListToStaticSpace(aid, cid, *list); err != nil {
		return err
	}

	return nil
}

func synchronizeNewDeleteActivityCommentLikeFromDB2Redis(aid, cid string) error {
	commentId, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	list, err := redis.GetNewDeleteActivityCommentLikeList(aid, cid)
	if err != nil {
		return err
	}
	uids := []int64{}
	for _, uid := range *list {
		userId, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			continue
		}
		uids = append(uids, userId)
	}

	cl := dal.Executor.ActivityCommentLike
	if _, err = cl.WithContext(context.Background()).Where(cl.CommentID.Eq(commentId), cl.UserID.In(uids...)).Delete(); err != nil {
		return err
	}
	return nil
}

func SynchronizeActivityCommentLikeFromDB2Redis() error {
	ac := dal.Executor.ActivityComment
	acomments, err := ac.WithContext(context.Background()).Select(ac.ID, ac.ActivityID).Find()
	if err != nil {
		return err
	}

	for _, a := range acomments {
		cl := dal.Executor.ActivityCommentLike
		clikes, err := cl.WithContext(context.Background()).Find()
		if err != nil {
			return err
		}

		list := []string{}
		for _, acl := range clikes {
			list = append(list, fmt.Sprint(acl.UserID))
		}
		if err = redis.PutActivityCommentLikeInfo(fmt.Sprint(ac.ActivityID), fmt.Sprint(a.ID), &list); err != nil {
			return err
		}
	}
	return nil
}
