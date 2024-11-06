package synchronizer

import (
	"strconv"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/mw/redis"
)

func SynchronizeVideoVisitInfoRedis2DB(vid string) error {
	videoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return err
	}

	visitCount, err := redis.GetVideoVisitCount(vid)
	if err != nil {
		return err
	}

	exquery.UpdateVideoWithId(&model.Video{
		ID:         videoId,
		VisitCount: visitCount,
	})
	return nil
}

func SynchronizeVideoVisitInfoDB2Redis() error {
	videos, err := exquery.QueryVideoVisitCountAll()
	if err != nil {
		return err
	}
	for _, video := range videos {
		if err := redis.PutVideoVisitInfo(strconv.FormatInt(video.ID, 10), video.VisitCount); err != nil {
			return err
		}
	}

	return nil
}
