package redis

import "time"

func EmailCodeStore(email string, code string) error {
	if err := emailRedisClient.Set(email, code, 10*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func EmailCodeGet(email string) (string, error) {
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
