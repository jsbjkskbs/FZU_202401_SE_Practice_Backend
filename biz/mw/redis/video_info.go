package redis

import (
	"sfw/pkg/errno"
	"sync"

	"github.com/go-redis/redis"
)

func PutIPVisitInfo(vid, ip string) error {
	_, err := videoInfoClient.SAdd(`video/visit/`+vid, ip).Result()
	if err != nil {
		return err
	}
	return nil
}

func IsIPVisited(vid, ip string) (bool, error) {
	exist, err := videoInfoClient.Exists(`video/visit/` + vid).Result()
	if err != nil {
		return false, err
	}
	if exist == 0 {
		return false, nil
	}

	visited, err := videoInfoClient.SIsMember(`video/visit/`+vid, ip).Result()
	if err != nil {
		return false, err
	}
	return visited, nil
}

func IncrVideoVisitCount(vid string) error {
	_, err := videoInfoClient.ZIncrBy(`visit`, 1, vid).Result()
	if err != nil {
		return err
	}
	return nil
}

func DelIPVisitInfo(vid, ip string) error {
	_, err := videoInfoClient.SRem(`video/visit/`+vid, ip).Result()
	if err != nil {
		return err
	}
	return nil
}

func PutVideoLikeInfo(vid string, uidList *[]string) error {
	pipe := videoInfoClient.TxPipeline()
	pipe.Del(`video/like/` + vid)
	pipe.Del(`video/changed_like/` + vid)
	for _, item := range *uidList {
		pipe.SAdd(`video/like/`+vid, item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func PutVideoVisitInfo(vid string, visitCount int64) error {
	txpipe := videoInfoClient.TxPipeline()
	txpipe.ZAdd(`visit`, redis.Z{Score: float64(visitCount), Member: vid})
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
	_, err := videoInfoClient.ZAdd(`video/changed_like/`+vid, redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := videoInfoClient.SRem(`video/like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendVideoLikeInfoToStaticSpace(vid, uid string) error {
	if _, err := videoInfoClient.SAdd(`video/like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendVideoLikeListToStaticSpace(vid string, uidList []string) error {
	tx := videoInfoClient.TxPipeline()
	for _, uid := range uidList {
		tx.SAdd(`video/like/`+vid, uid)
	}
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoLikeInfoFromDynamicSpace(vid, uid string) error {
	if _, err := videoInfoClient.ZRem(`video/changed_like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteVideoLikeListFromDynamicSpace(vid string) error {
	if _, err := videoInfoClient.Del(`video/changed_like/` + vid).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveVideoLikeInfo(vid, uid string) error {
	if !IsVideoExist(vid) {
		return errno.ResourceNotFound
	}
	_, err := videoInfoClient.ZAdd(`video/changed_like/`+vid, redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := videoInfoClient.SRem(`video/like/`+vid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func GetVideoLikeList(vid string) (*[]string, error) {
	list, err := videoInfoClient.SMembers(`video/like/` + vid).Result()
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
	list, err := videoInfoClient.ZRangeByScore(`video/changed_like/`+vid, redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteVideoLikeList(vid string) (*[]string, error) {
	list, err := videoInfoClient.ZRangeByScore(`video/changed_like/`+vid, redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetVideoLikeCount(vid string) (int64, error) {
	var count int64
	var err error
	if count, err = videoInfoClient.SCard(`video/like/` + vid).Result(); err != nil {
		return -1, err
	}
	if nCount, err := videoInfoClient.ZCount(`video/changed_like/`+vid, `1`, `1`).Result(); err != nil {
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

func isVideoLikedByUserInStaticSpace(vid, uid string) (bool, error) {
	exist, err := videoInfoClient.Exists(`video/like/` + vid).Result()
	if err != nil {
		return false, err
	}
	if exist == 0 {
		return false, nil
	}

	isLiked, err := videoInfoClient.SIsMember(`video/like/`+vid, uid).Result()
	if err != nil {
		return false, err
	}

	return isLiked, nil
}

func isVideoLikedByUserInDynamicSpace(vid, uid string) (bool, error) {
	exist, err := videoInfoClient.Exists(`video/changed_like/` + vid).Result()
	if err != nil {
		return false, err
	}
	if exist == 0 {
		return false, nil
	}

	score, err := videoInfoClient.ZScore(`video/changed_like/`+vid, uid).Result()
	if err != nil {
		return false, err
	}

	return score == 1, nil
}

func IsVideoLikedByUser(vid, uid string) (bool, error) {
	inStaticSpace, err := isVideoLikedByUserInStaticSpace(vid, uid)
	if err != nil {
		return false, err
	}
	if inStaticSpace {
		return true, nil
	}

	inDynamicSpace, err := isVideoLikedByUserInDynamicSpace(vid, uid)
	if err != nil {
		return false, err
	}

	return inDynamicSpace, nil
}

func IsVideoExist(vid string) bool {
	_, err := videoInfoClient.ZScore(`visit`, vid).Result()
	return err == nil
}

func DeleteVideo(vid string) error {
	wg := sync.WaitGroup{}
	wg.Add(4)
	errs := make(chan error, 4)
	go func() {
		err := DeleteVideoCommentInfo(vid)
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		_, err := videoInfoClient.Del(`video/like/` + vid).Result()
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		_, err := videoInfoClient.Del(`video/changed_like/` + vid).Result()
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		_, err := videoInfoClient.ZRem(`visit`, vid).Result()
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
