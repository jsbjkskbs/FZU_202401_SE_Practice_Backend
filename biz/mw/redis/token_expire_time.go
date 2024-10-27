package redis

import "time"

func TokenExpireTimeStore(key string, expire int64, ttl time.Duration) error {
	if err := tokenExpireTimeClient.Set(key, expire, ttl).Err(); err != nil {
		return err
	}
	return nil
}

func TokenExpireTimeGet(key string) (int64, error) {
	if exist, err := tokenExpireTimeClient.Exists(key).Result(); err != nil {
		return 0, err
	} else if exist == 0 {
		return 0, nil
	}

	expire, err := tokenExpireTimeClient.Get(key).Int64()
	if err != nil {
		return 0, err
	}
	return expire, nil
}

func TokenExpireTimeDel(key string) error {
	if err := tokenExpireTimeClient.Del(key).Err(); err != nil {
		return err
	}
	return nil
}
