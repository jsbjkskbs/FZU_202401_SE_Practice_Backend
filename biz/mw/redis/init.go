package redis

import (
	"github.com/go-redis/redis"
)

var (
	emailRedisClient *redis.Client
)

func Load() {
	emailRedisClient = redis.NewClient(&redis.Options{
		Addr:     EmailCodeRedisClient.Addr,
		Password: EmailCodeRedisClient.Password,
		DB:       EmailCodeRedisClient.DB,
	})

	if _, err := emailRedisClient.Ping().Result(); err != nil {
		panic(err)
	}
}
