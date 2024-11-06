package synchronizer

import (
	"fmt"
	"strconv"
	"sync"

	"sfw/biz/dal/exquery"
	"sfw/biz/mw/redis"
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

	exist, err := exquery.QueryVideoCommentExistById(commentId)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}

	list, err := redis.GetNewUpdateVideoCommentLikeList(vid, cid)
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

	if err = exquery.DeleteVideoCommentLikeByCommentIdAndUserIds(commentId, uids); err != nil {
		return err
	}
	if err = exquery.InsertVideoCommentLikeByUserIds(commentId, uids); err != nil {
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

	if err = exquery.DeleteVideoCommentLikeByCommentIdAndUserIds(commentId, uids); err != nil {
		return err
	}
	return nil
}

func SynchronizeVideoCommentLikeFromDB2Redis() error {
	vcomments, err := exquery.QueryVideoCommentAllIdAndVideoId()
	if err != nil {
		return err
	}

	for _, v := range vcomments {
		clikes, err := exquery.QueryVideoCommentLikeAllUserIdByCommentId(v.ID)
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

	exist, err := exquery.QueryActivityCommentExistById(commentId)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}

	list, err := redis.GetNewUpdateActivityCommentLikeList(aid, cid)
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

	if err = exquery.DeleteActivityCommentLikeByCommentIdAndUserIds(commentId, uids); err != nil {
		return err
	}
	if err = exquery.InsertActivityCommentLikeByUserIds(commentId, uids); err != nil {
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

	if err = exquery.DeleteActivityCommentLikeByCommentIdAndUserIds(commentId, uids); err != nil {
		return err
	}
	return nil
}

func SynchronizeActivityCommentLikeFromDB2Redis() error {
	acomments, err := exquery.QueryActivityCommentAllIdAndActivityId()
	if err != nil {
		return err
	}

	for _, a := range acomments {
		clikes, err := exquery.QueryActivityCommentLikeAllUserIdByCommentId(a.ID)
		if err != nil {
			return err
		}

		list := []string{}
		for _, acl := range clikes {
			list = append(list, fmt.Sprint(acl.UserID))
		}
		if err = redis.PutActivityCommentLikeInfo(fmt.Sprint(a.ActivityID), fmt.Sprint(a.ID), &list); err != nil {
			return err
		}
	}
	return nil
}
