package gorse

import (
	"context"
	"sfw/pkg/errno"

	"github.com/zhenghaoz/gorse/client"
)

var (
	Cli    *client.GorseClient
	Url    = ""
	ApiKey = ""
)

// ping checks the connection to the Gorse server.
// ping 检查与Gorse服务器的连接。
func ping() (err error) {
	if Cli == nil {
		return errno.InternalServerError
	}
	ctx := context.TODO()
	_, err = Cli.GetRecommend(ctx, "test", "", 10)
	return
}

// Load loads the configuration of Gorse.
// It will panic if the connection to the Gorse server is failed.
// It is recommended to call this function in the init function of the package.
// Load 加载Gorse的配置。
// 如果连接到Gorse服务器失败，它将引发panic。
// 建议在包的init函数中调用此函数。
func Load() {
	Cli = client.NewGorseClient(Url, ApiKey)
	if err := ping(); err != nil {
		panic(err)
	}
}
