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

func SynchronizeVideoLikeFromRedis2DB(vid string) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		if err := synchronizeNewInsertVideoLikeFromRedis2DB(vid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		if err := synchronizeNewDeleteVideoLikeFromRedis2DB(vid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return err
	default:
		return redis.DeleteVideoLikeListFromDynamicSpace(vid)
	}
}

// synchronizeNewInsertVideoLikeFromRedis2DB is a function to synchronize new insert video like from redis to database
func synchronizeNewInsertVideoLikeFromRedis2DB(vid string) error {
	videoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return err
	}

	v := dal.Executor.Video
	exist, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return nil
	}

	list, err := redis.GetNewUpdateVideoLikeList(vid)
	if err != nil {
		return err
	}
	uids := []int64{}
	vlikes := []*model.VideoLike{}
	for _, uid := range *list {
		userId, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			continue
		}
		uids = append(uids, userId)
		vlikes = append(vlikes, &model.VideoLike{VideoID: videoId, UserID: userId})
	}

	vl := dal.Executor.VideoLike
	if _, err = vl.WithContext(context.Background()).Where(vl.VideoID.Eq(videoId), vl.UserID.In(uids...)).Delete(); err != nil {
		return err
	}

	if err = vl.WithContext(context.Background()).Create(vlikes...); err != nil {
		return err
	}

	if err = redis.AppendVideoLikeListToStaticSpace(vid, *list); err != nil {
		return err
	}
	return nil
}

func synchronizeNewDeleteVideoLikeFromRedis2DB(vid string) error {
	videoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return err
	}

	list, err := redis.GetNewDeleteVideoLikeList(vid)
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

	vl := dal.Executor.VideoLike
	if _, err = vl.WithContext(context.Background()).Where(vl.VideoID.Eq(videoId), vl.UserID.In(uids...)).Delete(); err != nil {
		return err
	}
	return nil
}

func SynchronizeVideoLikeFromDB2Redis() error {
	v := dal.Executor.Video
	videos, err := v.WithContext(context.Background()).Select(v.ID).Find()
	if err != nil {
		return err
	}

	for _, video := range videos {
		vl := dal.Executor.VideoLike
		vlikes, err := vl.WithContext(context.Background()).Where(vl.VideoID.Eq(video.ID)).Find()
		if err != nil {
			return err
		}

		list := []string{}
		for _, vlike := range vlikes {
			list = append(list, fmt.Sprint(vlike.UserID))
		}

		if err = redis.PutVideoLikeInfo(fmt.Sprint(video.ID), &list); err != nil {
			return err
		}
	}

	return nil
}
