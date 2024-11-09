package service

import (
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"sfw/biz/service/model_converter"
	"sfw/pkg/utils/scheduler"

	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/video"
	"sfw/biz/model/base"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/pkg/errno"
	"sfw/pkg/oss"
	"sfw/pkg/utils/checker"
	"sfw/pkg/utils/generator"
)

var videoService = NewVideoService(nil, new(app.RequestContext))

func TestNewVideoPublishEvent(t *testing.T) {
	type testCase struct {
		name                      string
		req                       *video.VideoPublishReq
		errorIsExist              bool
		expectedError             string
		mockTokenErrorReturn      error
		mockInfoStoreErrorReturn  error
		mockUploadErrorReturn     error
		mockUploadUptokenReturn   string
		mockUploadUploadKeyReturn string
		expectedResult            *video.VideoPublishRespData
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoPublishReq{
				AccessToken: "1111",
			},
			errorIsExist:         true,
			expectedError:        errno.AccessTokenInvalidErrorMsg,
			mockTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ParamInvalid",
			req: &video.VideoPublishReq{
				AccessToken: "1111",
				Title:       "111",
				Description: "111",
				Category:    "运动",
				Labels:      make([]string, 0),
			},
			errorIsExist:  true,
			expectedError: "视频信息不合法",
		},
		{
			name: "InfoStoreFail",
			req: &video.VideoPublishReq{
				AccessToken: "1111",
				Title:       "111",
				Description: "111",
				Category:    "运动",
				Labels:      make([]string, 1),
			},
			errorIsExist:             true,
			expectedError:            errno.DatabaseCallErrorMsg,
			mockInfoStoreErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "UploadFail",
			req: &video.VideoPublishReq{
				AccessToken: "1111",
				Title:       "111",
				Description: "111",
				Category:    "运动",
				Labels:      make([]string, 1),
			},
			errorIsExist:          true,
			expectedError:         errno.InternalServerErrorMsg,
			mockUploadErrorReturn: errno.InternalServerError,
		},
		{
			name: "Success",
			req: &video.VideoPublishReq{
				AccessToken: "1111",
				Title:       "111",
				Description: "111",
				Category:    "运动",
				Labels:      make([]string, 1),
			},
			errorIsExist:              false,
			mockUploadUptokenReturn:   "111",
			mockUploadUploadKeyReturn: "111",
			expectedResult: &video.VideoPublishRespData{
				UploadURL: oss.UploadUrl,
				UploadKey: "111",
				Uptoken:   "111",
			},
		},
	}

	generator.VideoIDGenerator, _ = generator.NewSnowflake(2)
	checker.CategoryMap["运动"] = 1

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockTokenErrorReturn).Build()
			mockey.Mock((*generator.Snowflake).Generate).Return(111).Build()
			mockey.Mock(redis.VideoUploadInfoStore).Return(tc.mockInfoStoreErrorReturn).Build()
			mockey.Mock(oss.UploadVideo).Return(tc.mockUploadUptokenReturn, tc.mockUploadUploadKeyReturn, tc.mockUploadErrorReturn).Build()

			result, err := videoService.NewPublishEvent(tc.req)

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

