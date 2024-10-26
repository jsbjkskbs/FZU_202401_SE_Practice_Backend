package auth

import (
	"sfw/biz/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func Auth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		jwt.AccessTokenJwtMiddleware.MiddlewareFunc(),
	)
}
