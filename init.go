package main

import (
	"context"
	"fmt"
	"runtime"
	"sfw/biz/dal"
	"sfw/biz/mw/generator/snowflake"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/mw/sentinel"
	"sfw/pkg/errno"
	"sfw/pkg/utils/configure"
	"sfw/pkg/utils/mail"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/opensergo/sentinel/adapter"
)

func InstallSentinel(h *server.Hertz) {
	h.Use(
		adapter.SentinelServerMiddleware(
			adapter.WithServerResourceExtractor(func(ctx context.Context, c *app.RequestContext) string {
				return "default"
			}),
			adapter.WithServerBlockFallback(func(ctx context.Context, c *app.RequestContext) {
				c.AbortWithStatusJSON(consts.StatusTooManyRequests, utils.H{
					"code": errno.QueryLimit.Code,
					"msg":  errno.QueryLimit.Message,
				})
			}),
		),
	)
}

func Initialize() {
	checkEnv()
	snowflake.Init()
	jwt.AccessTokenJwtInit()
	jwt.RefreshTokenJwtInit()

	// Initialize your application here.
	configureLoader := configure.NewConfLoader(&configure.ConfigureOption{
		ConfigName:    "config",
		ConfigType:    "yaml",
		ConfigPath:    ".",
		RegisterParam: []interface{}{},
		Register:      ConfigureRegister,
		LogPrefix:     "|Config Loader|",
		LogSuffix:     "",
		LogFunc:       hlog.Info,
		WarnFunc:      hlog.Warn,
		ErrorFunc:     hlog.Error,
		FatalFunc:     hlog.Fatal,
		Silent:        false,
	})
	if err := configureLoader.Run(); err != nil {
		hlog.Fatal("|Config Loader|", err)
	}
}

func ConfigureRegister(...any) {
	configure.RuleTable = []*configure.ConfigureRule{
		{
			RuleName: "mysql",
			Level:    configure.LevelFatal,

			LoadMethodParam: []interface{}{},
			LoadMethod: func(v ...any) error {
				cmap := configure.GlobalConfig.GetStringMap("MySQL")
				ok := false
				if dal.DSN, ok = cmap["dsn"].(string); !ok {
					return errno.InternalServerError
				}
				dal.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger:      func(v ...any) {},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
		{
			RuleName: "gorse",
			Level:    configure.LevelFatal,

			LoadMethodParam: []interface{}{},
			LoadMethod: func(v ...any) error {
				cmap := configure.GlobalConfig.GetStringMap("Gorse")
				ok := false
				if gorse.Url, ok = cmap["url"].(string); !ok {
					return errno.InternalServerError
				}
				if gorse.ApiKey, ok = cmap["apikey"].(string); !ok {
					return errno.InternalServerError
				}
				gorse.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger:      func(v ...any) {},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
		{
			RuleName: "Email",
			Level:    configure.LevelFatal,

			LoadMethodParam: []interface{}{},
			LoadMethod: func(v ...any) error {
				cmap := configure.GlobalConfig.GetStringMap("Email")
				ok := false
				mail.Config = new(mail.EmailStationConfig)
				if mail.Config.Address, ok = cmap["address"].(string); !ok {
					return errno.InternalServerError
				}
				if mail.Config.Port, ok = cmap["port"].(int); !ok {
					return errno.InternalServerError
				}
				if mail.Config.Username, ok = cmap["username"].(string); !ok {
					return errno.InternalServerError
				}
				if mail.Config.Password, ok = cmap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if mail.Config.ConnPoolSize, ok = cmap["conn_pool_size"].(int); !ok {
					return errno.InternalServerError
				}
				mail.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger:      func(v ...any) {},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
		{
			RuleName: "Redis",
			Level:    configure.LevelFatal,

			LoadMethodParam: []interface{}{},
			LoadMethod: func(v ...any) error {
				cmap := configure.GlobalConfig.GetStringMap("Redis")
				emap, ok := cmap["email"].(map[string]interface{})
				if !ok {
					return errno.InternalServerError
				}
				if redis.EmailCodeRedisClient.Addr, ok = emap["addr"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.EmailCodeRedisClient.Password, ok = emap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.EmailCodeRedisClient.DB, ok = emap["db"].(int); !ok {
					return errno.InternalServerError
				}
				redis.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger:      func(v ...any) {},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
		{
			RuleName: "Sentinel",
			Level:    configure.LevelFatal,

			LoadMethodParam: []interface{}{},
			LoadMethod: func(v ...any) error {
				sentinel.Rules = configure.GlobalConfig.GetStringMap("Sentinel")
				sentinel.Load()
				return nil
			},

			SuccessTriggerParam: []interface{}{},
			SuccessTrigger:      func(v ...any) {},

			FailedTriggerParam: []interface{}{},
			FailedTrigger:      func(v ...any) {},
		},
	}
}

func checkEnv() {
	if !(runtime.GOOS == "linux" || runtime.GOOS == "darwin") {
		panic(
			fmt.Sprint(
				" Do you want to run the server on a non-linux or non-darwin platform? 😅\n",
				"Hhmmm... 🤔\n",
				"I bet you're a crazy follower of "+runtime.GOOS+"! 🤓👆",
			))
	}
}
