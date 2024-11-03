package service

import (
	"testing"
	"time"

	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/user"
	"sfw/biz/model/base"
	"sfw/biz/mw/redis"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/utils/encrypt"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/mail"
	"sfw/pkg/utils/mfa"
)

var service = new(UserService)

func TestNewRegisterEvent(t *testing.T) {
	type testCase struct {
		name            string
		req             *user.UserRegisterReq
		errorIsExist    bool
		expectedError   string
		mockExistReturn bool
	}

	testCases := []testCase{
		{
			name: "UsernameIsEmpty",
			req: &user.UserRegisterReq{
				Username: "",
				Password: "123",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:  true,
			expectedError: "用户名不符合规范",
		},
		{
			name: "UsernameContainsWhiteSpace",
			req: &user.UserRegisterReq{
				Username: "jk skj",
				Password: "123",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:  true,
			expectedError: "用户名不符合规范",
		},
		{
			name: "PasswordIsEmpty",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:  true,
			expectedError: "密码不符合规范",
		},
		{
			name: "PasswordContainsWhiteSpace",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "12 3",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:  true,
			expectedError: "密码不符合规范",
		},
		{
			name: "CodeError",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "jkskj123456",
				Email:    "test@example.com",
				Code:     "12345678",
			},
			errorIsExist:  true,
			expectedError: "验证码错误、不存在或已过期",
		},
		{
			name: "IsExist",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "jkskj123456",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:    true,
			expectedError:   "用户名或邮箱已存在",
			mockExistReturn: true,
		},
		{
			name: "Success",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "jkskj123456",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:    false,
			mockExistReturn: false,
		},
	}

	defer mockey.UnPatchAll()

	generator.UserIDGenerator, _ = generator.NewSnowflake(1)

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(redis.EmailCodeGet).Return("123456", nil).Build()
			mockey.Mock(exquery.QueryUserExistByUsernameOrEmail).Return(tc.mockExistReturn, nil).Build()

			mockey.Mock(generator.UserIDGenerator.Generate).Return(111).Build()
			mockey.Mock(exquery.InsertUser).Return(nil).Build()
			mockey.Mock(redis.EmailCodeDel).Return(nil).Build()

			err := service.NewRegisterEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				time.Sleep(20 * time.Second)
			}
		})
	}
}

func TestNewSecurityEmailCodeEvent(t *testing.T) {
	type testCase struct {
		name            string
		req             *user.UserSecurityEmailCodeReq
		errorIsExist    bool
		expectedError   string
		mockExistReturn bool
	}

	testCases := []testCase{
		{
			name: "EmailIsEmpty",
			req: &user.UserSecurityEmailCodeReq{
				Email: "test@example.com",
			},
			errorIsExist:    true,
			expectedError:   "邮箱已存在",
			mockExistReturn: true,
		},
		{
			name: "Success",
			req: &user.UserSecurityEmailCodeReq{
				Email: "test@example.com",
			},
			errorIsExist: false,
		},
	}

	defer mockey.UnPatchAll()
	mail.Config = new(mail.EmailStationConfig)
	mail.Station = mail.NewEmailStation(*mail.Config)

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryUserExistByEmail).Return(tc.mockExistReturn, nil).Build()
			mockey.Mock(redis.EmailCodeStore).Return(nil).Build()
			mockey.Mock(mockey.GetMethod(mail.Station, "Send")).Return().Build()

			err := service.NewSecurityEmailCodeEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewLoginEvent(t *testing.T) {
	type testCase struct {
		name             string
		req              *user.UserLoginReq
		errorIsExist     bool
		expectedError    string
		mockUserReturn   *model.User
		mockVerifyReturn bool
	}

	testCases := []testCase{
		{
			name: "UserIsNotExist",
			req: &user.UserLoginReq{
				Username: "jkskj",
				Password: "jkskj123456",
			},
			errorIsExist:   true,
			expectedError:  "用户不存在",
			mockUserReturn: nil,
		},
		{
			name: "PasswordIsWrong",
			req: &user.UserLoginReq{
				Username: "jkskj",
				Password: "jkskj123456",
			},
			errorIsExist:  true,
			expectedError: "密码错误",
			mockUserReturn: &model.User{
				Password:  "1111",
				MfaEnable: false,
			},
		},
		{
			name: "MFAMissCode",
			req: &user.UserLoginReq{
				Username: "jkskj",
				Password: "jkskj123456",
			},
			errorIsExist:  true,
			expectedError: errno.MfaAuthFailedErrorMsg,
			mockUserReturn: &model.User{
				Password:  encrypt.EncryptBySHA256WithSalt("jkskj123456", encrypt.GetSalt()),
				MfaEnable: true,
			},
		},
		{
			name: "MFAFail",
			req: &user.UserLoginReq{
				Username: "jkskj",
				Password: "jkskj123456",
				MfaCode:  new(string),
			},
			errorIsExist:  true,
			expectedError: errno.MfaAuthFailedErrorMsg,
			mockUserReturn: &model.User{
				Password:  encrypt.EncryptBySHA256WithSalt("jkskj123456", encrypt.GetSalt()),
				MfaEnable: true,
			},
			mockVerifyReturn: false,
		},
		{
			name: "Success",
			req: &user.UserLoginReq{
				Username: "jkskj",
				Password: "jkskj123456",
				MfaCode:  new(string),
			},
			errorIsExist: false,
			mockUserReturn: &model.User{
				Password:  encrypt.EncryptBySHA256WithSalt("jkskj123456", encrypt.GetSalt()),
				MfaEnable: true,
			},
			mockVerifyReturn: true,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryUserByUsername).Return(tc.mockUserReturn, nil).Build()
			mockey.Mock((*mfa.AuthController).VerifyTOTP).Return(tc.mockVerifyReturn).Build()
			mockey.Mock(model_converter.UserWithTokenDal2Resp).Return(&base.UserWithToken{}).Build()

			_, err := service.NewLoginEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
