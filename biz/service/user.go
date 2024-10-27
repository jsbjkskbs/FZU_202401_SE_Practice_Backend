package service

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/user"
	"sfw/biz/model/base"
	"sfw/biz/mw/generator/snowflake"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/service/checker"
	"sfw/biz/service/service_converter"
	"sfw/pkg/errno"
	"sfw/pkg/oss"
	"sfw/pkg/utils/encrypt"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/mail"
	"sfw/pkg/utils/mfa"
	"strconv"
	"time"

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
	// check username and password
	err := checker.CheckUsername(req.Username)
	if err != nil {
		return errno.CustomError.WithMessage(err.Error())
	}
	err = checker.CheckPassword(req.Password)
	if err != nil {
		return errno.CustomError.WithMessage(err.Error())
	}

	// check email and code
	code, err := redis.EmailCodeGet(req.Email)
	if err != nil {
		return errno.DatabaseCallError
	}
	if code != req.Code {
		return errno.CustomError.WithMessage("验证码错误、不存在或已过期")
	}

	// check username and email exist
	u := dal.Executor.User
	exist, err := u.WithContext(service.ctx).Where(u.Username.Eq(req.Username)).Count()
	if err != nil {
		return errno.DatabaseCallError
	}
	if exist != 0 {
		return errno.CustomError.WithMessage("用户名已存在")
	}

	// create user
	err = u.WithContext(service.ctx).Create(&model.User{
		ID:       snowflake.UserIDGenerator.Generate(),
		Username: req.Username,
		Password: encrypt.EncryptBySHA256WithSalt(req.Password, encrypt.GetSalt()),
		Email:    req.Email,
	})
	if err != nil {
		return errno.DatabaseCallError
	}

	// not need to check error, because it's not a critical operation
	go redis.EmailCodeDel(req.Email)
	return nil
}

func (service *UserService) NewSecurityEmailCodeEvent(req *user.UserSecurityEmailCodeReq) error {
	u := dal.Executor.User
	exist, err := u.WithContext(service.ctx).Where(u.Email.Eq(req.Email)).Count()
	if err != nil {
		return errno.DatabaseCallError
	}
	if exist != 0 {
		return errno.CustomError.WithMessage("邮箱已存在")
	}

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
		if req.MfaCode == nil {
			return nil, errno.MfaAuthFailed
		}
		passed := mfa.NewAuthController(fmt.Sprint(user.ID), *req.MfaCode, user.MfaSecret).VerifyTOTP()
		if !passed {
			return nil, errno.MfaAuthFailed
		}
	}
	user.AvatarURL = oss.Key2Url(user.AvatarURL)
	return user, nil
}

func (service *UserService) NewInfoEvent(req *user.UserInfoReq) (*model.User, error) {
	u := dal.Executor.User
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid
	}
	user, err := u.WithContext(service.ctx).Where(u.ID.Eq(uid)).First()
	if err != nil {
		return nil, errno.DatabaseCallError
	}
	if user == nil {
		return nil, errno.CustomError.WithMessage("用户不存在")
	}
	user.AvatarURL = oss.Key2Url(user.AvatarURL)
	return user, nil
}

func (service *UserService) NewFollowerCountEvent(req *user.UserFollowerCountReq) (int64, error) {
	f := dal.Executor.Follow
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return 0, errno.ParamInvalid
	}
	count, err := f.WithContext(service.ctx).Where(f.FollowedID.Eq(uid)).Count()
	if err != nil {
		return 0, errno.DatabaseCallError
	}
	return count, nil
}

func (service *UserService) NewFollowingCountEvent(req *user.UserFollowingCountReq) (int64, error) {
	f := dal.Executor.Follow
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return 0, errno.ParamInvalid
	}
	count, err := f.WithContext(service.ctx).Where(f.FollowerID.Eq(uid)).Count()
	if err != nil {
		return 0, errno.DatabaseCallError
	}
	return count, nil
}

func (service *UserService) NewLikeCountEvent(req *user.UserLikeCountReq) (int64, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return 0, errno.ParamInvalid
	}

	sum, err := exquery.QueryUserLikeCount(uid)
	if err != nil {
		return 0, errno.DatabaseCallError
	}

	return sum, nil
}

