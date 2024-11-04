package synchronizer

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/mw/redis"
)

func SynchronizeActivityLikeFromRedis2DB(vid string) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		if err := synchronizeNewInsertActivityLikeFromRedis2DB(vid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		if err := synchronizeNewDeleteActivityLikeFromRedis2DB(vid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return err
	default:
		return redis.DeleteActivityLikeListFromDynamicSpace(vid)
	}
}

func synchronizeNewInsertActivityLikeFromRedis2DB(vid string) error {
	activityId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return err
	}

	a := dal.Executor.Activity
	exist, err := a.WithContext(context.Background()).Where(a.ID.Eq(activityId)).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return nil
	}

	list, err := redis.GetNewUpdateActivityLikeList(vid)
	if err != nil {
		return err
	}
	alikes := []*model.ActivityLike{}
	uids := []int64{}
	for _, uid := range *list {
		userId, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			continue
		}
		alikes = append(alikes, &model.ActivityLike{ActivityID: activityId, UserID: userId})
		uids = append(uids, userId)
	}

	al := dal.Executor.ActivityLike
	if _, err = al.WithContext(context.Background()).Where(al.ActivityID.Eq(activityId), al.UserID.In(uids...)).Delete(); err != nil {
		return err
	}

	if err = al.WithContext(context.Background()).Create(alikes...); err != nil {
		return err
	}

	if err = redis.AppendActivityLikeListToStaticSpace(vid, *list); err != nil {
		return err
	}

	if err = redis.DeleteActivityLikeListFromDynamicSpace(vid); err != nil {
		return err
	}
	return nil
}

func synchronizeNewDeleteActivityLikeFromRedis2DB(vid string) error {
	activityId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return err
	}

	list, err := redis.GetNewDeleteActivityLikeList(vid)
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

	al := dal.Executor.ActivityLike
	if _, err = al.WithContext(context.Background()).Where(al.ActivityID.Eq(activityId), al.UserID.In(uids...)).Delete(); err != nil {
		return err
	}
	return nil
}

func SynchronizeActivityLikeFromDB2Redis() error {
	a := dal.Executor.Activity
	activities, err := a.WithContext(context.Background()).Select(a.ID).Find()
	if err != nil {
		return err
	}

	al := dal.Executor.ActivityLike
	for _, activity := range activities {
		l, err := al.WithContext(context.Background()).Where(al.ActivityID.Eq(activity.ID)).Find()
		if err != nil {
			return err
		}
		likes := []string{}
		for _, item := range l {
			likes = append(likes, fmt.Sprint(item.UserID))
		}
		if err = redis.PutActivityLikeInfo(fmt.Sprint(activity.ID), &likes); err != nil {
			return err
		}
	}
	return nil
}
