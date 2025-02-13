package redis

import "time"

func VideoUploadInfoStore(key string, kv map[string]interface{}, ttl time.Duration) error {
	tx := videoClient.TxPipeline()
	tx.HMSet(key, kv)
	tx.Expire(key, ttl)
	if _, err := tx.Exec(); err != nil {
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
