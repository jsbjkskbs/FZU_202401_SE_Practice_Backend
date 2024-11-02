package jwt

import (
	"time"
)

var (
	// AccessTokenJwtMiddleware AccessToken JWT中间件
	AccessTokenJwtMiddleware *JWTMiddleware
	// RefreshTokenJwtMiddleware RefreshToken JWT中间件
	RefreshTokenJwtMiddleware *JWTMiddleware

	// AccessTokenExpireTime AccessToken过期时间
	AccessTokenExpireTime = time.Hour * 4
	// RefreshTokenExpireTime RefreshToken过期时间
	RefreshTokenExpireTime = time.Hour * 72

	// AccessTokenIdentityKey AccessToken身份标识
	AccessTokenIdentityKey = "access_token_field"
	// RefreshTokenIdentityKey RefreshToken身份标识
	RefreshTokenIdentityKey = "refresh_token_field"

	AccessTokenKey  = []byte("access_token_key")
	RefreshTokenKey = []byte("refresh_token_key")
)

func Init() {
	AccessTokenJwtMiddleware = NewJWTMiddleware(JWTMiddlewareOptions{
		IdentityKey:                 AccessTokenIdentityKey,
		Timeout:                     AccessTokenExpireTime,
		Key:                         AccessTokenKey,
		TookenLookup:                "query:Access-Token,header:Access-Token",
		WithoutDefaultTokenHeadName: true,
	})
	RefreshTokenJwtMiddleware = NewJWTMiddleware(JWTMiddlewareOptions{
		IdentityKey:                 RefreshTokenIdentityKey,
		Timeout:                     RefreshTokenExpireTime,
		Key:                         RefreshTokenKey,
		TookenLookup:                "query:Refresh-Token,header:Refresh-Token",
		WithoutDefaultTokenHeadName: true,
	})
	err := AccessTokenJwtMiddleware.Init()
	if err != nil {
		panic(err)
	}
	err = RefreshTokenJwtMiddleware.Init()
	if err != nil {
		panic(err)
	}
}
