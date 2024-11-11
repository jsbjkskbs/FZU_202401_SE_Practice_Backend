package service

import (
	"testing"

	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/utils/checker"

	"github.com/bytedance/mockey"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stretchr/testify/assert"
	"sfw/biz/dal/exquery"
	"sfw/biz/model/api/report"
	"sfw/biz/mw/jwt"
	"sfw/pkg/errno"
	"sfw/pkg/utils/generator"
)

var reportService = NewReportService(nil, new(app.RequestContext))

func TestNewReportVideoEvent(t *testing.T) {
	type testCase struct {
		name                        string
		req                         *report.ReportVideoReq
		errorIsExist                bool
		expectedError               string
		mockQueryVideoErrorReturn   error
		mockQueryVideoResultReturn  bool
		mockInsertReportErrorReturn error
		mockQueryCountErrorReturn   error
		mockQueryCountResultReturn  int64
		mockAccessTokenErrorReturn  error
		mockAccessTokenResultReturn int64
	}

	testCases := []testCase{
		{
			name: "VideoIDParamInvalid",
			req: &report.ReportVideoReq{
				VideoID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "视频ID不合法",
		},
		{
			name: "QueryVideoFail",
			req: &report.ReportVideoReq{
				VideoID: "111",
			},
			errorIsExist:              true,
			expectedError:             errno.DatabaseCallErrorMsg,
			mockQueryVideoErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoIsNotExist",
			req: &report.ReportVideoReq{
				VideoID: "111",
			},
			errorIsExist:  true,
			expectedError: "视频不存在",
		},
		{
			name: "QueryCountFail",
			req: &report.ReportVideoReq{
				VideoID: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryCountErrorReturn:  errno.DatabaseCallError,
			mockQueryVideoResultReturn: true,
		},
		{
			name: "CountExceedTheLimit",
			req: &report.ReportVideoReq{
				VideoID: "111",
			},
			errorIsExist:               true,
			expectedError:              "您已经举报过该视频多次，请耐心等待处理结果",
			mockQueryCountResultReturn: 4,
			mockQueryVideoResultReturn: true,
		},
		{
			name: "AccessTokenFail",
			req: &report.ReportVideoReq{
				VideoID: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "InsertReportFail",
			req: &report.ReportVideoReq{
				VideoID: "111",
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockInsertReportErrorReturn: errno.DatabaseCallError,
			mockQueryVideoResultReturn:  true,
		},
		{
			name: "Success",
			req: &report.ReportVideoReq{
				VideoID: "111",
			},
			errorIsExist:               false,
			mockQueryVideoResultReturn: true,
		},
	}

	generator.VideoReportIDGenerator, _ = generator.NewSnowflake(7)

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoExistById).Return(tc.mockQueryVideoResultReturn, tc.mockQueryVideoErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(exquery.QueryVideoReportCountByUserIdAndVideoId).Return(tc.mockQueryCountResultReturn, tc.mockQueryCountErrorReturn).Build()
			mockey.Mock(exquery.InsertVideoReport).Return(tc.mockInsertReportErrorReturn).Build()
			mockey.Mock(generator.VideoReportIDGenerator.Generate).Return(111).Build()

			err := reportService.NewReportVideoEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewReportActivityEvent(t *testing.T) {
	type testCase struct {
		name                          string
		req                           *report.ReportActivityReq
		errorIsExist                  bool
		expectedError                 string
		mockQueryActivityErrorReturn  error
		mockQueryActivityResultReturn bool
		mockInsertReportErrorReturn   error
		mockQueryCountErrorReturn     error
		mockQueryCountResultReturn    int64
		mockAccessTokenErrorReturn    error
		mockAccessTokenResultReturn   int64
	}

	testCases := []testCase{
		{
			name: "ActivityIDParamInvalid",
			req: &report.ReportActivityReq{
				ActivityID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "动态ID不合法",
		},
		{
			name: "QueryActivityFail",
			req: &report.ReportActivityReq{
				ActivityID: "111",
			},
			errorIsExist:                 true,
			expectedError:                errno.DatabaseCallErrorMsg,
			mockQueryActivityErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ActivityIsNotExist",
			req: &report.ReportActivityReq{
				ActivityID: "111",
			},
			errorIsExist:  true,
			expectedError: "动态不存在",
		},
		{
			name: "QueryCountFail",
			req: &report.ReportActivityReq{
				ActivityID: "111",
			},
			errorIsExist:                  true,
			expectedError:                 errno.DatabaseCallErrorMsg,
			mockQueryCountErrorReturn:     errno.DatabaseCallError,
			mockQueryActivityResultReturn: true,
		},
		{
			name: "CountExceedTheLimit",
			req: &report.ReportActivityReq{
				ActivityID: "111",
			},
			errorIsExist:                  true,
			expectedError:                 "您已经举报过该动态多次，请耐心等待处理结果",
			mockQueryCountResultReturn:    4,
			mockQueryActivityResultReturn: true,
		},
		{
			name: "AccessTokenFail",
			req: &report.ReportActivityReq{
				ActivityID: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "InsertReportFail",
			req: &report.ReportActivityReq{
				ActivityID: "111",
			},
			errorIsExist:                  true,
			expectedError:                 errno.DatabaseCallErrorMsg,
			mockInsertReportErrorReturn:   errno.DatabaseCallError,
			mockQueryActivityResultReturn: true,
		},
		{
			name: "Success",
			req: &report.ReportActivityReq{
				ActivityID: "111",
			},
			errorIsExist:                  false,
			mockQueryActivityResultReturn: true,
		},
	}

	generator.ActivityReportIDGenerator, _ = generator.NewSnowflake(8)

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryActivityExistById).Return(tc.mockQueryActivityResultReturn, tc.mockQueryActivityErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(exquery.QueryActivityReportCountByUserIdAndActivityId).Return(tc.mockQueryCountResultReturn, tc.mockQueryCountErrorReturn).Build()
			mockey.Mock(exquery.InsertActivityReport).Return(tc.mockInsertReportErrorReturn).Build()
			mockey.Mock(generator.ActivityReportIDGenerator.Generate).Return(111).Build()

			err := reportService.NewReportActivityEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewReportANewReportCommentEvent(t *testing.T) {
	type testCase struct {
		name                         string
		req                          *report.ReportCommentReq
		errorIsExist                 bool
		expectedError                string
		mockQueryCommentErrorReturn  error
		mockQueryCommentResultReturn bool
		mockInsertReportErrorReturn  error
		mockQueryCountErrorReturn    error
		mockQueryCountResultReturn   int64
		mockAccessTokenErrorReturn   error
		mockAccessTokenResultReturn  int64
	}

	testCases := []testCase{
		{
			name: "ActivityIDParamInvalid",
			req: &report.ReportCommentReq{
				CommentType: "aaa",
			},
			errorIsExist:  true,
			expectedError: "评论类型错误",
		},
		{
			name: "VideoAccessTokenFail",
			req: &report.ReportCommentReq{
				CommentType: "video",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ActivityAccessTokenFail",
			req: &report.ReportCommentReq{
				CommentType: "activity",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "VideoFromMediaIDParamInvalid",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "视频ID不合法",
		},
		{
			name: "ActivityFromMediaIDParamInvalid",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "动态ID不合法",
		},
		{
			name: "VideoCommentIDParamInvalid",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "111",
				CommentID:   "aaa",
			},
			errorIsExist:  true,
			expectedError: "评论ID不合法",
		},
		{
			name: "ActivityCommentIDParamInvalid",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "111",
				CommentID:   "aaa",
			},
			errorIsExist:  true,
			expectedError: "评论ID不合法",
		},
		{
			name: "QueryVideoCommentFail",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockQueryCommentErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "QueryActivityCommentFail",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockQueryCommentErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoCommentIsNotExist",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:  true,
			expectedError: "评论不存在或视频与评论索引不匹配",
		},
		{
			name: "ActivityCommentIsNotExist",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:  true,
			expectedError: "评论不存在或动态与评论索引不匹配",
		},
		{
			name: "VideoQueryCountFail",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 true,
			expectedError:                errno.DatabaseCallErrorMsg,
			mockQueryCountErrorReturn:    errno.DatabaseCallError,
			mockQueryCommentResultReturn: true,
		},
		{
			name: "ActivityQueryCountFail",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 true,
			expectedError:                errno.DatabaseCallErrorMsg,
			mockQueryCountErrorReturn:    errno.DatabaseCallError,
			mockQueryCommentResultReturn: true,
		},
		{
			name: "VideoCountExceedTheLimit",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 true,
			expectedError:                "您已经举报过该评论多次，请耐心等待处理结果",
			mockQueryCountResultReturn:   4,
			mockQueryCommentResultReturn: true,
		},
		{
			name: "ActivityCountExceedTheLimit",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 true,
			expectedError:                "您已经举报过该评论多次，请耐心等待处理结果",
			mockQueryCountResultReturn:   4,
			mockQueryCommentResultReturn: true,
		},
		{
			name: "VideoInsertReportFail",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 true,
			expectedError:                errno.DatabaseCallErrorMsg,
			mockInsertReportErrorReturn:  errno.DatabaseCallError,
			mockQueryCommentResultReturn: true,
		},
		{
			name: "ActivityInsertReportFail",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 true,
			expectedError:                errno.DatabaseCallErrorMsg,
			mockInsertReportErrorReturn:  errno.DatabaseCallError,
			mockQueryCommentResultReturn: true,
		},
		{
			name: "VideoSuccess",
			req: &report.ReportCommentReq{
				CommentType: "video",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 false,
			mockQueryCommentResultReturn: true,
		},
		{
			name: "ActivitySuccess",
			req: &report.ReportCommentReq{
				CommentType: "activity",
				FromMediaID: "111",
				CommentID:   "111",
			},
			errorIsExist:                 false,
			mockQueryCommentResultReturn: true,
		},
	}

	generator.VideoCommentReportIDGenerator, _ = generator.NewSnowflake(9)
	generator.ActivityCommentReportIDGenerator, _ = generator.NewSnowflake(10)

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			switch tc.req.CommentType {
			case "video":
				mockey.Mock(exquery.QueryVideoCommentExistByIdAndVideoId).Return(tc.mockQueryCommentResultReturn, tc.mockQueryCommentErrorReturn).Build()
				mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
				mockey.Mock(exquery.QueryVideoCommentReportCountByUserIdAndCommentId).Return(tc.mockQueryCountResultReturn, tc.mockQueryCountErrorReturn).Build()
				mockey.Mock(exquery.InsertVideoCommentReport).Return(tc.mockInsertReportErrorReturn).Build()
				mockey.Mock(generator.VideoCommentReportIDGenerator.Generate).Return(111).Build()
			case "activity":
				mockey.Mock(exquery.QueryActivityCommentExistByIdAndActivityId).Return(tc.mockQueryCommentResultReturn, tc.mockQueryCommentErrorReturn).Build()
				mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
				mockey.Mock(exquery.QueryActivityCommentReportCountByUserIdAndCommentId).Return(tc.mockQueryCountResultReturn, tc.mockQueryCountErrorReturn).Build()
				mockey.Mock(exquery.InsertActivityCommentReport).Return(tc.mockInsertReportErrorReturn).Build()
				mockey.Mock(generator.ActivityCommentReportIDGenerator.Generate).Return(111).Build()
			}

			err := reportService.NewReportCommentEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewAdminVideoReportListEvent(t *testing.T) {
	type testCase struct {
		name                        string
		req                         *report.AdminVideoReportListReq
		errorIsExist                bool
		expectedError               string
		mockQueryReportErrorReturn  error
		mockQueryReportResultReturn []*model.VideoReport
		mockQueryReportCountReturn  int64
		mockConvertResultReturn     []*base.VideoReport
		expectedResult              *report.AdminVideoReportListRespData
	}

	testCases := []testCase{
		{
			name:                       "QueryReportFail",
			req:                        &report.AdminVideoReportListReq{},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name:                        "Success",
			req:                         &report.AdminVideoReportListReq{},
			errorIsExist:                false,
			mockQueryReportResultReturn: make([]*model.VideoReport, 0),
			expectedResult: &report.AdminVideoReportListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoReportByBasicInfoPaged).Return(tc.mockQueryReportResultReturn, tc.mockQueryReportCountReturn, tc.mockQueryReportErrorReturn).Build()
			mockey.Mock(model_converter.VideoReportDal2Resp).Return(tc.mockConvertResultReturn).Build()

			result, err := reportService.NewAdminVideoReportListEvent(tc.req)

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

func TestNewAdminActivityReportListEvent(t *testing.T) {
	type testCase struct {
		name                        string
		req                         *report.AdminActivityReportListReq
		errorIsExist                bool
		expectedError               string
		mockQueryReportErrorReturn  error
		mockQueryReportResultReturn []*model.ActivityReport
		mockQueryReportCountReturn  int64
		mockConvertResultReturn     []*base.ActivityReport
		expectedResult              *report.AdminActivityReportListRespData
	}

	testCases := []testCase{
		{
			name:                       "QueryReportFail",
			req:                        &report.AdminActivityReportListReq{},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name:                        "Success",
			req:                         &report.AdminActivityReportListReq{},
			errorIsExist:                false,
			mockQueryReportResultReturn: make([]*model.ActivityReport, 0),
			expectedResult: &report.AdminActivityReportListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryActivityReportByBasicInfoPaged).Return(tc.mockQueryReportResultReturn, tc.mockQueryReportCountReturn, tc.mockQueryReportErrorReturn).Build()
			mockey.Mock(model_converter.ActivityReportDal2Resp).Return(tc.mockConvertResultReturn).Build()

			result, err := reportService.NewAdminActivityReportListEvent(tc.req)

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

func TestNewAdminCommentReportListEvent(t *testing.T) {
	type testCase struct {
		name                                string
		req                                 *report.AdminCommentReportListReq
		errorIsExist                        bool
		expectedError                       string
		mockQueryReportErrorReturn          error
		mockQueryActivityReportResultReturn []*model.ActivityCommentReport
		mockQueryVideoReportResultReturn    []*model.VideoCommentReport
		mockQueryReportCountReturn          int64
		mockConvertResultReturn             []*base.CommentReport
		expectedResult                      *report.AdminCommentReportListRespData
	}

	testCases := []testCase{
		{
			name: "ActivityIDParamInvalid",
			req: &report.AdminCommentReportListReq{
				CommentType: "aaa",
			},
			errorIsExist:  true,
			expectedError: "评论类型错误",
		},
		{
			name: "QueryVideoReportFail",
			req: &report.AdminCommentReportListReq{
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "QueryActivityReportFail",
			req: &report.AdminCommentReportListReq{
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoSuccess",
			req: &report.AdminCommentReportListReq{
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:                     false,
			mockQueryVideoReportResultReturn: make([]*model.VideoCommentReport, 0),
			expectedResult: &report.AdminCommentReportListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
		},
		{
			name: "ActivitySuccess",
			req: &report.AdminCommentReportListReq{
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:                        false,
			mockQueryActivityReportResultReturn: make([]*model.ActivityCommentReport, 0),
			expectedResult: &report.AdminCommentReportListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoCommentReportByBasicInfoPaged).Return(tc.mockQueryVideoReportResultReturn, tc.mockQueryReportCountReturn, tc.mockQueryReportErrorReturn).Build()
			mockey.Mock(exquery.QueryActivityCommentReportByBasicInfoPaged).Return(tc.mockQueryActivityReportResultReturn, tc.mockQueryReportCountReturn, tc.mockQueryReportErrorReturn).Build()
			mockey.Mock(model_converter.ActivityCommentReportDal2Resp).Return(tc.mockConvertResultReturn).Build()
			mockey.Mock(model_converter.VideoCommentReportDal2Resp).Return(tc.mockConvertResultReturn).Build()

			result, err := reportService.NewAdminCommentReportListEvent(tc.req)

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

func TestNewAdminVideoReportHandleEvent(t *testing.T) {
	type testCase struct {
		name                        string
		req                         *report.AdminVideoReportHandleReq
		errorIsExist                bool
		expectedError               string
		mockQueryReportErrorReturn  error
		mockQueryReportResultReturn bool
		mockAccessTokenErrorReturn  error
		mockAccessTokenResultReturn int64
		mockUpdateReportErrorReturn error
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &report.AdminVideoReportHandleReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ReportIDParamInvalid",
			req: &report.AdminVideoReportHandleReq{
				AccessToken: "111",
				ReportID:    "aaa",
			},
			errorIsExist:  true,
			expectedError: "举报ID不合法",
		},
		{
			name: "QueryFail",
			req: &report.AdminVideoReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ReportIsNotExist",
			req: &report.AdminVideoReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
			},
			errorIsExist:  true,
			expectedError: "举报不存在或已经处理",
		},
		{
			name: "ActionTypeParamInvalid",
			req: &report.AdminVideoReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				ActionType:  666,
			},
			errorIsExist:                true,
			expectedError:               "操作类型不合法",
			mockQueryReportResultReturn: true,
		},
		{
			name: "UpdateFail",
			req: &report.AdminVideoReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				ActionType:  common.ActionTypeOff,
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockQueryReportResultReturn: true,
			mockUpdateReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &report.AdminVideoReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist:                false,
			mockQueryReportResultReturn: true,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(exquery.QueryVideoReportExistByIdAndStatus).Return(tc.mockQueryReportResultReturn, tc.mockQueryReportErrorReturn).Build()
			mockey.Mock(exquery.UpdateVideoReportById).Return(tc.mockUpdateReportErrorReturn).Build()

			err := reportService.NewAdminVideoReportHandleEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewAdminActivityReportHandleEvent(t *testing.T) {
	type testCase struct {
		name                        string
		req                         *report.AdminActivityReportHandleReq
		errorIsExist                bool
		expectedError               string
		mockQueryReportErrorReturn  error
		mockQueryReportResultReturn bool
		mockAccessTokenErrorReturn  error
		mockAccessTokenResultReturn int64
		mockUpdateReportErrorReturn error
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &report.AdminActivityReportHandleReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ReportIDParamInvalid",
			req: &report.AdminActivityReportHandleReq{
				AccessToken: "111",
				ReportID:    "aaa",
			},
			errorIsExist:  true,
			expectedError: "举报ID不合法",
		},
		{
			name: "QueryFail",
			req: &report.AdminActivityReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ReportIsNotExist",
			req: &report.AdminActivityReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
			},
			errorIsExist:  true,
			expectedError: "举报不存在或已经处理",
		},
		{
			name: "ActionTypeParamInvalid",
			req: &report.AdminActivityReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				ActionType:  666,
			},
			errorIsExist:                true,
			expectedError:               "操作类型不合法",
			mockQueryReportResultReturn: true,
		},
		{
			name: "UpdateFail",
			req: &report.AdminActivityReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				ActionType:  common.ActionTypeOff,
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockQueryReportResultReturn: true,
			mockUpdateReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &report.AdminActivityReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist:                false,
			mockQueryReportResultReturn: true,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(exquery.QueryActivityReportExistByIdAndStatus).Return(tc.mockQueryReportResultReturn, tc.mockQueryReportErrorReturn).Build()
			mockey.Mock(exquery.UpdateActivityReportById).Return(tc.mockUpdateReportErrorReturn).Build()

			err := reportService.NewAdminActivityReportHandleEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewAdminCommentReportHandleEvent(t *testing.T) {
	type testCase struct {
		name                        string
		req                         *report.AdminCommentReportHandleReq
		errorIsExist                bool
		expectedError               string
		mockQueryReportErrorReturn  error
		mockQueryReportResultReturn bool
		mockAccessTokenErrorReturn  error
		mockAccessTokenResultReturn int64
		mockUpdateReportErrorReturn error
	}

	testCases := []testCase{
		{
			name: "ActivityIDParamInvalid",
			req: &report.AdminCommentReportHandleReq{
				CommentType: "aaa",
			},
			errorIsExist:  true,
			expectedError: "评论类型错误",
		},
		{
			name: "VideoAccessTokenFail",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ActivityAccessTokenFail",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "VideoReportIDParamInvalid",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "aaa",
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:  true,
			expectedError: "举报ID不合法",
		},
		{
			name: "ActivityReportIDParamInvalid",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "aaa",
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:  true,
			expectedError: "举报ID不合法",
		},
		{
			name: "VideoQueryFail",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ActivityQueryFail",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoReportIsNotExist",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:  true,
			expectedError: "举报不存在或已经处理",
		},
		{
			name: "ActivityReportIsNotExist",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:  true,
			expectedError: "举报不存在或已经处理",
		},
		{
			name: "VideoActionTypeParamInvalid",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeVideo,
				ActionType:  3,
			},
			errorIsExist:                true,
			expectedError:               "操作类型不合法",
			mockQueryReportResultReturn: true,
		},
		{
			name: "ActivityActionTypeParamInvalid",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeActivity,
				ActionType:  3,
			},
			errorIsExist:                true,
			expectedError:               "操作类型不合法",
			mockQueryReportResultReturn: true,
		},
		{
			name: "VideoUpdateFail",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockQueryReportResultReturn: true,
			mockUpdateReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ActivityUpdateFail",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockQueryReportResultReturn: true,
			mockUpdateReportErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoSuccess",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeVideo,
			},
			errorIsExist:                false,
			mockQueryReportResultReturn: true,
		},
		{
			name: "ActivitySuccess",
			req: &report.AdminCommentReportHandleReq{
				AccessToken: "111",
				ReportID:    "111",
				CommentType: common.CommentTypeActivity,
			},
			errorIsExist:                false,
			mockQueryReportResultReturn: true,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
			switch tc.req.CommentType {
			case common.CommentTypeVideo:
				mockey.Mock(exquery.QueryVideoCommentReportExistByIdAndStatus).Return(tc.mockQueryReportResultReturn, tc.mockQueryReportErrorReturn).Build()
				mockey.Mock(exquery.UpdateVideoCommentReportById).Return(tc.mockUpdateReportErrorReturn).Build()
			case common.CommentTypeActivity:
				mockey.Mock(exquery.QueryActivityCommentReportExistByIdAndStatus).Return(tc.mockQueryReportResultReturn, tc.mockQueryReportErrorReturn).Build()
				mockey.Mock(exquery.UpdateActivityCommentReportById).Return(tc.mockUpdateReportErrorReturn).Build()
			}

			err := reportService.NewAdminCommentReportHandleEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewAdminVideoListEvent(t *testing.T) {
	type testCase struct {
		name                    string
		req                     *report.AdminVideoListReq
		errorIsExist            bool
		expectedError           string
		mockQueryErrorReturn    error
		mockQueryResultReturn   []*model.Video
		mockQueryCountReturn    int64
		mockConvertErrorReturn  error
		mockConvertResultReturn []*base.Video
		expectedResult          *report.AdminVideoListRespData
	}

	category := "运动"
	testCases := []testCase{
		{
			name: "CategoryIsNotExist",
			req: &report.AdminVideoListReq{
				Category: new(string),
			},
			errorIsExist:  true,
			expectedError: "分类不存在",
		},
		{
			name: "QueryFail",
			req: &report.AdminVideoListReq{
				Category: &category,
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name:                   "ConvertFail",
			req:                    &report.AdminVideoListReq{},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
			mockQueryResultReturn:  make([]*model.Video, 0),
		},
		{
			name:                  "Success",
			req:                   &report.AdminVideoListReq{},
			errorIsExist:          false,
			mockQueryResultReturn: make([]*model.Video, 0),
			expectedResult: &report.AdminVideoListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
		},
	}

	checker.CategoryMap["运动"] = 1

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoByCategoryPaged).Return(tc.mockQueryResultReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := reportService.NewAdminVideoListEvent(tc.req)

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

func TestNewAdminVideoHandleEvent(t *testing.T) {
	type testCase struct {
		name                  string
		req                   *report.AdminVideoHandleReq
		errorIsExist          bool
		expectedError         string
		mockQueryErrorReturn  error
		mockQueryResultReturn bool
		mockUpdateErrorReturn error
		expectedResult        *report.AdminVideoListRespData
	}

	testCases := []testCase{
		{
			name: "VideoIDInvalid",
			req: &report.AdminVideoHandleReq{
				VideoID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "视频ID不合法",
		},
		{
			name: "QueryFail",
			req: &report.AdminVideoHandleReq{
				VideoID: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoIsNotExist",
			req: &report.AdminVideoHandleReq{
				VideoID: "111",
			},
			errorIsExist:  true,
			expectedError: "视频不存在或已经处理",
		},
		{
			name: "ActionTypeInvalid",
			req: &report.AdminVideoHandleReq{
				VideoID:    "111",
				ActionType: 3,
			},
			errorIsExist:          true,
			expectedError:         "操作类型不合法",
			mockQueryResultReturn: true,
		},
		{
			name: "UpdateFail",
			req: &report.AdminVideoHandleReq{
				VideoID:    "111",
				ActionType: common.ActionTypeOff,
			},
			errorIsExist:          true,
			expectedError:         errno.DatabaseCallErrorMsg,
			mockQueryResultReturn: true,
			mockUpdateErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &report.AdminVideoHandleReq{
				VideoID:    "111",
				ActionType: common.ActionTypeOn,
			},
			errorIsExist:          false,
			mockQueryResultReturn: true,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoExistByIdAndStatus).Return(tc.mockQueryResultReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock(exquery.UpdateVideoWithId).Return(tc.mockUpdateErrorReturn).Build()

			err := reportService.NewAdminVideoHandleEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
