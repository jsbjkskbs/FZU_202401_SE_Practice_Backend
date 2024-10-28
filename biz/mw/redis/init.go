package redis

import (
	"github.com/go-redis/redis"
)

var (
	emailRedisClient      *redis.Client
	tokenExpireTimeClient *redis.Client
	videoClient           *redis.Client
	videoInfoClient       *redis.Client
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

	videoClient = redis.NewClient(&redis.Options{
		Addr:     VideoClient.Addr,
		Password: VideoClient.Password,
		DB:       VideoClient.DB,
	})

	videoInfoClient = redis.NewClient(&redis.Options{
		Addr:     VideoInfoClient.Addr,
		Password: VideoInfoClient.Password,
		DB:       VideoInfoClient.DB,
	})

	if _, err := emailRedisClient.Ping().Result(); err != nil {
		panic(err)
	}

	if _, err := tokenExpireTimeClient.Ping().Result(); err != nil {
		panic(err)
	}

	if _, err := videoClient.Ping().Result(); err != nil {
		panic(err)
	}

	if _, err := videoInfoClient.Ping().Result(); err != nil {
		panic(err)
	}

}
