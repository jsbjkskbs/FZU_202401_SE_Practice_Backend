package jwt

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/model/api/user"
	"sfw/pkg/errno"
	"sfw/pkg/utils"
	"sfw/pkg/utils/encrypt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var (
	// AccessTokenJwtMiddleware AccessToken JWT中间件
	AccessTokenJwtMiddleware *jwt.HertzJWTMiddleware
	// RefreshTokenJwtMiddleware RefreshToken JWT中间件
	RefreshTokenJwtMiddleware *jwt.HertzJWTMiddleware

	// AccessTokenExpireTime AccessToken过期时间
	AccessTokenExpireTime = time.Hour * 1
	// RefreshTokenExpireTime RefreshToken过期时间
	RefreshTokenExpireTime = time.Hour * 72

	// AccessTokenIdentityKey AccessToken身份标识
	AccessTokenIdentityKey = "access_token_field"
	// RefreshTokenIdentityKey RefreshToken身份标识
	RefreshTokenIdentityKey = "refresh_token_field"
)

// PayloadIdentityData 载荷身份数据
type PayloadIdentityData struct {
	// Uid 用户ID
	Uid string
}

// AccessTokenJwtInit 初始化AccessToken JWT
func AccessTokenJwtInit() {
	var err error
	AccessTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:                         []byte("access_token_key_123456"),
		TokenLookup:                 "query:Access-Token,header:Access-Token",
		Timeout:                     AccessTokenExpireTime,
		IdentityKey:                 AccessTokenIdentityKey,
		WithoutDefaultTokenHeadName: true,

		// this func would never be called now
		// use GenerateAccessToken instead
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest user.UserLoginReq
			if err := c.BindAndValidate(&loginRequest); err != nil {
				return nil, err
			}
			u := dal.Executor.User
			user, err := u.WithContext(context.Background()).Where(u.Username.Eq(loginRequest.Username)).First()
			if err != nil {
				return nil, err
			}
			if user.Password != encrypt.EncryptBySHA256WithSalt(loginRequest.Password, encrypt.GetSalt()) {
				return nil, errno.AccountOrPasswordInvalid
			}
			c.Set("user_id", user.ID)
			return PayloadIdentityData{Uid: fmt.Sprint(user.ID)}, nil
		},

		// data为Authenticator返回的interface{}
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(PayloadIdentityData); ok {
				return jwt.MapClaims{
					AccessTokenJwtMiddleware.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
		},

		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return utils.CreateBaseHttpResponse(e).Msg
		},
	})

	if err != nil {
		panic(err)
	}
}

// RefreshTokenJwtInit 初始化RefreshToken JWT
func RefreshTokenJwtInit() {
	var err error
	RefreshTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:                         []byte("refresh_token_key_abcdef"),
		TokenLookup:                 "query:Refresh-Token,header:Refresh-Token",
		Timeout:                     RefreshTokenExpireTime,
		IdentityKey:                 RefreshTokenIdentityKey,
		WithoutDefaultTokenHeadName: true,

		// this func would never be called now
		// use GenerateAccessToken instead
		// 只在LoginHandler触发
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			uid, exist := c.Get("user_id")
			if !exist {
				return nil, errno.RefreshTokenInvalid
			}
			return PayloadIdentityData{Uid: fmt.Sprint(uid)}, nil
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &PayloadIdentityData{
				Uid: claims[RefreshTokenJwtMiddleware.IdentityKey].(string),
			}
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(PayloadIdentityData); ok {
				return jwt.MapClaims{
					RefreshTokenJwtMiddleware.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
		},

		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			return true
		},
	})
	if err != nil {
		panic(err)
	}
}
