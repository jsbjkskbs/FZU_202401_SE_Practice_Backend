package synchronizer

import (
	"fmt"
	"strconv"
	"sync"

	"sfw/biz/dal/exquery"
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

	exist, err := exquery.QueryActivityExistById(activityId)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}

	list, err := redis.GetNewUpdateActivityLikeList(vid)
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

	if err = exquery.DeleteActivityLikeByUserIds(activityId, uids); err != nil {
		return err
	}

	if err = exquery.InsertActivityLikeByUserIds(activityId, uids); err != nil {
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

	if err = exquery.DeleteActivityLikeByUserIds(activityId, uids); err != nil {
		return err
	}
	return nil
}

func SynchronizeActivityLikeFromDB2Redis() error {
	activities, err := exquery.QueryActivityLikeActivityIds()
	if err != nil {
		return err
	}

	for _, activity := range activities {
		l, err := exquery.QueryActivityLikeUserIdsByActivityId(activity.ActivityID)
		if err != nil {
			return err
		}
		likes := []string{}
		for _, item := range l {
			likes = append(likes, fmt.Sprint(item.UserID))
		}
		if err = redis.PutActivityLikeInfo(fmt.Sprint(activity.ActivityID), &likes); err != nil {
			return err
		}
	}
	return nil
}
