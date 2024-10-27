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
		jwt.AccessTokenJwtMiddleware.MiddlewareFunc(),
		checkTokenExpireTime,
	)
}

func checkTokenExpireTime(ctx context.Context, c *app.RequestContext) {
	uid, err := jwt.CovertJWTPayloadToString(ctx, c)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.AccessTokenInvalid.Code,
			"msg":  errno.AccessTokenInvalid.Message,
		})
		c.Abort()
		return
	}
	tokenExpireTime := jwt.GetAccessTokenExpireAt(ctx, c).Add(-jwt.AccessTokenExpireTime).Unix()
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
