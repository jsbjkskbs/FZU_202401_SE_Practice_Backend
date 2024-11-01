package jwt

import (
	"context"
	"encoding/json"
	"sfw/pkg/errno"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	gojwt "github.com/golang-jwt/jwt/v4"
)

// IsRefreshTokenAvailable 刷新Token是否有效
func IsRefreshTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := RefreshTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := RefreshTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(RefreshTokenJwtMiddleware.IdentityKey, identity)
	}
	if !RefreshTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}
	return true
}

// GetRefreshTokenExpireAt 获取刷新Token过期时间
func GetRefreshTokenExpireAt(token string) (time.Time, error) {
	data, err := AccessTokenJwtMiddleware.ParseTokenString(token)
	if err != nil {
		return time.Time{}, err
	}
	exp, ok := data.Claims.(gojwt.MapClaims)["exp"]
	if !ok {
		return time.Time{}, errno.AccessTokenInvalid
	}
	switch v := exp.(type) {
	case float64:
		return time.Unix(int64(v), 0), nil
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(n, 0), nil
	default:
		return time.Time{}, errno.AccessTokenInvalid
	}
}

// GenerateRefreshToken 生成刷新Token
func GenerateRefreshToken(uid string) string {
	tokenString, _, _ := RefreshTokenJwtMiddleware.TokenGenerator(PayloadIdentityData{Uid: uid})
	return tokenString
}
