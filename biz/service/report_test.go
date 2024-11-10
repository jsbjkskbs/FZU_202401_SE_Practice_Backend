package service

import (
	"testing"

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
