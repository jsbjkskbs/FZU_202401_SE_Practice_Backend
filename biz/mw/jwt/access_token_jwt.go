package jwt

import (
	"context"
	"encoding/json"
	"sfw/pkg/errno"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// GenerateAccessToken 生成AccessToken
func GenerateAccessToken(ctx context.Context, c *app.RequestContext) string {
	v, _ := c.Get(RefreshTokenJwtMiddleware.IdentityKey)
	data := PayloadIdentityData{
		Uid: v.(*PayloadIdentityData).Uid,
	}
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(data)
	return tokenString
}

// IsAccessTokenAvailable AccessToken是否有效
func IsAccessTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := AccessTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := AccessTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(AccessTokenJwtMiddleware.IdentityKey, identity)
	}
	if !AccessTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}
	return true
}

// GetAccessTokenExpireAt 获取AccessToken过期时间
func GetAccessTokenExpireAt(ctx context.Context, c *app.RequestContext) time.Time {
	claims, _ := AccessTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
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

// ExtractUserIdWhenAuthorized 提取用户ID
func ExtractUserIdWhenAuthorized(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	data, exist := c.Get(AccessTokenJwtMiddleware.IdentityKey)
	if !exist {
		return nil, errno.AccessTokenInvalid
	}
	return data, nil
}

// CovertJWTPayloadToString 将JWT Payload转换为字符串
func CovertJWTPayloadToString(ctx context.Context, c *app.RequestContext) (string, error) {
	data, err := ExtractUserIdWhenAuthorized(ctx, c)
	if err != nil {
		return ``, err
	}
	return data.(map[string]interface{})["Uid"].(string), nil
}
