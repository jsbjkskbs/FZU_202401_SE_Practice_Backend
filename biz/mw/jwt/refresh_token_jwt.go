package jwt

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
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
func GetRefreshTokenExpireAt(ctx context.Context, c *app.RequestContext) time.Time {
	claims, _ := RefreshTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	switch v := claims["exp"].(type) {
	case nil:
		return time.Time{}
	case float64:
		return time.Unix(int64(v), 0)
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return time.Time{}
		}
		return time.Unix(n, 0)
	default:
		return time.Time{}
	}
}

// GenerateRefreshToken 生成刷新Token
func GenerateRefreshToken(ctx context.Context, c *app.RequestContext) string {
	v, _ := c.Get(AccessTokenJwtMiddleware.IdentityKey)
	data := PayloadIdentityData{
		Uid: v.(*PayloadIdentityData).Uid,
	}
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(data)
	return tokenString
}
