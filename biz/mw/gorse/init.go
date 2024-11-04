package gorse

import (
	"context"

	"sfw/pkg/errno"

	"github.com/zhenghaoz/gorse/client"
)

var (
	cli    *client.GorseClient
	Url    = ""
	ApiKey = ""
)

// ping checks the connection to the Gorse server.
// ping 检查与Gorse服务器的连接。
func ping() error {
	if cli == nil {
		return errno.InternalServerError
	}
	health, err := cli.HealthLive(context.Background())
	if err != nil {
		return err
	}
	if !health.Ready {
		return errno.InternalServerError
	}
	return nil
}

// Load loads the configuration of Gorse.
// It will panic if the connection to the Gorse server is failed.
// It is recommended to call this function in the init function of the package.
// Load 加载Gorse的配置。
// 如果连接到Gorse服务器失败，它将引发panic。
// 建议在包的init函数中调用此函数。
func Load() {
	cli = client.NewGorseClient(Url, ApiKey)
	if err := ping(); err != nil {
		panic(err)
	}
}
