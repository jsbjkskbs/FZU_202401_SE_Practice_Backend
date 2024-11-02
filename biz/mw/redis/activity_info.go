package redis

import (
	"github.com/go-redis/redis"
)

func PutActivityLikeInfo(vid string, uidList *[]string) error {
	pipe := activityInfoClient.TxPipeline()
	pipe.Del(`activity/like/` + vid)
	pipe.Del(`activity/changed_like/` + vid)
	for _, item := range *uidList {
		pipe.SAdd(`activity/like/`+vid, item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func AppendActivityLikeInfo(vid, uid string) error {
	_, err := activityInfoClient.ZAdd(`activity/changed_like/`+vid, redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := activityInfoClient.SRem(`activity/like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityLikeInfoToStaticSpace(vid, uid string) error {
	if _, err := activityInfoClient.SAdd(`activity/like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityLikeListToStaticSpace(vid string, uidList []string) error {
	tx := activityInfoClient.TxPipeline()
	for _, uid := range uidList {
		tx.SAdd(`activity/like/`+vid, uid)
	}
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityLikeInfoFromDynamicSpace(vid, uid string) error {
	if _, err := activityInfoClient.ZRem(`activity/changed_like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteActivityLikeListFromDynamicSpace(vid string) error {
	if _, err := activityInfoClient.Del(`activity/changed_like/` + vid).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveActivityLikeInfo(vid, uid string) error {
	_, err := activityInfoClient.ZAdd(`activity/changed_like/`+vid, redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := activityInfoClient.SRem(`activity/like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func GetActivityLikeList(vid string) (*[]string, error) {
	list, err := activityInfoClient.SMembers(`activity/like/` + vid).Result()
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

func GetNewUpdateActivityLikeList(vid string) (*[]string, error) {
	list, err := activityInfoClient.ZRangeByScore(`activity/changed_like/`+vid, redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteActivityLikeList(vid string) (*[]string, error) {
	list, err := activityInfoClient.ZRangeByScore(`activity/changed_like/`+vid, redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetActivityLikeCount(vid string) (int64, error) {
	var count int64
	var err error
	if count, err = activityInfoClient.SCard(`activity/like/` + vid).Result(); err != nil {
		return -1, err
	}
	if nCount, err := activityInfoClient.ZCount(`activity/changed_like/`+vid, `1`, `1`).Result(); err != nil {
		return -1, err
	} else {
		return count + nCount, nil
	}
}