func (service *UserService) NewAvatarUploadEvent(req *user.UserAvatarUploadReq) (string, error) {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return "", errno.AccessTokenInvalid
	}
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return "", errno.InternalServerError
	}

	uptoken, err := oss.UploadAvatar(uid, id)
	if err != nil {
		return "", errno.InternalServerError
	}
	return uptoken, nil
}

func (service *UserService) NewMfaQrcodeEvent(req *user.UserMfaQrcodeReq) (*user.UserMfaQrcodeData, error) {
	id, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	info, err := mfa.NewAuthController(id, "", "").GenerateTOTP()
	if err != nil {
		return nil, errno.MfaGenerateFailed
	}

	qrcode := encrypt.EncodeUrlToQrcodeAsPng(info.Url)
	return &user.UserMfaQrcodeData{
		Qrcode: qrcode,
		Secret: info.Secret,
	}, nil
}

func (service *UserService) NewMfaBindEvent(req *user.UserMfaBindReq) error {
	id, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.InternalServerError
	}

	passed := mfa.NewAuthController(id, req.Code, req.Secret).VerifyTOTP()
	if !passed {
		return errno.MfaAuthFailed
	}

	u := dal.Executor.User
	_, err = u.WithContext(service.ctx).Where(u.ID.Eq(uid)).Updates(model.User{
		MfaEnable: true,
		MfaSecret: req.Secret,
	})
	if err != nil {
		return errno.DatabaseCallError
	}
	return nil
}

func (service *UserService) NewSearchEvent(req *user.UserSearchReq) (*[]*base.User, bool, int64, int64, error) {
	req.PageNum = max(req.PageNum, 0)
	req.PageSize = max(req.PageSize, 1)

	u := dal.Executor.User
	users, count, err := u.
		WithContext(service.ctx).
		Where(u.Username.Like(fmt.Sprintf("%%%s%%", req.Keyword))).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))
	if err != nil {
		return nil, true, req.PageNum, req.PageSize, errno.DatabaseCallError
	}
	for _, user := range users {
		user.AvatarURL = oss.Key2Url(user.AvatarURL)
	}
	return service_converter.UserListDal2Resp(&users), count <= (req.PageNum+1)*req.PageSize, req.PageNum, req.PageSize, nil
}

func (service *UserService) NewSecurityPasswordRetrieve(req *user.UserPasswordRetrieveReq) error {
	var (
		user *model.User
		err  error
	)

	u := dal.Executor.User
	code := generator.GenerateAlnumString(generator.AlnumGeneratorOption{
		Length:    6,
		UseNumber: true,
	})

	switch req.Otype {
	case "email":
		user, err = u.WithContext(service.ctx).Where(u.Email.Eq(req.Oid)).First()
		if err != nil {
			return errno.DatabaseCallError
		}
		if user == nil {
			return errno.CustomError.WithMessage("邮箱不存在")
		}
		mail.Station.Send(&mail.Email{
			To:      []string{req.Oid},
			Subject: "noreply",
			HTML:    fmt.Sprintf(mail.HTML, "FuliFuli", code, "FuliFuli", "FuliFuli"),
		})
		if err := redis.EmailCodeStore(req.Oid, code); err != nil {
			return errno.DatabaseCallError
		}
	default:
		return errno.CustomError.WithMessage("暂不支持该类型: " + req.Otype)
	}
	err = redis.TokenExpireTimeStore(fmt.Sprint(user.ID), time.Now().Unix(), jwt.AccessTokenExpireTime+1*time.Minute)
	if err != nil {
		return errno.DatabaseCallError
	}

	return nil
}

func (servcie *UserService) NewSecurityPasswordResetEvent(req *user.UserPasswordResetReq) error {
	if err := checker.CheckPassword(req.Password); err != nil {
		return err
	}

	switch req.Otype {
	case "email":
		code, err := redis.EmailCodeGet(req.Oid)
		if err != nil {
			return errno.DatabaseCallError
		}
		if code != req.Code {
			return errno.CustomError.WithMessage("验证码错误、不存在或已过期")
		}
		u := dal.Executor.User
		_, err = u.WithContext(servcie.ctx).
			Where(u.Email.Eq(req.Oid)).
			Update(u.Password, encrypt.EncryptBySHA256WithSalt(req.Password, encrypt.GetSalt()))
		if err != nil {
			return errno.DatabaseCallError
		}
		go redis.EmailCodeDel(req.Oid)
	default:
		return errno.CustomError.WithMessage("暂不支持该类型: " + req.Otype)
	}
	return nil
}
