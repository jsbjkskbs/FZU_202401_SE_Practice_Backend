package service

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/user"
	"sfw/biz/model/base"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/oss"
	"sfw/pkg/utils/checker"
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
		return errno.CustomError.WithMessage("用户名不符合规范")
	}
	err = checker.CheckPassword(req.Password)
	if err != nil {
		return errno.CustomError.WithMessage("密码不符合规范")
	}

	// check email and code
	code, err := redis.EmailCodeGet(req.Email)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if code != req.Code {
		return errno.CustomError.WithMessage("验证码错误、不存在或已过期")
	}

	// check username and email exist
	exist, err := exquery.QueryUserExistByUsernameOrEmail(req.Username, req.Email)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist {
		return errno.CustomError.WithMessage("用户名或邮箱已存在")
	}

	// create user
	err = exquery.InsertUser(&model.User{
		ID:       generator.UserIDGenerator.Generate(),
		Username: req.Username,
		Password: encrypt.EncryptBySHA256WithSalt(req.Password, encrypt.GetSalt()),
		Email:    req.Email,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}

	// not need to check error, because it's not a critical operation
	go redis.EmailCodeDel(req.Email)
	return nil
}

func (service *UserService) NewSecurityEmailCodeEvent(req *user.UserSecurityEmailCodeReq) error {
	exist, err := exquery.QueryUserExistByEmail(req.Email)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist {
		return errno.CustomError.WithMessage("邮箱已存在")
	}

	code := generator.GenerateAlnumString(
		generator.AlnumGeneratorOption{
			Length:    6,
			UseNumber: true,
		},
	)
	if err := redis.EmailCodeStore(req.Email, code); err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	mail.Station.Send(
		&mail.Email{
			To:      []string{req.Email},
			Subject: "noreply",
			HTML:    fmt.Sprintf(mail.HTML, "FuliFuli", code, "FuliFuli", "FuliFuli"),
		},
	)
	return nil
}

func (service *UserService) NewLoginEvent(req *user.UserLoginReq) (*base.UserWithToken, error) {
	user, err := exquery.QueryUserByUsername(req.Username)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if user == nil {
		return nil, errno.CustomError.WithMessage("用户不存在")
	}
	if user.Password != encrypt.EncryptBySHA256WithSalt(req.Password, encrypt.GetSalt()) {
		return nil, errno.CustomError.WithMessage("密码错误")
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
	return model_converter.UserWithTokenDal2Resp(user), nil
}

func (service *UserService) NewInfoEvent(req *user.UserInfoReq) (*base.User, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("用户ID错误")
	}
	user, err := exquery.QueryUserByID(uid)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if user == nil {
		return nil, errno.CustomError.WithMessage("用户不存在")
	}
	return model_converter.UserDal2Resp(user), nil
}

func (service *UserService) NewFollowerCountEvent(req *user.UserFollowerCountReq) (int64, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return 0, errno.ParamInvalid.WithMessage("用户ID错误")
	}
	count, err := exquery.QueryFollowerCountByUserID(uid)
	if err != nil {
		return 0, errno.DatabaseCallError.WithInnerError(err)
	}
	return count, nil
}

func (service *UserService) NewFollowingCountEvent(req *user.UserFollowingCountReq) (int64, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return 0, errno.ParamInvalid.WithMessage("用户ID错误")
	}
	count, err := exquery.QueryFollowingCountByUserID(uid)
	if err != nil {
		return 0, errno.DatabaseCallError.WithInnerError(err)
	}
	return count, nil
}

func (service *UserService) NewLikeCountEvent(req *user.UserLikeCountReq) (int64, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return 0, errno.ParamInvalid.WithInnerError(err)
	}

	sum, err := exquery.QueryUserLikeCount(uid)
	if err != nil {
		return 0, errno.DatabaseCallError.WithInnerError(err)
	}

	return sum, nil
}

func (service *UserService) NewAvatarUploadEvent(req *user.UserAvatarUploadReq) (*user.UserAvatarUploadData, error) {
	id, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	uptoken, uploadKey, err := oss.UploadAvatar(fmt.Sprint(id), id)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}
	return &user.UserAvatarUploadData{
		UploadURL: oss.UploadUrl,
		UploadKey: uploadKey,
		Uptoken:   uptoken,
	}, nil
}

func (service *UserService) NewMfaQrcodeEvent(req *user.UserMfaQrcodeReq) (*user.UserMfaQrcodeData, error) {
	id, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	info, err := mfa.NewAuthController(fmt.Sprint(id), "", "").GenerateTOTP()
	if err != nil {
		return nil, errno.MfaGenerateFailed.WithInnerError(err)
	}

	qrcode := encrypt.EncodeUrlToQrcodeAsPng(info.Url)
	return &user.UserMfaQrcodeData{
		Qrcode: qrcode,
		Secret: info.Secret,
	}, nil
}

func (service *UserService) NewMfaBindEvent(req *user.UserMfaBindReq) error {
	id, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	passed := mfa.NewAuthController(fmt.Sprint(id), req.Code, req.Secret).VerifyTOTP()
	if !passed {
		return errno.MfaAuthFailed
	}

	u := dal.Executor.User
	_, err = u.WithContext(service.ctx).Where(u.ID.Eq(id)).Updates(model.User{
		MfaEnable: true,
		MfaSecret: req.Secret,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *UserService) NewSearchEvent(req *user.UserSearchReq) (*user.UserSearchRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	users, count, err := exquery.QueryUserFuzzyByUsernamePaged("%"+req.Keyword+"%", req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &user.UserSearchRespData{
		Items:    *model_converter.UserListDal2Resp(&users),
		IsEnd:    count < req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *UserService) NewSecurityPasswordRetrieveEmail(req *user.UserPasswordRetrieveEmailReq) error {
	var (
		user *model.User
		err  error
	)

	code := generator.GenerateAlnumString(generator.AlnumGeneratorOption{
		Length:    6,
		UseNumber: true,
	})

	user, err = exquery.QueryUserByEmail(req.Email)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if user == nil {
		return errno.CustomError.WithMessage("用户不存在")
	}
	mail.Station.Send(&mail.Email{
		To:      []string{req.Email},
		Subject: "noreply",
		HTML:    fmt.Sprintf(mail.HTML, "FuliFuli", code, "FuliFuli", "FuliFuli"),
	})
	if err := redis.EmailCodeStore(req.Email, code); err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	err = redis.TokenExpireTimeStore(fmt.Sprint(user.ID), time.Now().Unix(), jwt.RefreshTokenExpireTime-1*time.Minute)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}

	return nil
}

func (servcie *UserService) NewSecurityPasswordResetEmailEvent(req *user.UserPasswordResetEmailReq) error {
	if err := checker.CheckPassword(req.Password); err != nil {
		return errno.CustomError.WithMessage("密码不符合规范")
	}

	code, err := redis.EmailCodeGet(req.Email)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if code != req.Code {
		return errno.CustomError.WithMessage("验证码错误、不存在或已过期")
	}
	err = exquery.UpdateUserWithEmail(&model.User{
		Email:    req.Email,
		Password: encrypt.EncryptBySHA256WithSalt(req.Password, encrypt.GetSalt()),
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	go redis.EmailCodeDel(req.Email)
	return nil
}
