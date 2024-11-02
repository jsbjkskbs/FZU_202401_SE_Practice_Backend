package redis

import (
	"github.com/go-redis/redis"
)

func PutVideoCommentLikeInfo(cid string, uidList *[]string) error {
	pipe := commentInfoClient.TxPipeline()
	pipe.Del(`comment/video/like/` + cid)
	pipe.Del(`comment/video/changed_like/` + cid)
	for _, item := range *uidList {
		pipe.SAdd(`comment/video/like/`+cid, item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func AppendVideoCommentLikeInfo(cid, uid string) error {
	_, err := commentInfoClient.ZAdd(`comment/video/changed_like/`+cid, redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(`comment/video/like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendVideoCommentLikeInfoToStaticSpace(cid, uid string) error {
	if _, err := commentInfoClient.SAdd(`comment/video/like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendVideoCommentLikeListToStaticSpace(cid string, uidList []string) error {
	tx := commentInfoClient.TxPipeline()
	for _, uid := range uidList {
		tx.SAdd(`comment/video/like/`+cid, uid)
	}
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentLikeInfoFromDynamicSpace(cid, uid string) error {
	if _, err := commentInfoClient.ZRem(`comment/video/changed_like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentLikeListFromDynamicSpace(cid string) error {
	if _, err := commentInfoClient.Del(`comment/video/changed_like/` + cid).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveVideoCommentLikeInfo(cid, uid string) error {
	_, err := commentInfoClient.ZAdd(`comment/video/changed_like/`+cid, redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(`comment/video/like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func GetVideoCommentLikeList(cid string) (*[]string, error) {
	list, err := commentInfoClient.SMembers(`comment/video/like/` + cid).Result()
	if err != nil {
		return nil, err
	}
	nList, err := GetNewUpdateVideoLikeList(cid)
	if err != nil {
		return nil, err
	}
	list = append(list, *nList...)
	return &list, nil
}

func GetNewUpdateVideoCommentLikeList(cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(`comment/video/changed_like/`+cid, redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteVideoCommentLikeList(cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(`comment/video/changed_like/`+cid, redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetVideoCommentLikeCount(cid string) (int64, error) {
	var count int64
	var err error
	if count, err = commentInfoClient.SCard(`comment/video/like/` + cid).Result(); err != nil {
		return -1, err
	}
	if nCount, err := commentInfoClient.ZCount(`comment/video/changed_like/`+cid, `1`, `1`).Result(); err != nil {
		return -1, err
	} else {
		return count + nCount, nil
	}
}

func PutActivityCommentLikeInfo(cid string, uidList *[]string) error {
	pipe := commentInfoClient.TxPipeline()
	pipe.Del(`comment/activity/like/` + cid)
	pipe.Del(`comment/activity/changed_like/` + cid)
	for _, item := range *uidList {
		pipe.SAdd(`comment/activity/like/`+cid, item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func AppendActivityCommentLikeInfo(cid, uid string) error {
	_, err := commentInfoClient.ZAdd(`comment/activity/changed_like/`+cid, redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(`comment/activity/like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityCommentLikeInfoToStaticSpace(cid, uid string) error {
	if _, err := commentInfoClient.SAdd(`comment/activity/like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityCommentLikeListToStaticSpace(cid string, uidList []string) error {
	tx := commentInfoClient.TxPipeline()
	for _, uid := range uidList {
		tx.SAdd(`comment/activity/like/`+cid, uid)
	}
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentLikeInfoFromDynamicSpace(cid, uid string) error {
	if _, err := commentInfoClient.ZRem(`comment/activity/changed_like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentLikeListFromDynamicSpace(cid string) error {
	if _, err := commentInfoClient.Del(`comment/activity/changed_like/` + cid).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveActivityCommentLikeInfo(cid, uid string) error {
	_, err := commentInfoClient.ZAdd(`comment/activity/changed_like/`+cid, redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(`comment/activity/like/`+cid, uid).Result(); err != nil {
		return err
	}
	return nil
}

func GetActivityCommentLikeList(cid string) (*[]string, error) {
	list, err := commentInfoClient.SMembers(`comment/activity/like/` + cid).Result()
	if err != nil {
		return nil, err
	}
	nList, err := GetNewUpdateVideoLikeList(cid)
	if err != nil {
		return nil, err
	}
	list = append(list, *nList...)
	return &list, nil
}

func GetNewUpdateActivityCommentLikeList(cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(`comment/activity/changed_like/`+cid, redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteActivityCommentLikeList(cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(`comment/activity/changed_like/`+cid, redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetActivityCommentLikeCount(cid string) (int64, error) {
	var count int64
	var err error
	if count, err = commentInfoClient.SCard(`comment/activity/like/` + cid).Result(); err != nil {
		return -1, err
	}
	if nCount, err := commentInfoClient.ZCount(`comment/activity/changed_like/`+cid, `1`, `1`).Result(); err != nil {
		return -1, err
	} else {
		return count + nCount, nil
	}
}
