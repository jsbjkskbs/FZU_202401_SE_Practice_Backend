package redis

import (
	"sfw/pkg/errno"
	"strconv"

	"github.com/go-redis/redis"
)

func PutVideoLikeInfo(vid string, uidList *[]string) error {
	pipe := videoInfoClient.TxPipeline()
	pipe.Del(`l:` + vid)
	pipe.Del(`nl:` + vid)
	for _, item := range *uidList {
		pipe.SAdd(`l:`+vid, item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func PutVideoVisitInfo(vid, visitCount, category string) error {
	score, _ := strconv.ParseFloat(visitCount, 64)
	txpipe := videoInfoClient.TxPipeline()
	txpipe.ZAdd(`visit`, redis.Z{Score: score, Member: vid})
	txpipe.ZAdd(`visit_`+category, redis.Z{Score: score, Member: vid})
	_, err := txpipe.Exec()
	if err != nil {
		return err
	}
	return nil
}

func AppendVideoLikeInfo(vid, uid string) error {
	if !IsVideoExist(vid) {
		return errno.ResourceNotFound
	}
	_, err := videoInfoClient.ZAdd(`nl:`+vid, redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := videoInfoClient.SRem(`l:`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendVideoLikeInfoToStaticSpace(vid, uid string) error {
	if _, err := videoInfoClient.SAdd(`l:`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteVideoLikeInfoFromDynamicSpace(vid, uid string) error {
	if _, err := videoInfoClient.ZRem(`nl:`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveVideoLikeInfo(vid, uid string) error {
	if !IsVideoExist(vid) {
		return errno.ResourceNotFound
	}
	_, err := videoInfoClient.ZAdd(`nl:`+vid, redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := videoInfoClient.SRem(`l:`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func IncrVideoVisitInfo(vid string, category string) error {
	txpipe := videoInfoClient.TxPipeline()
	txpipe.ZIncrBy(`visit`, 1, vid)
	txpipe.ZIncrBy(`visit_`+category, 1, vid)
	_, err := txpipe.Exec()
	if err != nil {
		return err
	}
	return nil
}

func IsVideoLikedByUser(vid, uid string) (bool, error) {
	exist, err := videoInfoClient.SIsMember(`l:`+vid, uid).Result()
	if err != nil {
		return false, err
	}
	if !exist {
		_, err := videoInfoClient.ZRank(`nl:`+vid, uid).Result()
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return true, nil
	}
}

func GetVideoLikeList(vid string) (*[]string, error) {
	list, err := videoInfoClient.SMembers(`l:` + vid).Result()
	if err != nil {
		return nil, err
	}
	nList, err := GetNewUpdateVideoLikeList(vid)
	if err != nil {
		return nil, err
	}
	list = append(list, *nList...)
	return &list, nil
}

func GetNewUpdateVideoLikeList(vid string) (*[]string, error) {
	list, err := videoInfoClient.ZRangeByScore(`nl:`+vid, redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteVideoLikeList(vid string) (*[]string, error) {
	list, err := videoInfoClient.ZRangeByScore(`nl:`+vid, redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetVideoLikeCount(vid string) (int64, error) {
	var count int64
	var err error
	if count, err = videoInfoClient.SCard(`l:` + vid).Result(); err != nil {
		return -1, err
	}
	if nCount, err := videoInfoClient.ZCount(`nl:`+vid, `1`, `1`).Result(); err != nil {
		return -1, err
	} else {
		return count + nCount, nil
	}
}

func GetVideoVisitCount(vid string) (int64, error) {
	_, err := videoInfoClient.ZRank(`visit`, vid).Result()
	if err != nil {
		return -1, err
	}
	s, err := videoInfoClient.ZScore(`visit`, vid).Result()
	if err != nil {
		return -1, err
	}
	return int64(s), nil
}

func GetVideoPopularList(pageNum, pageSize int64) (*[]string, error) {
	list, err := videoInfoClient.ZRevRange(`visit`, (pageNum-1)*pageSize, pageNum*pageSize-1).Result()
	if err != nil {
		return nil, err
	}
	return &list, err
}

func IsVideoExist(vid string) bool {
	_, err := videoInfoClient.ZScore(`visit`, vid).Result()
	return err == nil
}
