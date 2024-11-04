package webs

import (
	"context"

	"sfw/biz/mw/jwt"
	"sfw/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func _wsAuth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		tokenAuthFunc(),
	)
}

func tokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if !jwt.AccessTokenJwtMiddleware.IsTokenAvailable(ctx, c) {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, utils.H{
				"code": errno.AccessTokenInvalidErrorCode,
				"msg":  errno.AccessTokenInvalidErrorMsg,
			})
			return
		}
		c.Next(ctx)
	}
}
