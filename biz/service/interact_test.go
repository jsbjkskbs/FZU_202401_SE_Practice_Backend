package service

import (
	"testing"
	"time"

	"github.com/bytedance/mockey"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stretchr/testify/assert"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/interact"
	"sfw/biz/model/base"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/scheduler"
)

var interactService = NewInteractService(nil, new(app.RequestContext))

func TestNewLikeVideoActionEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *interact.InteractLikeVideoActionReq
		errorIsExist               bool
		expectedError              string
		mockAccessTokenErrorReturn error
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &interact.InteractLikeVideoActionReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ParamInvalid",
			req: &interact.InteractLikeVideoActionReq{
				AccessToken: "111",
				ActionType:  4,
			},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name: "SuccessAppendVideoLikeInfo",
			req: &interact.InteractLikeVideoActionReq{
				AccessToken: "111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist: false,
		},
		{
			name: "SuccessRemoveVideoLikeInfo",
			req: &interact.InteractLikeVideoActionReq{
				AccessToken: "111",
				ActionType:  common.ActionTypeOff,
			},
			errorIsExist: false,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(redis.AppendVideoLikeInfo).Return(nil).Build()
			mockey.Mock(redis.RemoveVideoLikeInfo).Return(nil).Build()
			mockey.Mock((*scheduler.Scheduler).Start).Return().Build()

			err := interactService.NewLikeVideoActionEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				time.Sleep(1 * time.Second)
			}
		})
	}
}

func TestNewLikeActivityActionEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *interact.InteractLikeActivityActionReq
		errorIsExist               bool
		expectedError              string
		mockAccessTokenErrorReturn error
	}

	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &interact.InteractLikeActivityActionReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ParamInvalid",
			req: &interact.InteractLikeActivityActionReq{
				AccessToken: "111",
				ActionType:  4,
			},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name: "SuccessAppendActivityLikeInfo",
			req: &interact.InteractLikeActivityActionReq{
				AccessToken: "111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist: false,
		},
		{
			name: "SuccessRemoveActivityLikeInfo",
			req: &interact.InteractLikeActivityActionReq{
				AccessToken: "111",
				ActionType:  common.ActionTypeOff,
			},
			errorIsExist: false,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(redis.AppendActivityLikeInfo).Return(nil).Build()
			mockey.Mock(redis.RemoveActivityLikeInfo).Return(nil).Build()
			mockey.Mock((*scheduler.Scheduler).Start).Return().Build()

			err := interactService.NewLikeActivityActionEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				time.Sleep(1 * time.Second)
			}
		})
	}
}

func TestNewLikeCommentEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *interact.InteractLikeCommentActionReq
		errorIsExist               bool
		expectedError              string
		mockAccessTokenErrorReturn error
	}

	testCases := []testCase{
		{
			name: "CommentTypeParamInvalid",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: "aaa",
			},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name: "VideoAccessTokenFail",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeVideo,
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "VideoActionTypeParamInvalid",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeVideo,
				AccessToken: "111",
				ActionType:  4,
			},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name: "SuccessAppendVideoCommentLikeInfo",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeVideo,
				AccessToken: "111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist: false,
		},
		{
			name: "SuccessRemoveVideoCommentLikeInfo",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeVideo,
				AccessToken: "111",
				ActionType:  common.ActionTypeOff,
			},
			errorIsExist: false,
		},
		{
			name: "ActivityAccessTokenFail",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeActivity,
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ActivityActionTypeParamInvalid",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeActivity,
				AccessToken: "111",
				ActionType:  4,
			},
			errorIsExist:  true,
			expectedError: errno.ParamInvalidErrorMsg,
		},
		{
			name: "SuccessAppendActivityCommentLikeInfo",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeVideo,
				AccessToken: "111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist: false,
		},
		{
			name: "SuccessRemoveActivityCommentLikeInfo",
			req: &interact.InteractLikeCommentActionReq{
				CommentType: common.CommentTypeVideo,
				AccessToken: "111",
				ActionType:  common.ActionTypeOff,
			},
			errorIsExist: false,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(redis.AppendVideoCommentLikeInfo).Return(nil).Build()
			mockey.Mock(redis.RemoveVideoCommentLikeInfo).Return(nil).Build()
			mockey.Mock(redis.AppendActivityCommentLikeInfo).Return(nil).Build()
			mockey.Mock(redis.RemoveActivityCommentLikeInfo).Return(nil).Build()
			mockey.Mock((*scheduler.Scheduler).Start).Return().Build()

			err := interactService.NewLikeCommentEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				time.Sleep(1 * time.Second)
			}
		})
	}
}

func TestNewLikeVideoListEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *interact.InteractLikeVideoListReq
		errorIsExist               bool
		expectedError              string
		expectedResult             *interact.InteractLikeVideoListRespData
		mockQueryErrorReturn       error
		mockQueryVideoReturn       []*model.Video
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    []*base.Video
	}

	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: &interact.InteractLikeVideoListReq{
				UserID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "无效的用户ID",
		},
		{
			name: "QueryFail",
			req: &interact.InteractLikeVideoListReq{
				UserID: "111",
			},
			errorIsExist:         true,
			expectedError:        errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "AccessTokenFail",
			req: &interact.InteractLikeVideoListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ConvertFail",
			req: &interact.InteractLikeVideoListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &interact.InteractLikeVideoListReq{
				UserID: "111",
			},
			errorIsExist: false,
			expectedResult: &interact.InteractLikeVideoListRespData{
				Items:    []*base.Video{},
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

			mockey.Mock(exquery.QueryVideoLikedByUserIdPaged).Return(tc.mockQueryVideoReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := interactService.NewLikeVideoListEvent(tc.req)

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

func TestNewCommentVideoPublishEvent(t *testing.T) {
	type testCase struct {
		name                         string
		req                          *interact.InteractCommentVideoPublishReq
		errorIsExist                 bool
		expectedError                string
		mockQueryRootIdErrorReturn   error
		mockQueryRootIdExistReturn   bool
		mockQueryParentIdErrorReturn error
		mockQueryParentIdExistReturn bool
		mockQueryVideoErrorReturn    error
		mockQueryVideoExistReturn    bool
		mockInsertErrorReturn        error
		mockAccessTokenErrorReturn   error
	}

	rootID := "111"
	wrongID := "0"
	parentID := "111"
	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "RootIDParamInvalid",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				RootID:      new(string),
			},
			errorIsExist:  true,
			expectedError: "无效的根评论ID",
		},
		{
			name: "WrongRootID",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				RootID:      &wrongID,
			},
			errorIsExist:  true,
			expectedError: "无效的根评论ID",
		},
		{
			name: "QueryRootIdFail",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				RootID:      &rootID,
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryRootIdErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "RootIdIsNotExist",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				RootID:      &rootID,
			},
			errorIsExist:  true,
			expectedError: "根评论不存在",
		},
		{
			name: "ParamInvalid",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				ParentID:    new(string),
			},
			errorIsExist:  true,
			expectedError: "父评论ID必须与根评论ID同时存在",
		},
		{
			name: "ParentIDParamInvalid",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				ParentID:    new(string),
				RootID:      &rootID,
			},
			errorIsExist:               true,
			expectedError:              "无效的父评论ID",
			mockQueryRootIdExistReturn: true,
		},
		{
			name: "WrongParentID",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				ParentID:    &wrongID,
				RootID:      &rootID,
			},
			errorIsExist:               true,
			expectedError:              "无效的父评论ID",
			mockQueryRootIdExistReturn: true,
		},
		{
			name: "QueryParentIDFail",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				ParentID:    &parentID,
				RootID:      &rootID,
			},
			errorIsExist:                 true,
			expectedError:                errno.DatabaseCallErrorMsg,
			mockQueryParentIdErrorReturn: errno.DatabaseCallError,
			mockQueryRootIdExistReturn:   true,
		},
		{
			name: "ParentIdIsNotExist",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				ParentID:    &parentID,
				RootID:      &rootID,
			},
			errorIsExist:               true,
			expectedError:              "父评论不存在",
			mockQueryRootIdExistReturn: true,
		},
		{
			name: "VideoIDParamInvalid",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				VideoID:     "aa",
			},
			errorIsExist:  true,
			expectedError: "无效的视频ID",
		},
		{
			name: "QueryVideoFail",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				VideoID:     "111",
			},
			errorIsExist:              true,
			expectedError:             errno.DatabaseCallErrorMsg,
			mockQueryVideoErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoIsNotExist",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				VideoID:     "111",
			},
			errorIsExist:  true,
			expectedError: "视频不存在",
		},
		{
			name: "InsertFail",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				VideoID:     "111",
			},
			errorIsExist:              true,
			expectedError:             errno.DatabaseCallErrorMsg,
			mockInsertErrorReturn:     errno.DatabaseCallError,
			mockQueryVideoExistReturn: true,
		},
		{
			name: "Success",
			req: &interact.InteractCommentVideoPublishReq{
				AccessToken: "111",
				VideoID:     "111",
			},
			errorIsExist:              false,
			mockQueryVideoExistReturn: true,
		},
	}

	generator.VideoCommentIDGenerator, _ = generator.NewSnowflake(4)

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoCommentExistByIdParentIdAndRootId).Return(tc.mockQueryRootIdExistReturn, tc.mockQueryRootIdErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(exquery.QueryVideoCommentExistByIdAndRootId).Return(tc.mockQueryParentIdExistReturn, tc.mockQueryParentIdErrorReturn).Build()
			mockey.Mock(exquery.QueryVideoExistById).Return(tc.mockQueryVideoExistReturn, tc.mockQueryVideoErrorReturn).Build()
			mockey.Mock(exquery.InsertVideoComment).Return(tc.mockInsertErrorReturn).Build()
			mockey.Mock(generator.VideoCommentIDGenerator.Generate).Return(111).Build()

			err := interactService.NewCommentVideoPublishEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewCommentVideoListEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *interact.InteractCommentVideoListReq
		errorIsExist               bool
		expectedError              string
		mockQueryErrorReturn       error
		mockQueryCommentReturn     []*model.VideoComment
		mockQueryCountReturn       int64
		mockAccessTokenErrorReturn error
		mockQueryVideoErrorReturn  error
		mockQueryVideoExistReturn  bool
		mockConvertErrorReturn     error
		mockConvertResultReturn    *[]*base.Comment
		expectedResult             *interact.InteractCommentVideoListRespData
	}

	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: &interact.InteractCommentVideoListReq{
				VideoID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "无效的视频ID",
		},
		{
			name: "QueryVideoFail",
			req: &interact.InteractCommentVideoListReq{
				VideoID: "111",
			},
			errorIsExist:              true,
			expectedError:             errno.DatabaseCallErrorMsg,
			mockQueryVideoErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "VideoIsNotExist",
			req: &interact.InteractCommentVideoListReq{
				VideoID: "111",
			},
			errorIsExist:  true,
			expectedError: "视频不存在",
		},
		{
			name: "QueryCommentFail",
			req: &interact.InteractCommentVideoListReq{
				VideoID: "111",
			},
			errorIsExist:              true,
			expectedError:             errno.DatabaseCallErrorMsg,
			mockQueryErrorReturn:      errno.DatabaseCallError,
			mockQueryVideoExistReturn: true,
		},
		{
			name: "AccessTokenFail",
			req: &interact.InteractCommentVideoListReq{
				VideoID:     "111",
				AccessToken: new(string),
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
			mockQueryVideoExistReturn:  true,
		},
		{
			name: "ConvertFail",
			req: &interact.InteractCommentVideoListReq{
				VideoID:     "111",
				AccessToken: new(string),
			},
			errorIsExist:              true,
			expectedError:             errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn:    errno.DatabaseCallError,
			mockQueryVideoExistReturn: true,
		},
		{
			name: "Success",
			req: &interact.InteractCommentVideoListReq{
				VideoID: "111",
			},
			errorIsExist: false,
			expectedResult: &interact.InteractCommentVideoListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn:   new([]*base.Comment),
			mockQueryVideoExistReturn: true,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryVideoExistById).Return(tc.mockQueryVideoExistReturn, tc.mockQueryVideoErrorReturn).Build()
			mockey.Mock(exquery.QueryVideoRootCommentByVideoIdPaged).Return(tc.mockQueryCommentReturn, tc.mockQueryCountReturn, tc.mockQueryErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.VideoCommentDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()

			result, err := interactService.NewCommentVideoListEvent(tc.req)

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
