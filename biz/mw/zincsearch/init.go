package zincsearch

var (
	Client    *ZincClient
	ClientOpt ZincClientOption
)

func Load() {
	Client = NewZincClient(&ClientOpt)
	if err := Client.Ping(); err != nil {
		panic(err)
	}
}
