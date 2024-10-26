package redis

type _Redis struct {
	Addr     string
	Password string
	DB       int
}

var (
	EmailCodeRedisClient _Redis
)
