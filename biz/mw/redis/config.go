package redis

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

var (
	EmailRedisClient      RedisConf
	TokenExpireTimeClient RedisConf
	VideoClient           RedisConf
	VideoInfoClient       RedisConf
	ActivityInfoClient    RedisConf
	CommentInfoClient     RedisConf
)
