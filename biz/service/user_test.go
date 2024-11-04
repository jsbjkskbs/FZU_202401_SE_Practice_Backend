package service

import (
	"testing"
	"time"

	"sfw/biz/mw/jwt"

	"sfw/pkg/oss"

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
		name                   string
		req                    *user.UserRegisterReq
		errorIsExist           bool
		expectedError          string
		mockExistReturn        bool
		mockExistErrorReturn   error
		mockInsertErrorReturn  error
		mockCodeGetErrorReturn error
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
			name: "CodeGetError",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "jkskj123456",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockCodeGetErrorReturn: errno.DatabaseCallError,
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
			name: "QueryError",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "jkskj123456",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockExistErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "InsertError",
			req: &user.UserRegisterReq{
				Username: "jkskj",
				Password: "jkskj123456",
				Email:    "test@example.com",
				Code:     "123456",
			},
			errorIsExist:          true,
			expectedError:         errno.DatabaseCallErrorMsg,
			mockInsertErrorReturn: errno.DatabaseCallError,
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

			mockey.Mock(redis.EmailCodeGet).Return("123456", tc.mockCodeGetErrorReturn).Build()
			mockey.Mock(exquery.QueryUserExistByUsernameOrEmail).Return(tc.mockExistReturn, tc.mockExistErrorReturn).Build()

			mockey.Mock(generator.UserIDGenerator.Generate).Return(111).Build()
			mockey.Mock(exquery.InsertUser).Return(tc.mockInsertErrorReturn).Build()
			mockey.Mock(redis.EmailCodeDel).Return(nil).Build()

			err := service.NewRegisterEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				time.Sleep(5 * time.Second)
			}
		})
	}
}

