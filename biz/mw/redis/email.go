package redis

import (
	"time"
)

func EmailCodeStore(email string, code string) error {
	if err := emailRedisClient.Set(email, code, 10*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func EmailCodeTTL(email string) (time.Duration, error) {
	if exist, err := emailRedisClient.Exists(email).Result(); err != nil {
		return 0, err
	} else if exist == 0 {
		return 0, nil
	}

	ttl, err := emailRedisClient.TTL(email).Result()
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

func EmailCodeGet(email string) (string, error) {
	exist, err := emailRedisClient.Exists(email).Result()
	if err != nil {
		return "", err
	} else if exist == 0 {
		return "", nil
	}

	code, err := emailRedisClient.Get(email).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

func EmailCodeDel(email string) error {
	if err := emailRedisClient.Del(email).Err(); err != nil {
		return err
	}
	return nil
}
