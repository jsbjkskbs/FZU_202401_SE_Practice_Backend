package zincsearch

var (
	Client    *Logger
	ClientOpt LoggerOption
)

func Load() {
	Client = NewLogger(&ClientOpt)
	if err := Client.Ping(); err != nil {
		panic(err)
	}
}