func TestNewSecurityEmailCodeEvent(t *testing.T) {
	type testCase struct {
		name                     string
		req                      *user.UserSecurityEmailCodeReq
		errorIsExist             bool
		expectedError            string
		mockExistReturn          bool
		mockCodeStoreErrorReturn error
		mockQueryErrorReturn     error
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
			name: "QueryError",
			req: &user.UserSecurityEmailCodeReq{
				Email: "test@example.com",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "CodeStoreError",
			req: &user.UserSecurityEmailCodeReq{
				Email: "test@example.com",
			},
			errorIsExist:             true,
			expectedError:            errno.DatabaseCallErrorMsg,
			mockCodeStoreErrorReturn: errno.DatabaseCallError,
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

			mockey.Mock(exquery.QueryUserExistByEmail).Return(tc.mockExistReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock(redis.EmailCodeStore).Return(tc.mockCodeStoreErrorReturn).Build()
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
		mockErrorReturn  error
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
			name: "QueryError",
			req: &user.UserLoginReq{
				Username: "jkskj",
				Password: "jkskj123456",
			},
			errorIsExist:    true,
			expectedError:   errno.DatabaseCallErrorMsg,
			mockErrorReturn: errno.DatabaseCallError,
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

			mockey.Mock(exquery.QueryUserByUsername).Return(tc.mockUserReturn, tc.mockErrorReturn).Build()
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

func TestNewInfoEvent(t *testing.T) {
	type testCase struct {
		name            string
		req             *user.UserInfoReq
		errorIsExist    bool
		expectedError   string
		mockUserReturn  *model.User
		mockErrorReturn error
	}

	testCases := []testCase{
		{
			name:          "ParamInvalid",
			req:           &user.UserInfoReq{UserID: ""},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name:          "IsNotExist",
			req:           &user.UserInfoReq{UserID: "111"},
			errorIsExist:  true,
			expectedError: "用户不存在",
		},
		{
			name:            "QueryError",
			req:             &user.UserInfoReq{UserID: "111"},
			errorIsExist:    true,
			expectedError:   errno.DatabaseCallErrorMsg,
			mockErrorReturn: errno.DatabaseCallError,
		},
		{
			name:         "Success",
			req:          &user.UserInfoReq{UserID: "111"},
			errorIsExist: false,
			mockUserReturn: &model.User{
				ID: 111,
			},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryUserByID).Return(tc.mockUserReturn, tc.mockErrorReturn).Build()
			mockey.Mock(model_converter.UserDal2Resp).Return(&base.User{}).Build()

			_, err := service.NewInfoEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewFollowerCountEvent(t *testing.T) {
	type testCase struct {
		name            string
		req             *user.UserFollowerCountReq
		errorIsExist    bool
		expectedError   string
		mockCountReturn int64
		mockErrorReturn error
		expectedResult  int64
	}

	testCases := []testCase{
		{
			name:          "ParamInvalid",
			req:           &user.UserFollowerCountReq{UserID: ""},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name:            "QueryError",
			req:             &user.UserFollowerCountReq{UserID: "111"},
			errorIsExist:    true,
			expectedError:   errno.DatabaseCallErrorMsg,
			mockErrorReturn: errno.DatabaseCallError,
		},
		{
			name:            "Success",
			req:             &user.UserFollowerCountReq{UserID: "111"},
			errorIsExist:    false,
			expectedResult:  1,
			mockCountReturn: 1,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryFollowerCountByUserID).Return(tc.mockCountReturn, tc.mockErrorReturn).Build()

			result, err := service.NewFollowerCountEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}

func TestNewFollowingCountEvent(t *testing.T) {
	type testCase struct {
		name            string
		req             *user.UserFollowingCountReq
		errorIsExist    bool
		expectedError   string
		mockCountReturn int64
		mockErrorReturn error
		expectedResult  int64
	}

	testCases := []testCase{
		{
			name:          "ParamInvalid",
			req:           &user.UserFollowingCountReq{UserID: ""},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name:            "QueryError",
			req:             &user.UserFollowingCountReq{UserID: "111"},
			errorIsExist:    true,
			expectedError:   errno.DatabaseCallErrorMsg,
			mockErrorReturn: errno.DatabaseCallError,
		},
		{
			name:            "Success",
			req:             &user.UserFollowingCountReq{UserID: "111"},
			errorIsExist:    false,
			expectedResult:  1,
			mockCountReturn: 1,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryFollowingCountByUserID).Return(tc.mockCountReturn, tc.mockErrorReturn).Build()

			result, err := service.NewFollowingCountEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}

func TestNewLikeCountEvent(t *testing.T) {
	type testCase struct {
		name            string
		req             *user.UserLikeCountReq
		errorIsExist    bool
		expectedError   string
		mockCountReturn int64
		mockErrorReturn error
		expectedResult  int64
	}

	testCases := []testCase{
		{
			name:          "ParamInvalid",
			req:           &user.UserLikeCountReq{UserID: ""},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name:            "QueryError",
			req:             &user.UserLikeCountReq{UserID: "111"},
			errorIsExist:    true,
			expectedError:   errno.DatabaseCallErrorMsg,
			mockErrorReturn: errno.DatabaseCallError,
		},
		{
			name:            "Success",
			req:             &user.UserLikeCountReq{UserID: "111"},
			errorIsExist:    false,
			expectedResult:  1,
			mockCountReturn: 1,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryUserLikeCount).Return(tc.mockCountReturn, tc.mockErrorReturn).Build()

			result, err := service.NewLikeCountEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}

func TestNewAvatarUploadEvent(t *testing.T) {
	type testCase struct {
		name                      string
		req                       *user.UserAvatarUploadReq
		errorIsExist              bool
		expectedError             string
		mockTokenErrorReturn      error
		mockUploadUploadKeyReturn string
		mockUploadUptokenReturn   string
		mockUploadErrorReturn     error
		expectedResult            *user.UserAvatarUploadData
	}

	testCases := []testCase{
		{
			name:                 "AccessTokenInvalid",
			req:                  &user.UserAvatarUploadReq{AccessToken: ""},
			errorIsExist:         true,
			expectedError:        errno.AccessTokenInvalidErrorMsg,
			mockTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name:                  "UploadError",
			req:                   &user.UserAvatarUploadReq{AccessToken: "111"},
			errorIsExist:          true,
			expectedError:         errno.InternalServerErrorMsg,
			mockUploadErrorReturn: errno.InternalServerError,
		},
		{
			name:                      "Success",
			req:                       &user.UserAvatarUploadReq{AccessToken: "111"},
			errorIsExist:              false,
			mockUploadUploadKeyReturn: "111",
			mockUploadUptokenReturn:   "111",
			expectedResult: &user.UserAvatarUploadData{
				UploadURL: oss.UploadUrl,
				UploadKey: "111",
				Uptoken:   "111",
			},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockTokenErrorReturn).Build()
			mockey.Mock(oss.UploadAvatar).Return(tc.mockUploadUptokenReturn, tc.mockUploadUploadKeyReturn, tc.mockUploadErrorReturn).Build()

			result, err := service.NewAvatarUploadEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}

func TestNewMfaQrcodeEvent(t *testing.T) {
	type testCase struct {
		name                    string
		req                     *user.UserMfaQrcodeReq
		errorIsExist            bool
		expectedError           string
		mockTokenErrorReturn    error
		mockQRCodeReturn        string
		mockInfoReturn          *mfa.MfaAuthInfo
		mockGenerateErrorReturn error
		expectedResult          *user.UserMfaQrcodeData
	}

	testCases := []testCase{
		{
			name:                 "AccessTokenInvalid",
			req:                  &user.UserMfaQrcodeReq{AccessToken: ""},
			errorIsExist:         true,
			expectedError:        errno.AccessTokenInvalidErrorMsg,
			mockTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name:                    "GenerateError",
			req:                     &user.UserMfaQrcodeReq{AccessToken: "111"},
			errorIsExist:            true,
			expectedError:           errno.MfaGenerateFailedErrorMsg,
			mockGenerateErrorReturn: errno.MfaGenerateFailed,
		},
		{
			name:         "Success",
			req:          &user.UserMfaQrcodeReq{AccessToken: "111"},
			errorIsExist: false,
			mockInfoReturn: &mfa.MfaAuthInfo{
				Url:    "111",
				Secret: "111",
			},
			mockQRCodeReturn: "111",
			expectedResult: &user.UserMfaQrcodeData{
				Qrcode: "111",
				Secret: "111",
			},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockTokenErrorReturn).Build()
			mockey.Mock((*mfa.AuthController).GenerateTOTP).Return(tc.mockInfoReturn, tc.mockGenerateErrorReturn).Build()
			mockey.Mock(encrypt.EncodeUrlToQrcodeAsPng).Return(tc.mockQRCodeReturn).Build()

			result, err := service.NewMfaQrcodeEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
