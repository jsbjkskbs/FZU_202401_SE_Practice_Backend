package service

import (
	"github.com/bytedance/mockey"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stretchr/testify/assert"
	"sfw/biz/model/api/interact"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/service/common"
	"sfw/pkg/errno"
	"sfw/pkg/utils/scheduler"
	"testing"
	"time"
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