func TestNewCoverUploadEvent(t *testing.T) {
	type testCase struct {
		name                      string
		req                       *video.VideoCoverUploadReq
		errorIsExist              bool
		expectedError             string
		mockTokenIdReturn         int64
		mockTokenErrorReturn      error
		mockQueryErrorReturn      error
		mockQueryExistReturn      bool
		mockUploadUptokenReturn   string
		mockUploadUploadKeyReturn string
		mockUploadErrorReturn     error
		expectedResult            *video.VideoCoverUploadRespData
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoCoverUploadReq{
				AccessToken: "1111",
			},
			errorIsExist:         true,
			expectedError:        errno.AccessTokenInvalidErrorMsg,
			mockTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ParamInvalid",
			req: &video.VideoCoverUploadReq{
				AccessToken: "1111",
				VideoID:     "aaa",
			},
			errorIsExist:  true,
			expectedError: "视频ID错误",
		},
		{
			name: "QueryFail",
			req: &video.VideoCoverUploadReq{
				AccessToken: "1111",
				VideoID:     "1",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ResourceNotFound",
			req: &video.VideoCoverUploadReq{
				AccessToken: "1111",
				VideoID:     "1",
			},
			errorIsExist:  true,
			expectedError: "视频不存在",
		},
		{
			name: "UploadFail",
			req: &video.VideoCoverUploadReq{
				AccessToken: "1111",
				VideoID:     "1",
			},
			errorIsExist:          true,
			expectedError:         errno.InternalServerErrorMsg,
			mockQueryExistReturn:  true,
			mockUploadErrorReturn: errno.InternalServerError,
		},
		{
			name: "Success",
			req: &video.VideoCoverUploadReq{
				AccessToken: "1111",
				VideoID:     "1",
			},
			errorIsExist:              false,
			mockQueryExistReturn:      true,
			mockUploadUptokenReturn:   "111",
			mockUploadUploadKeyReturn: "111",
			expectedResult: &video.VideoCoverUploadRespData{
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
			mockey.Mock(exquery.QueryVideoExistByIdAndUserId).Return(tc.mockQueryExistReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock(oss.UploadVideoCover).Return(tc.mockUploadUptokenReturn, tc.mockUploadUploadKeyReturn, tc.mockUploadErrorReturn).Build()

			result, err := videoService.NewCoverUploadEvent(tc.req)

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

func TestNewVideoFeedEvent(t *testing.T) {
	type testCase struct {
		name                     string
		req                      *video.VideoFeedReq
		errorIsExist             bool
		expectedError            string
		mockRecommendVidReturn   []string
		mockRecommendErrorReturn error
		mockQueryErrorReturn     error
		mockQueryVideoReturn     *model.Video
		expectedResult           []*base.Video
	}

	category := "运动"
	vid := make([]string, 0)
	vid = append(vid, "1")
	testCases := []testCase{
		{
			name: "GetRecommendFail",
			req: &video.VideoFeedReq{
				Category: &category,
			},
			errorIsExist:             true,
			expectedError:            errno.InternalServerErrorMsg,
			mockRecommendErrorReturn: errno.InternalServerError,
		},
		{
			name:                     "GetRecommendWithCategoryFail",
			req:                      &video.VideoFeedReq{},
			errorIsExist:             true,
			expectedError:            errno.InternalServerErrorMsg,
			mockRecommendErrorReturn: errno.InternalServerError,
		},
		{
			name:                   "ParseIntFail",
			req:                    &video.VideoFeedReq{},
			errorIsExist:           true,
			expectedError:          errno.InternalServerErrorMsg,
			mockRecommendVidReturn: make([]string, 1),
		},
		{
			name:                   "ResourceNotFound",
			req:                    &video.VideoFeedReq{},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockRecommendVidReturn: vid,
			mockQueryErrorReturn:   errno.DatabaseCallError,
		},
		{
			name:                   "UploadFail",
			req:                    &video.VideoFeedReq{},
			errorIsExist:           true,
			expectedError:          "视频不存在",
			mockRecommendVidReturn: vid,
		},
		{
			name:                   "Success",
			req:                    &video.VideoFeedReq{},
			errorIsExist:           false,
			mockRecommendVidReturn: make([]string, 0),
			mockQueryVideoReturn:   &model.Video{},
			expectedResult:         []*base.Video{},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			if tc.req.Category != nil {
				mockey.Mock(gorse.GetRecommendWithCategory).Return(tc.mockRecommendVidReturn, tc.mockRecommendErrorReturn).Build()
			} else {
				mockey.Mock(gorse.GetRecommend).Return(tc.mockRecommendVidReturn, tc.mockRecommendErrorReturn).Build()
			}
			mockey.Mock(exquery.QueryVideoById).Return(tc.mockQueryVideoReturn, tc.mockQueryErrorReturn).Build()

			result, err := videoService.NewFeedEvent(tc.req)

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

func TestNewCustomFeedEvent(t *testing.T) {
	type testCase struct {
		name                     string
		req                      *video.VideoCustomFeedReq
		errorIsExist             bool
		expectedError            string
		mockTokenErrorReturn     error
		mockRecommendVidReturn   []string
		mockRecommendErrorReturn error
		mockQueryErrorReturn     error
		mockQueryVideoReturn     *model.Video
		expectedResult           []*base.Video
	}

	category := "运动"
	vid := make([]string, 0)
	vid = append(vid, "1")
	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoCustomFeedReq{
				AccessToken: "1111",
			},
			errorIsExist:         true,
			expectedError:        errno.AccessTokenInvalidErrorMsg,
			mockTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "GetRecommendFail",
			req: &video.VideoCustomFeedReq{
				AccessToken: "111",
				Category:    &category,
			},
			errorIsExist:             true,
			expectedError:            errno.InternalServerErrorMsg,
			mockRecommendErrorReturn: errno.InternalServerError,
		},
		{
			name: "GetRecommendWithCategoryFail",
			req: &video.VideoCustomFeedReq{
				AccessToken: "111",
			},
			errorIsExist:             true,
			expectedError:            errno.InternalServerErrorMsg,
			mockRecommendErrorReturn: errno.InternalServerError,
		},
		{
			name: "ParseIntFail",
			req: &video.VideoCustomFeedReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          errno.InternalServerErrorMsg,
			mockRecommendVidReturn: make([]string, 1),
		},
		{
			name: "ResourceNotFound",
			req: &video.VideoCustomFeedReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockRecommendVidReturn: vid,
			mockQueryErrorReturn:   errno.DatabaseCallError,
		},
		{
			name: "UploadFail",
			req: &video.VideoCustomFeedReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          "视频不存在",
			mockRecommendVidReturn: vid,
		},
		{
			name: "Success",
			req: &video.VideoCustomFeedReq{
				AccessToken: "111",
			},
			errorIsExist:           false,
			mockRecommendVidReturn: make([]string, 0),
			mockQueryVideoReturn:   &model.Video{},
			expectedResult:         []*base.Video{},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockTokenErrorReturn).Build()
			if tc.req.Category != nil {
				mockey.Mock(gorse.GetRecommendWithCategory).Return(tc.mockRecommendVidReturn, tc.mockRecommendErrorReturn).Build()
			} else {
				mockey.Mock(gorse.GetRecommend).Return(tc.mockRecommendVidReturn, tc.mockRecommendErrorReturn).Build()
			}
			mockey.Mock(exquery.QueryVideoById).Return(tc.mockQueryVideoReturn, tc.mockQueryErrorReturn).Build()

			result, err := videoService.NewCustomFeedEvent(tc.req)

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

func TestNewCategoriesEvent(t *testing.T) {
	type testCase struct {
		name           string
		req            *video.VideoCategoriesReq
		errorIsExist   bool
		expectedError  string
		expectedResult []string
	}

	category := "运动"
	var categories []string
	categories = make([]string, 0)
	categories = append(categories, category)

	testCases := []testCase{
		{
			name:           "Success",
			req:            &video.VideoCategoriesReq{},
			errorIsExist:   false,
			expectedResult: categories,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.MockValue(&checker.Categories).To(categories)

			result, err := videoService.NewCategoriesEvent(tc.req)

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

func TestNewVideoInfoEvent(t *testing.T) {
	type testCase struct {
		name                          string
		req                           *video.VideoInfoReq
		errorIsExist                  bool
		expectedError                 string
		expectedResult                *base.Video
		mockQueryErrorReturn          error
		mockQueryVideoReturn          *model.Video
		mockAccessTokenErrorReturn    error
		mockConvertErrorReturn        error
		mockVisited                   bool
		mockVisitedErrorReturn        error
		mockVisitCountErrorReturn     error
		mockPutIPVisitInfoErrorReturn error
		mockConvertResultReturn       *base.Video
	}

	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: &video.VideoInfoReq{
				VideoID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "视频ID错误",
		},
		{
			name: "QueryFail",
			req: &video.VideoInfoReq{
				VideoID: "1",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "IsIPVisitedFail",
			req: &video.VideoInfoReq{
				VideoID: "1",
			},
			errorIsExist:            false,
			expectedResult:          &base.Video{},
			mockVisited:             true,
			mockConvertResultReturn: &base.Video{},
			mockVisitedErrorReturn:  errno.DatabaseCallError,
		},
		{
			name: "VisitCountFail",
			req: &video.VideoInfoReq{
				VideoID: "1",
			},
			errorIsExist:              false,
			expectedResult:            &base.Video{},
			mockVisited:               true,
			mockConvertResultReturn:   &base.Video{},
			mockVisitCountErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "PutIPVisitInfoFail",
			req: &video.VideoInfoReq{
				VideoID: "1",
			},
			errorIsExist:                  false,
			expectedResult:                &base.Video{},
			mockVisited:                   true,
			mockConvertResultReturn:       &base.Video{},
			mockPutIPVisitInfoErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &video.VideoInfoReq{
				VideoID: "1",
			},
			errorIsExist:            false,
			expectedResult:          &base.Video{},
			mockConvertResultReturn: &base.Video{},
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoById).Return(tc.mockQueryVideoReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()
			mockey.Mock(redis.IsIPVisited).Return(tc.mockVisited, tc.mockVisitedErrorReturn).Build()
			mockey.Mock(videoService.c.ClientIP).Return("1").Build()
			mockey.Mock(redis.IncrVideoVisitCount).Return(tc.mockVisitCountErrorReturn).Build()
			mockey.Mock((*scheduler.Scheduler).Start).Return().Build()
			mockey.Mock(redis.PutIPVisitInfo).Return(tc.mockVisitCountErrorReturn).Build()

			result, err := videoService.NewInfoEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
			time.Sleep(1 * time.Second)
		})
	}
}

func TestNewUserListEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *video.VideoListReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *video.VideoListRespData
		mockQueryErrorReturn       error
		mockQueryResultReturn      []*model.Video
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    []*base.Video
	}

	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: &video.VideoListReq{
				UserID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "用户ID错误",
		},
		{
			name: "QueryFail",
			req: &video.VideoListReq{
				UserID: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "AccessTokenFail",
			req: &video.VideoListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ConvertFail",
			req: &video.VideoListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:           true,
			expectedError:          errno.InternalServerErrorMsg,
			mockConvertErrorReturn: errno.InternalServerError,
		},
		{
			name: "Success",
			req: &video.VideoListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist: false,
			expectedResult: &video.VideoListRespData{
				Items:    make([]*base.Video, 0),
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: make([]*base.Video, 0),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoByUserIdAndStatusPaged).Return(tc.mockQueryResultReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := videoService.NewListEvent(tc.req)

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

func TestNewSubmitAllEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *video.VideoSubmitAllReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *video.VideoSubmitAllRespData
		mockQueryErrorReturn       error
		mockQueryVideoReturn       []*model.Video
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    []*base.Video
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoSubmitAllReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "QueryFail",
			req: &video.VideoSubmitAllReq{
				AccessToken: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ConvertFail",
			req: &video.VideoSubmitAllReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          errno.InternalServerErrorMsg,
			mockConvertErrorReturn: errno.InternalServerError,
		},
		{
			name: "Success",
			req: &video.VideoSubmitAllReq{
				AccessToken: "111",
			},
			errorIsExist: false,
			expectedResult: &video.VideoSubmitAllRespData{
				Items:    make([]*base.Video, 0),
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: make([]*base.Video, 0),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoByUserIdPaged).Return(tc.mockQueryVideoReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := videoService.NewSubmitAllEvent(tc.req)

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

func TestNewSubmitReviewEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *video.VideoSubmitReviewReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *video.VideoSubmitReviewRespData
		mockQueryErrorReturn       error
		mockQueryVideoReturn       []*model.Video
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    []*base.Video
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoSubmitReviewReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "QueryFail",
			req: &video.VideoSubmitReviewReq{
				AccessToken: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ConvertFail",
			req: &video.VideoSubmitReviewReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &video.VideoSubmitReviewReq{
				AccessToken: "111",
			},
			errorIsExist: false,
			expectedResult: &video.VideoSubmitReviewRespData{
				Items:    make([]*base.Video, 0),
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: make([]*base.Video, 0),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoByUserIdAndStatusPaged).Return(tc.mockQueryVideoReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := videoService.NewSubmitReviewEvent(tc.req)

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

func TestNewSubmitLockedEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *video.VideoSubmitLockedReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *video.VideoSubmitLockedRespData
		mockQueryErrorReturn       error
		mockQueryVideoReturn       []*model.Video
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    []*base.Video
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoSubmitLockedReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "QueryFail",
			req: &video.VideoSubmitLockedReq{
				AccessToken: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ConvertFail",
			req: &video.VideoSubmitLockedReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          errno.InternalServerErrorMsg,
			mockConvertErrorReturn: errno.InternalServerError,
		},
		{
			name: "Success",
			req: &video.VideoSubmitLockedReq{
				AccessToken: "111",
			},
			errorIsExist: false,
			expectedResult: &video.VideoSubmitLockedRespData{
				Items:    make([]*base.Video, 0),
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: make([]*base.Video, 0),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoByUserIdAndStatusPaged).Return(tc.mockQueryVideoReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := videoService.NewSubmitLockedEvent(tc.req)

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

func TestNewSumitPassedEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *video.VideoSubmitPassedReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *video.VideoSubmitPassedRespData
		mockQueryErrorReturn       error
		mockQueryVideoReturn       []*model.Video
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    []*base.Video
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoSubmitPassedReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "QueryFail",
			req: &video.VideoSubmitPassedReq{
				AccessToken: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ConvertFail",
			req: &video.VideoSubmitPassedReq{
				AccessToken: "111",
			},
			errorIsExist:           true,
			expectedError:          errno.InternalServerErrorMsg,
			mockConvertErrorReturn: errno.InternalServerError,
		},
		{
			name: "Success",
			req: &video.VideoSubmitPassedReq{
				AccessToken: "111",
			},
			errorIsExist: false,
			expectedResult: &video.VideoSubmitPassedRespData{
				Items:    make([]*base.Video, 0),
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: make([]*base.Video, 0),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoByUserIdAndStatusPaged).Return(tc.mockQueryVideoReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := videoService.NewSumitPassedEvent(tc.req)

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

func TestNewVideoSearchEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *video.VideoSearchReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *video.VideoSearchRespData
		mockQueryErrorReturn       error
		mockQueryVideoReturn       []*model.Video
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    []*base.Video
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &video.VideoSearchReq{
				AccessToken: new(string),
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "QueryFail",
			req: &video.VideoSearchReq{
				AccessToken: new(string),
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "ConvertFail",
			req: &video.VideoSearchReq{
				AccessToken: new(string),
			},
			errorIsExist:           true,
			expectedError:          errno.InternalServerErrorMsg,
			mockConvertErrorReturn: errno.InternalServerError,
		},
		{
			name: "Success",
			req: &video.VideoSearchReq{
				AccessToken: new(string),
			},
			errorIsExist: false,
			expectedResult: &video.VideoSearchRespData{
				Items:    make([]*base.Video, 0),
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: make([]*base.Video, 0),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoFuzzyByKeywordPaged).Return(tc.mockQueryVideoReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := videoService.NewSearchEvent(tc.req)

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
