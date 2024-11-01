package auth

import (
	"context"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Auth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		checkTokenVaild,
		checkTokenExpireTime,
	)
}

func checkTokenVaild(ctx context.Context, c *app.RequestContext) {
	if !jwt.IsAccessTokenAvailable(ctx, c) {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.AccessTokenInvalid.Code,
			"msg":  errno.AccessTokenInvalid.Message,
		})
		c.Abort()
	}
}

func checkTokenExpireTime(ctx context.Context, c *app.RequestContext) {
	token := string(c.GetHeader("Access-Token"))
	payload, expire, err := jwt.GetBasicDataFromAccessToken(token)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.AccessTokenInvalid.Code,
			"msg":  errno.AccessTokenInvalid.Message,
		})
		c.Abort()
		return
	}
	uid, ok := payload.(string)
	if !ok {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.AccessTokenInvalid.Code,
			"msg":  errno.AccessTokenInvalid.Message,
		})
		c.Abort()
		return
	}
	tokenExpireTime := expire.Add(-jwt.AccessTokenExpireTime).Unix()
	latestExpireTime, err := redis.TokenExpireTimeGet(uid)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.DatabaseCallError.Code,
			"msg":  errno.DatabaseCallError.Message,
		})
		c.Abort()
		return
	}

	if latestExpireTime == 0 {
		return
	}

	if tokenExpireTime <= latestExpireTime {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.AccessTokenForbidden.Code,
			"msg":  errno.AccessTokenForbidden.Message,
		})
		c.Abort()
		return
	}
}
