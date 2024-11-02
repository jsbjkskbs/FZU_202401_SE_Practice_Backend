package jwt

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	gojwt "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
)

type JWTMiddlewareOptions struct {
	IdentityKey                 string
	Timeout                     time.Duration
	Key                         []byte
	TookenLookup                string
	WithoutDefaultTokenHeadName bool
}

type JWTMiddleware struct {
	jwt          *jwt.HertzJWTMiddleware
	identityKey  string
	invaildError error
}

type PayloadData struct {
	Payload string
}

func NewJWTMiddleware(opts JWTMiddlewareOptions) *JWTMiddleware {
	return &JWTMiddleware{
		jwt: &jwt.HertzJWTMiddleware{
			IdentityKey:                 opts.IdentityKey,
			Timeout:                     opts.Timeout,
			Key:                         opts.Key,
			TokenLookup:                 opts.TookenLookup,
			WithoutDefaultTokenHeadName: opts.WithoutDefaultTokenHeadName,

			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(PayloadData); ok {
					return jwt.MapClaims{
						opts.IdentityKey: v,
					}
				}
				return jwt.MapClaims{}
			},
		},
		identityKey:  opts.IdentityKey,
		invaildError: errors.New(`invalid token`),
	}
}

func (mw *JWTMiddleware) Init() error {
	if mw.jwt == nil {
		return errors.New(`jwt middleware not initialized`)
	}
	var err error
	mw.jwt, err = jwt.New(mw.jwt)
	if err != nil {
		return err
	}
	return nil
}

func (mw *JWTMiddleware) GenerateToken(payload string) string {
	tokenString, _, _ := mw.jwt.TokenGenerator(PayloadData{Payload: payload})
	return tokenString
}

func (mw *JWTMiddleware) IsTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := mw.jwt.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < mw.jwt.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < mw.jwt.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := mw.jwt.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(mw.jwt.IdentityKey, identity)
	}
	if !mw.jwt.Authorizator(identity, ctx, c) {
		return false
	}
	return true
}

func (mw *JWTMiddleware) GetTokenExpireAt(token string) (time.Time, error) {
	_, expiredAt, err := mw.GetBasicDataFromToken(token)
	return expiredAt, err
}

func (mw *JWTMiddleware) ExtractPayloadFromToken(token string) (string, error) {
	payload, _, err := mw.GetBasicDataFromToken(token)
	if err != nil {
		return ``, err
	}
	p, ok := payload.(string)
	if !ok {
		return ``, errors.New(`payload type error`)
	}
	return p, nil
}

func (mw *JWTMiddleware) ConvertJWTPayloadToInt64(token string) (int64, error) {
	payload, err := mw.ExtractPayloadFromToken(token)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(payload, 10, 64)
}

func (mw *JWTMiddleware) GetBasicDataFromToken(token string) (interface{}, time.Time, error) {
	data, err := mw.jwt.ParseTokenString(token)
	if err != nil {
		return ``, time.Time{}, err
	}
	expiredAt := time.Time{}
	exp, ok := data.Claims.(gojwt.MapClaims)["exp"]
	if !ok {
		return ``, time.Time{}, errors.New(`exp not found`)
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
		return ``, time.Time{}, errors.New(`exp type error`)
	}
	payload, ok := data.Claims.(gojwt.MapClaims)[mw.identityKey].(map[string]interface{})["Payload"]
	if !ok {
		return ``, time.Time{}, errors.New(`payload not found`)
	}
	return payload, expiredAt, nil
}
