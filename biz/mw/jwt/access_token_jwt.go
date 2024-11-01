package jwt

import (
	"context"
	"encoding/json"
	"sfw/pkg/errno"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	gojwt "github.com/golang-jwt/jwt/v4"
)

// GenerateAccessToken 生成AccessToken
func GenerateAccessToken(uid string) string {
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(PayloadIdentityData{Uid: uid})
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
func GetAccessTokenExpireAt(token string) (time.Time, error) {
	_, expiredAt, err := GetBasicDataFromAccessToken(token)
	return expiredAt, err
}

func ExtractUserIdFromAccessToken(token string) (string, error) {
	payload, _, err := GetBasicDataFromAccessToken(token)
	if err != nil {
		return ``, err
	}
	uid, ok := payload.(string)
	if !ok {
		return ``, errno.AccessTokenInvalid
	}
	return uid, nil
}

func ConvertJWTPayloadToInt64(token string) (int64, error) {
	uid, err := ExtractUserIdFromAccessToken(token)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(uid, 10, 64)
}

func GetBasicDataFromAccessToken(token string) (interface{}, time.Time, error) {
	data, err := AccessTokenJwtMiddleware.ParseTokenString(token)
	if err != nil {
		return ``, time.Time{}, err
	}
	expiredAt := time.Time{}
	exp, ok := data.Claims.(gojwt.MapClaims)["exp"]
	if !ok {
		return ``, time.Time{}, errno.AccessTokenInvalid
	}
	switch v := exp.(type) {
	case float64:
		expiredAt = time.Unix(int64(v), 0)
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return ``, time.Time{}, err
		}
		expiredAt = time.Unix(n, 0)
	default:
		return ``, time.Time{}, errno.AccessTokenInvalid
	}
	uid, ok := data.Claims.(gojwt.MapClaims)[AccessTokenIdentityKey].(map[string]interface{})["Uid"]
	if !ok {
		return ``, time.Time{}, errno.AccessTokenInvalid
	}
	return uid, expiredAt, nil
}
