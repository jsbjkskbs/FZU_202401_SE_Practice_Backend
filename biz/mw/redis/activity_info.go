package redis

import (
	"sync"

	"github.com/go-redis/redis"
)

func PutActivityLikeInfo(aid string, uidList *[]string) error {
	pipe := activityInfoClient.TxPipeline()
	pipe.Del(`activity/like/` + aid)
	pipe.Del(`activity/changed_like/` + aid)
	for _, item := range *uidList {
		pipe.SAdd(`activity/like/`+aid, item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func AppendActivityLikeInfo(aid, uid string) error {
	_, err := activityInfoClient.ZAdd(`activity/changed_like/`+aid, redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := activityInfoClient.SRem(`activity/like/`+aid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityLikeInfoToStaticSpace(aid, uid string) error {
	if _, err := activityInfoClient.SAdd(`activity/like/`+aid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityLikeListToStaticSpace(aid string, uidList []string) error {
	tx := activityInfoClient.TxPipeline()
	for _, uid := range uidList {
		tx.SAdd(`activity/like/`+aid, uid)
	}
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityLikeInfoFromDynamicSpace(aid, uid string) error {
	if _, err := activityInfoClient.ZRem(`activity/changed_like/`+aid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteActivityLikeListFromDynamicSpace(aid string) error {
	if _, err := activityInfoClient.Del(`activity/changed_like/` + aid).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveActivityLikeInfo(aid, uid string) error {
	_, err := activityInfoClient.ZAdd(`activity/changed_like/`+aid, redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := activityInfoClient.SRem(`activity/like/`+aid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func GetActivityLikeList(aid string) (*[]string, error) {
	list, err := activityInfoClient.SMembers(`activity/like/` + aid).Result()
	if err != nil {
		return nil, err
	}
	nList, err := GetNewUpdateVideoLikeList(aid)
	if err != nil {
		return nil, err
	}
	list = append(list, *nList...)
	return &list, nil
}

func GetNewUpdateActivityLikeList(aid string) (*[]string, error) {
	list, err := activityInfoClient.ZRangeByScore(`activity/changed_like/`+aid, redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteActivityLikeList(aid string) (*[]string, error) {
	list, err := activityInfoClient.ZRangeByScore(`activity/changed_like/`+aid, redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetActivityLikeCount(aid string) (int64, error) {
	var count int64
	var err error
	if count, err = activityInfoClient.SCard(`activity/like/` + aid).Result(); err != nil {
		return -1, err
	}
	if nCount, err := activityInfoClient.ZCount(`activity/changed_like/`+aid, `1`, `1`).Result(); err != nil {
		return -1, err
	} else {
		return count + nCount, nil
	}
}

func DeleteActivity(aid string) error {
	wg := sync.WaitGroup{}
	wg.Add(3)
	errs := make(chan error, 3)
	go func() {
		if err := DeleteActivityCommentInfo(aid); err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		_, err := activityInfoClient.Del(`activity/like/` + aid).Result()
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		err := activityInfoClient.Del(`activity/changed_like/` + aid).Err()
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return err
	default:
		return nil
	}
}
