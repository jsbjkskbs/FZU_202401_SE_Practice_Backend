package service

import (
	"testing"

	"github.com/bytedance/mockey"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stretchr/testify/assert"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/activity"
	"sfw/biz/model/base"
	"sfw/biz/mw/jwt"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/utils/generator"
)

var activityService = NewActivityService(nil, new(app.RequestContext))

func TestNewActivityFeedEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *activity.ActivityFeedReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *activity.ActivityFeedRespData
		mockQueryErrorReturn       error
		mockQueryActivityReturn    []*model.Activity
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    *[]*base.Activity
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &activity.ActivityFeedReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "QueryFail",
			req: &activity.ActivityFeedReq{
				AccessToken: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ConvertFail",
			req: &activity.ActivityFeedReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &activity.ActivityFeedReq{
				AccessToken: "111",
			},
			errorIsExist: false,
			expectedResult: &activity.ActivityFeedRespData{
				Items:    make([]*base.Activity, 0),
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: &[]*base.Activity{},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryActivityByFollowedIdPaged).Return(tc.mockQueryActivityReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.ActivityListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := activityService.NewFeedEvent(tc.req)

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

func TestNewActivityPublishEvent(t *testing.T) {
	type testCase struct {
		name                          string
		req                           *activity.ActivityPublishReq
		errorIsExist                  bool
		expectedError                 string
		mockQueryImageErrorReturn     error
		mockQueryImageExistReturn     bool
		mockInsertActivityErrorReturn error
		mockInsertImageErrorReturn    error
		mockAccessTokenErrorReturn    error
	}

	refActivity := "1"
	refVideo := "1"
	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "RefActivityParamInvalid",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
				RefActivity: new(string),
			},
			errorIsExist:  true,
			expectedError: "无效的动态ID",
		},
		{
			name: "RefVideoParamInvalid",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
				RefVideo:    new(string),
			},
			errorIsExist:  true,
			expectedError: "无效的视频ID",
		},
		{
			name: "ParamInvalid",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
				RefVideo:    &refActivity,
				RefActivity: &refVideo,
			},
			errorIsExist:  true,
			expectedError: "只能引用一个视频或一个动态",
		},
		{
			name: "ImageParamInvalid",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
				Image:       []string{"a"},
			},
			errorIsExist:  true,
			expectedError: "无效的图片ID: a",
		},
		{
			name: "QueryImageFail",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
				Image:       []string{"1"},
			},
			errorIsExist:              true,
			expectedError:             errno.DatabaseCallErrorMsg,
			mockQueryImageErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "QueryImageIsNotExist",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
				Image:       []string{"1"},
			},
			errorIsExist:  true,
			expectedError: "图片不存在: 1",
		},
		{
			name: "InsertImageFail",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
				Image:       []string{"1"},
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryImageExistReturn:  true,
			mockInsertImageErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "InsertActivityFail",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
			},
			errorIsExist:                  true,
			expectedError:                 errno.DatabaseCallErrorMsg,
			mockInsertActivityErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &activity.ActivityPublishReq{
				AccessToken: "111",
			},
			errorIsExist: false,
		},
	}

	generator.ActivityIDGenerator, _ = generator.NewSnowflake(3)
	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(generator.ActivityIDGenerator.Generate).Return(111).Build()
			mockey.Mock(exquery.QueryImageExistById).Return(tc.mockQueryImageExistReturn, tc.mockQueryImageErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(exquery.InsertActivity).Return(tc.mockInsertActivityErrorReturn).Build()
			mockey.Mock(exquery.InsertActivityImage).Return(tc.mockInsertImageErrorReturn).Build()

			err := activityService.NewPublishEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewActivityListEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *activity.ActivityListReq
		errorIsExist               bool
		expectedError              string
		mockQueryErrorReturn       error
		mockQueryActivityReturn    []*model.Activity
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    *[]*base.Activity
		expectedResult             *activity.ActivityListRespData
	}

	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: &activity.ActivityListReq{
				UserID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "无效的用户ID",
		},
		{
			name: "QueryFail",
			req: &activity.ActivityListReq{
				UserID: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "AccessTokenFail",
			req: &activity.ActivityListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ConvertFail",
			req: &activity.ActivityListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &activity.ActivityListReq{
				UserID: "111",
			},
			errorIsExist: false,
			expectedResult: &activity.ActivityListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: new([]*base.Activity),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryActivityByUserIdPaged).Return(tc.mockQueryActivityReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.ActivityListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := activityService.NewListEvent(tc.req)

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
