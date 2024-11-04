package redis

import (
	"strings"

	"github.com/go-redis/redis"
)

func PutVideoCommentLikeInfo(vid, cid string, uidList *[]string) error {
	pipe := commentInfoClient.TxPipeline()
	pipe.Del(strings.Join([]string{`comment/video/like`, vid, cid}, `/`))
	pipe.Del(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`))
	for _, item := range *uidList {
		pipe.SAdd(strings.Join([]string{`comment/video/like`, vid, cid}, `/`), item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func AppendVideoCommentLikeInfo(vid, cid, uid string) error {
	_, err := commentInfoClient.ZAdd(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`), redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(strings.Join([]string{`comment/video/like`, vid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendVideoCommentLikeInfoToStaticSpace(vid, cid, uid string) error {
	if _, err := commentInfoClient.SAdd(strings.Join([]string{`comment/video/like`, vid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendVideoCommentLikeListToStaticSpace(vid, cid string, uidList []string) error {
	tx := commentInfoClient.TxPipeline()
	for _, uid := range uidList {
		tx.SAdd(strings.Join([]string{`comment/video/like`, vid, cid}, `/`), uid)
	}
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentLikeInfoFromDynamicSpace(vid, cid, uid string) error {
	if _, err := commentInfoClient.ZRem(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentLikeListFromDynamicSpace(vid, cid string) error {
	if _, err := commentInfoClient.Del(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`)).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveVideoCommentLikeInfo(vid, cid, uid string) error {
	_, err := commentInfoClient.ZAdd(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`), redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(strings.Join([]string{`comment/video/like`, vid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func GetVideoCommentLikeList(vid, cid string) (*[]string, error) {
	list, err := commentInfoClient.SMembers(strings.Join([]string{`comment/video/like`, vid, cid}, `/`)).Result()
	if err != nil {
		return nil, err
	}
	nList, err := GetNewUpdateVideoCommentLikeList(vid, cid)
	if err != nil {
		return nil, err
	}
	list = append(list, *nList...)
	return &list, nil
}

func GetNewUpdateVideoCommentLikeList(vid, cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`), redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteVideoCommentLikeList(vid, cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`), redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetVideoCommentLikeCount(vid, cid string) (int64, error) {
	var count int64
	var err error
	if count, err = commentInfoClient.SCard(strings.Join([]string{`comment/video/like`, vid, cid}, `/`)).Result(); err != nil {
		return -1, err
	}
	if nCount, err := commentInfoClient.ZCount(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`), `1`, `1`).Result(); err != nil {
		return -1, err
	} else {
		return count + nCount, nil
	}
}

func DeleteVideoComment(vid, cid string) error {
	tx := commentInfoClient.TxPipeline()
	tx.Del(strings.Join([]string{`comment/video/like`, vid, cid}, `/`))
	tx.Del(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`))
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoCommentInfo(vid string) error {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = commentInfoClient.Scan(cursor, strings.Join([]string{`comment/video/like`, vid, `*`}, `/`), 1000).Result()
		if err != nil {
			return err
		}
		if len(keys) > 0 {
			if _, err := commentInfoClient.Del(keys...).Result(); err != nil {
				return err
			}
		}
		if cursor == 0 {
			break
		}
	}
	return nil
}

func PutActivityCommentLikeInfo(aid, cid string, uidList *[]string) error {
	pipe := commentInfoClient.TxPipeline()
	pipe.Del(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`))
	pipe.Del(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`))
	for _, item := range *uidList {
		pipe.SAdd(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`), item)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func AppendActivityCommentLikeInfo(aid, cid, uid string) error {
	_, err := commentInfoClient.ZAdd(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`), redis.Z{Score: 1, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityCommentLikeInfoToStaticSpace(aid, cid, uid string) error {
	if _, err := commentInfoClient.SAdd(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func AppendActivityCommentLikeListToStaticSpace(aid, cid string, uidList []string) error {
	tx := commentInfoClient.TxPipeline()
	for _, uid := range uidList {
		tx.SAdd(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`), uid)
	}
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentLikeInfoFromDynamicSpace(aid, cid, uid string) error {
	if _, err := commentInfoClient.ZRem(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentLikeListFromDynamicSpace(aid, cid string) error {
	if _, err := commentInfoClient.Del(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`)).Result(); err != nil {
		return err
	}
	return nil
}

func RemoveActivityCommentLikeInfo(aid, cid, uid string) error {
	_, err := commentInfoClient.ZAdd(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`), redis.Z{Score: 2, Member: uid}).Result()
	if err != nil {
		return err
	}
	if _, err := commentInfoClient.SRem(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`), uid).Result(); err != nil {
		return err
	}
	return nil
}

func GetActivityCommentLikeList(aid, cid string) (*[]string, error) {
	list, err := commentInfoClient.SMembers(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`)).Result()
	if err != nil {
		return nil, err
	}
	nList, err := GetNewUpdateActivityCommentLikeList(aid, cid)
	if err != nil {
		return nil, err
	}
	list = append(list, *nList...)
	return &list, nil
}

func GetNewUpdateActivityCommentLikeList(aid, cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`), redis.ZRangeBy{Min: `1`, Max: `1`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetNewDeleteActivityCommentLikeList(aid, cid string) (*[]string, error) {
	list, err := commentInfoClient.ZRangeByScore(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`), redis.ZRangeBy{Min: `2`, Max: `2`}).Result()
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func GetActivityCommentLikeCount(aid, cid string) (int64, error) {
	var count int64
	var err error
	if count, err = commentInfoClient.SCard(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`)).Result(); err != nil {
		return -1, err
	}
	if nCount, err := commentInfoClient.ZCount(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`), `1`, `1`).Result(); err != nil {
		return -1, err
	} else {
		return count + nCount, nil
	}
}

func isVideoCommentLikedByUserInStaticSpace(vid, cid, uid string) (bool, error) {
	exist, err := commentInfoClient.Exists(strings.Join([]string{`comment/video/like`, vid, cid}, `/`)).Result()
	if err != nil {
		return false, err
	}

	if exist == 0 {
		return false, nil
	}

	isLiked, err := commentInfoClient.SIsMember(strings.Join([]string{`comment/video/like`, vid, cid}, `/`), uid).Result()
	if err != nil {
		return false, err
	}

	return isLiked, nil
}

func isVideoCommentLikedByUserInDynamicSpace(vid, cid, uid string) (bool, error) {
	exist, err := commentInfoClient.Exists(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`)).Result()
	if err != nil {
		return false, err
	}

	if exist == 0 {
		return false, nil
	}

	score, err := commentInfoClient.ZScore(strings.Join([]string{`comment/video/changed_like`, vid, cid}, `/`), uid).Result()
	if err != nil {
		return false, err
	}

	return score == 1, nil
}

func isActivityCommentLikedByUserInStaticSpace(aid, cid, uid string) (bool, error) {
	exist, err := commentInfoClient.Exists(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`)).Result()
	if err != nil {
		return false, err
	}

	if exist == 0 {
		return false, nil
	}

	isLiked, err := commentInfoClient.SIsMember(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`), uid).Result()
	if err != nil {
		return false, err
	}
	return isLiked, nil
}

func isActivityCommentLikedByUserInDynamicSpace(aid, uid string) (bool, error) {
	exist, err := commentInfoClient.Exists(strings.Join([]string{`comment/activity/changed_like`, aid}, `/`)).Result()
	if err != nil {
		return false, err
	}

	if exist == 0 {
		return false, nil
	}

	score, err := commentInfoClient.ZScore(strings.Join([]string{`comment/activity/changed_like`, aid}, `/`), uid).Result()
	if err != nil {
		return false, err
	}

	return score == 1, nil
}

func IsVideoCommentLikedByUser(vid, cid, uid string) (bool, error) {
	inStaticSpace, err := isVideoCommentLikedByUserInStaticSpace(vid, cid, uid)
	if err != nil {
		return false, err
	}

	if inStaticSpace {
		return true, nil
	}

	inDynamicSpace, err := isVideoCommentLikedByUserInDynamicSpace(vid, cid, uid)
	if err != nil {
		return false, err
	}

	return inDynamicSpace, nil
}

func IsActivityCommentLikedByUser(aid, cid, uid string) (bool, error) {
	inStaticSpace, err := isActivityCommentLikedByUserInStaticSpace(aid, cid, uid)
	if err != nil {
		return false, err
	}

	if inStaticSpace {
		return true, nil
	}

	inDynamicSpace, err := isActivityCommentLikedByUserInDynamicSpace(aid, uid)
	if err != nil {
		return false, err
	}

	return inDynamicSpace, nil
}

func DeleteActivityComment(aid, cid string) error {
	tx := commentInfoClient.TxPipeline()
	tx.Del(strings.Join([]string{`comment/activity/like`, aid, cid}, `/`))
	tx.Del(strings.Join([]string{`comment/activity/changed_like`, aid, cid}, `/`))
	_, err := tx.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityCommentInfo(aid string) error {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = commentInfoClient.Scan(cursor, strings.Join([]string{`comment/activity/like`, aid, `*`}, `/`), 1000).Result()
		if err != nil {
			return err
		}
		if len(keys) > 0 {
			if _, err := commentInfoClient.Del(keys...).Result(); err != nil {
				return err
			}
		}
		if cursor == 0 {
			break
		}
	}
	return nil
}
