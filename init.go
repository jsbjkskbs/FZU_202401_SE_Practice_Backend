package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/mw/sentinel"
	"sfw/pkg/errno"
	"sfw/pkg/oss"
	"sfw/pkg/synchronizer"
	"sfw/pkg/utils/checker"
	"sfw/pkg/utils/configure"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/mail"
	"sfw/pkg/utils/scheduler"
	"time"

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
	generator.Init()
	scheduler.Init()
	jwt.Init()

	// Initialize your application here.
	configureLoader := configure.NewConfLoader(&configure.ConfigureOption{
		ConfigName:    "config",
		ConfigType:    "yaml",
		ConfigPath:    ".",
		RegisterParam: []interface{}{},
		Register:      ConfigureRegister,
		LogPrefix:     "Config Loader: ",
		LogSuffix:     "",
		LogFunc:       hlog.Info,
		WarnFunc:      hlog.Warn,
		ErrorFunc:     hlog.Error,
		FatalFunc:     hlog.Fatal,
		Silent:        false,
	})
	if err := configureLoader.Run(); err != nil {
		hlog.Fatal("Config Loader: ", err)
	}
	loadCategory()

	err := synchronizer.SynchronizeVideoVisitInfoDB2Redis()
	if err != nil {
		hlog.Fatal("Synchronize Task: synchronize video visit info from db to redis error", err)
	}
	hlog.Info("Synchronize Task: sychronize video visit info from db to redis success")

	err = synchronizer.SynchronizeVideoLikeFromDB2Redis()
	if err != nil {
		hlog.Fatal("Synchronize Task: synchronize video like from db to redis error", err)
	}
	hlog.Info("Synchronize Task: sychronize video like from db to redis success")

	err = synchronizer.SynchronizeActivityLikeFromDB2Redis()
	if err != nil {
		hlog.Fatal("Synchronize Task: synchronize activity like from db to redis error", err)
	}
	hlog.Info("Synchronize Task: sychronize activity like from db to redis success")

	err = synchronizer.SynchronizeVideoCommentLikeFromDB2Redis()
	if err != nil {
		hlog.Fatal("Synchronize Task: synchronize video comment like from db to redis error", err)
	}
	hlog.Info("Synchronize Task: sychronize video comment like from db to redis success")

	err = synchronizer.SynchronizeActivityCommentLikeFromDB2Redis()
	if err != nil {
		hlog.Fatal("Synchronize Task: synchronize activity comment like from db to redis error", err)
	}
	hlog.Info("Synchronize Task: sychronize activity comment like from db to redis success")
	hlog.Info("Synchronize Task: all synchronize task success")

	hlog.Info("Initialize success, ready to serve after 3 seconds")
	time.Sleep(3 * time.Second)
}

func loadCategory() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				c := dal.Executor.Category
				list := []model.Category{}
				err := c.WithContext(context.Background()).Scan(&list)
				if err != nil {
					hlog.Error("load category error", err)
				}
				for _, v := range list {
					checker.CategoryMap[v.CategoryName] = v.ID
				}
				checker.Categories = make([]string, len(checker.CategoryMap)+1)
				for k, v := range checker.CategoryMap {
					checker.Categories[v] = k
				}
				checker.Categories = checker.Categories[1:]
				hlog.Infof("Synchronizer: category loaded success[%v]", checker.CategoryMap)
				ticker.Reset(1 * time.Hour)
			}
		}
	}()
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
					return errors.New("mysql dsn not found")
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
				if redis.EmailRedisClient.Addr, ok = emap["addr"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.EmailRedisClient.Password, ok = emap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.EmailRedisClient.DB, ok = emap["db"].(int); !ok {
					return errno.InternalServerError
				}

				tmap, ok := cmap["token_expire_time"].(map[string]interface{})
				if !ok {
					return errno.InternalServerError
				}
				if redis.TokenExpireTimeClient.Addr, ok = tmap["addr"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.TokenExpireTimeClient.Password, ok = tmap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.TokenExpireTimeClient.DB, ok = tmap["db"].(int); !ok {
					return errno.InternalServerError
				}

				vmap, ok := cmap["video"].(map[string]interface{})
				if !ok {
					return errno.InternalServerError
				}
				if redis.VideoClient.Addr, ok = vmap["addr"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.VideoClient.Password, ok = vmap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.VideoClient.DB, ok = vmap["db"].(int); !ok {
					return errno.InternalServerError
				}

				vimap, ok := cmap["video_info"].(map[string]interface{})
				if !ok {
					return errno.InternalServerError
				}
				if redis.VideoInfoClient.Addr, ok = vimap["addr"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.VideoInfoClient.Password, ok = vimap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.VideoInfoClient.DB, ok = vimap["db"].(int); !ok {
					return errno.InternalServerError
				}

				aimap, ok := cmap["activity_info"].(map[string]interface{})
				if !ok {
					return errno.InternalServerError
				}
				if redis.ActivityInfoClient.Addr, ok = aimap["addr"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.ActivityInfoClient.Password, ok = aimap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.ActivityInfoClient.DB, ok = aimap["db"].(int); !ok {
					return errno.InternalServerError
				}

				cimap, ok := cmap["comment_info"].(map[string]interface{})
				if !ok {
					return errno.InternalServerError
				}
				if redis.CommentInfoClient.Addr, ok = cimap["addr"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.CommentInfoClient.Password, ok = cimap["password"].(string); !ok {
					return errno.InternalServerError
				}
				if redis.CommentInfoClient.DB, ok = cimap["db"].(int); !ok {
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
		{
			RuleName: "OSS",
			Level:    configure.LevelFatal,

			LoadMethodParam: []interface{}{},
			LoadMethod: func(v ...any) error {
				cmap := configure.GlobalConfig.GetStringMap("OSS")
				ok := false
				if oss.Bucket, ok = cmap["bucket"].(string); !ok {
					return errno.InternalServerError
				}
				if oss.AccessKey, ok = cmap["access_key"].(string); !ok {
					return errno.InternalServerError
				}
				if oss.SecretKey, ok = cmap["secret_key"].(string); !ok {
					return errno.InternalServerError
				}
				if oss.Domain, ok = cmap["domain"].(string); !ok {
					return errno.InternalServerError
				}
				if oss.CallbackUrl, ok = cmap["callback_url"].(string); !ok {
					return errno.InternalServerError
				}
				if oss.UploadUrl, ok = cmap["upload_url"].(string); !ok {
					return errno.InternalServerError
				}
				oss.Load()
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
