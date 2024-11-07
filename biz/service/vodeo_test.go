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

func TestNewPublishEvent(t *testing.T) {
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

func TestNewFeedEvent(t *testing.T) {
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
