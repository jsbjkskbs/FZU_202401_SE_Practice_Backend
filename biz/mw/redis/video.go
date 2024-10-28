package redis

import "github.com/go-redis/redis"

func VideoUploadInfoStore(key string, kv map[string]interface{}) error {
	if err := videoClient.HMSet(key, kv).Err(); err != nil {
		return err
	}
	return nil
}

func VideoUploadInfoGet(key string) (map[string]string, error) {
	if exist, err := videoClient.Exists(key).Result(); err != nil {
		return nil, err
	} else if exist == 0 {
		return nil, nil
	}

	info, err := videoClient.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func VideoUploadInfoExist(key string) (bool, error) {
	if exist, err := videoClient.Exists(key).Result(); err != nil {
		return false, err
	} else if exist == 0 {
		return false, nil
	}
	return true, nil
}

func VideoUploadInfoDel(key string) error {
	if err := videoClient.Del(key).Err(); err != nil {
		return err
	}
	return nil
}

func VideoRankJoin(key, category string) error {
	if err := videoClient.ZAdd("video_rank", redis.Z{Score: 0, Member: key}).Err(); err != nil {
		return err
	}
	if err := videoClient.ZAdd("video_rank_"+category, redis.Z{Score: 0, Member: key}).Err(); err != nil {
		return err
	}
	return nil
}

func VideoScoreAdd(key, category string, score float64) error {
	if err := videoClient.ZIncrBy("video_rank", score, key).Err(); err != nil {
		return err
	}
	if err := videoClient.ZIncrBy("video_rank_"+category, score, key).Err(); err != nil {
		return err
	}
	return nil
}

func VideoRankGetTopN(n int64) ([]string, error) {
	topN, err := videoClient.ZRevRange("video_rank", 0, n-1).Result()
	if err != nil {
		return nil, err
	}
	return topN, nil
}

func VideoRankGetTopNWithCategory(n int64, category string) ([]string, error) {
	topN, err := videoClient.ZRevRange("video_rank_"+category, 0, n-1).Result()
	if err != nil {
		return nil, err
	}
	return topN, nil
}

func VideoRankRemove(key, category string) error {
	if err := videoClient.ZRem("video_rank", key).Err(); err != nil {
		return err
	}
	if err := videoClient.ZRem("video_rank_"+category, key).Err(); err != nil {
		return err
	}
	return nil
}
