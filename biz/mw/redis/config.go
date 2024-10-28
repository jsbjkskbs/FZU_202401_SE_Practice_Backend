package redis

type _Redis struct {
	Addr     string
	Password string
	DB       int
}

var (
	EmailRedisClient      _Redis
	TokenExpireTimeClient _Redis
	VideoClient           _Redis
	VideoInfoClient       _Redis
)
