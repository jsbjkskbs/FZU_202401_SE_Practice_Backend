package service

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/user"
	"sfw/biz/mw/generator/snowflake"
	"sfw/biz/mw/redis"
	"sfw/pkg/errno"
	"sfw/pkg/utils/encrypt"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/mail"
	"sfw/pkg/utils/mfa"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c:   c,
	}
}

func (service *UserService) NewRegisterEvent(req *user.UserRegisterReq) error {
	code, err := redis.EmailCodeGet(req.Email)
	if err != nil {
		return errno.DatabaseCallError
	}
	if code != req.Code {
		return errno.CustomError.WithMessage("验证码错误、不存在或已过期")
	}
	u := dal.Executor.User
	exist, err := u.WithContext(service.ctx).Where(u.Username.Eq(req.Username)).Count()
	if err != nil {
		return errno.DatabaseCallError
	}
	if exist != 0 {
		return errno.CustomError.WithMessage("用户名已存在")
	}

	err = u.WithContext(service.ctx).Create(&model.User{
		ID:       snowflake.UserIDGenerator.Generate(),
		Username: req.Username,
		Password: encrypt.EncryptBySHA256WithSalt(req.Password, encrypt.GetSalt()),
		Email:    req.Email,
	})

	if err != nil {
		return errno.DatabaseCallError
	}
	return nil
}

func (service *UserService) NewSecurityEmailCodeEvent(req *user.UserSecurityEmailCodeReq) error {
	codeFormat := generator.AlnumGeneratorOption{
		Length:    6,
		UseNumber: true,
	}
	code := generator.GenerateAlnumString(codeFormat)
	if err := redis.EmailCodeStore(req.Email, code); err != nil {
		return errno.DatabaseCallError
	}
	email := &mail.Email{
		To:      []string{req.Email},
		Subject: "noreply",
		HTML:    fmt.Sprintf(mail.HTML, "FuliFuli", code, "FuliFuli", "FuliFuli"),
	}
	mail.Station.Send(email)
	return nil
}

func (service *UserService) NewLoginEvent(req *user.UserLoginReq) (*model.User, error) {
	u := dal.Executor.User
	user, err := u.WithContext(service.ctx).Where(u.Username.Eq(req.Username)).First()
	if err != nil {
		return nil, errno.DatabaseCallError
	}
	if user == nil {
		return nil, errno.CustomError.WithMessage("用户不存在")
	}
	if user.MfaEnable {
		passed := mfa.NewAuthController(fmt.Sprint(user.ID), *req.MfaCode, user.MfaSecret).VerifyTOTP()
		if !passed {
			return nil, errno.MfaAuthFailed
		}
	}
	return user, nil
}
