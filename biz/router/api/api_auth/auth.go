package auth

import (
	"context"

	"sfw/biz/dal/exquery"
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
	)
}

func checkTokenVaild(ctx context.Context, c *app.RequestContext) {
	token := string(c.GetHeader("Access-Token"))
	payload, expire, err := jwt.AccessTokenJwtMiddleware.GetBasicDataFromToken(token)
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

func AdminAuth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		checkTokenVaild,
		checkAdmin,
	)
}

func checkAdmin(ctx context.Context, c *app.RequestContext) {
	token := string(c.GetHeader("Access-Token"))
	uid, _ := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(token)
	u, err := exquery.QueryUserByID(uid)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.DatabaseCallError.Code,
			"msg":  errno.DatabaseCallError.Message,
		})
		c.Abort()
		return
	}
	if u.Role != "admin" {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.PowerNotEnough.Code,
			"msg":  errno.PowerNotEnough.Message,
		})
		c.Abort()
		return
	}
}
