package redis

import (
	"github.com/go-redis/redis"
)

var (
	emailRedisClient      *redis.Client
	tokenExpireTimeClient *redis.Client
)

func Load() {
	emailRedisClient = redis.NewClient(&redis.Options{
		Addr:     EmailRedisClient.Addr,
		Password: EmailRedisClient.Password,
		DB:       EmailRedisClient.DB,
	})

	tokenExpireTimeClient = redis.NewClient(&redis.Options{
		Addr:     TokenExpireTimeClient.Addr,
		Password: TokenExpireTimeClient.Password,
		DB:       TokenExpireTimeClient.DB,
	})

	if _, err := emailRedisClient.Ping().Result(); err != nil {
		panic(err)
	}

	if _, err := tokenExpireTimeClient.Ping().Result(); err != nil {
		panic(err)
	}
	
}
