package service

import (
	"testing"

	"github.com/bytedance/mockey"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stretchr/testify/assert"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/relation"
	"sfw/biz/model/base"
	"sfw/biz/mw/jwt"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
)

var relationService = NewRelationService(nil, new(app.RequestContext))

func TestNewFollowActionEvent(t *testing.T) {
	type testCase struct {
		name                        string
		req                         *relation.RelationFollowActionReq
		errorIsExist                bool
		expectedError               string
		mockQueryUserErrorReturn    error
		mockQueryUserResultReturn   bool
		mockDeleteFollowErrorReturn error
		mockQueryFollowErrorReturn  error
		mockQueryFollowResultReturn bool
		mockInsertFollowErrorReturn error
		mockAccessTokenErrorReturn  error
		mockAccessTokenResultReturn int64
	}

	testCases := []testCase{
		{
			name: "ToUserIDParamInvalid",
			req: &relation.RelationFollowActionReq{
				ToUserID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "用户ID不合法",
		},
		{
			name: "Can'tFocusOnYourself",
			req: &relation.RelationFollowActionReq{
				AccessToken: "1111",
				ToUserID:    "1111",
			},
			errorIsExist:                true,
			expectedError:               "不能关注自己",
			mockAccessTokenResultReturn: 1111,
		},
		{
			name: "QueryUserFail",
			req: &relation.RelationFollowActionReq{
				AccessToken: "1111",
				ToUserID:    "1111",
			},
			errorIsExist:             true,
			expectedError:            errno.DatabaseCallErrorMsg,
			mockQueryUserErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "UserIsNotExist",
			req: &relation.RelationFollowActionReq{
				AccessToken: "1111",
				ToUserID:    "1111",
			},
			errorIsExist:  true,
			expectedError: "用户不存在",
		},
		{
			name: "AccessTokenFail",
			req: &relation.RelationFollowActionReq{
				AccessToken: "1111",
				ToUserID:    "aaa",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "ActionTypeParamInvalid",
			req: &relation.RelationFollowActionReq{
				ToUserID:    "111",
				AccessToken: "1111",
				ActionType:  111,
			},
			errorIsExist:              true,
			expectedError:             "操作类型不合法",
			mockQueryUserResultReturn: true,
		},
		{
			name: "QueryFollowFail",
			req: &relation.RelationFollowActionReq{
				AccessToken: "1111",
				ToUserID:    "1111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryFollowErrorReturn: errno.DatabaseCallError,
			mockQueryUserResultReturn:  true,
		},
		{
			name: "FollowIsExist",
			req: &relation.RelationFollowActionReq{
				AccessToken: "1111",
				ToUserID:    "1111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist:                true,
			expectedError:               "已关注",
			mockQueryUserResultReturn:   true,
			mockQueryFollowResultReturn: true,
		},
		{
			name: "InsertFollowFail",
			req: &relation.RelationFollowActionReq{
				AccessToken: "1111",
				ToUserID:    "1111",
				ActionType:  common.ActionTypeOn,
			},
			errorIsExist:                true,
			expectedError:               errno.DatabaseCallErrorMsg,
			mockInsertFollowErrorReturn: errno.DatabaseCallError,
			mockQueryUserResultReturn:   true,
		},
		{
			name: "Success",
			req: &relation.RelationFollowActionReq{
				ToUserID:    "111",
				AccessToken: "1111",
				ActionType:  common.ActionTypeOff,
			},
			errorIsExist:              false,
			mockQueryUserResultReturn: true,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryUserExistByID).Return(tc.mockQueryUserResultReturn, tc.mockQueryUserErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(tc.mockAccessTokenResultReturn, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(exquery.QueryFollowExistByFollowerIDAndFollowedID).Return(tc.mockQueryFollowResultReturn, tc.mockQueryFollowErrorReturn).Build()
			mockey.Mock(exquery.InsertFollowRecord).Return(tc.mockInsertFollowErrorReturn).Build()
			mockey.Mock(exquery.DeleteFollowRecord).Return(tc.mockDeleteFollowErrorReturn).Build()

			err := relationService.NewFollowActionEvent(tc.req)

			if tc.errorIsExist {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewFollowListEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *relation.RelationFollowListReq
		errorIsExist               bool
		expectedError              string
		mockQueryFollowErrorReturn error
		mockQueryFollowReturn      []*model.Follow
		mockQueryFollowCountReturn int64
		mockQueryUserErrorReturn   error
		mockQueryUserReturn        *model.User
		mockAccessTokenErrorReturn error
		mockConvertErrorReturn     error
		mockConvertResultReturn    *[]*base.User
		expectedResult             *relation.RelationFollowListRespData
	}

	follow := make([]*model.Follow, 0)
	follow = append(follow, &model.Follow{
		FollowedID: 0,
		FollowerID: 0,
		CreatedAt:  0,
		DeletedAt:  0,
	})
	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: &relation.RelationFollowListReq{
				UserID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "用户ID错误",
		},
		{
			name: "QueryFollowFail",
			req: &relation.RelationFollowListReq{
				UserID: "111",
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryFollowErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "QueryUserFail",
			req: &relation.RelationFollowListReq{
				UserID: "111",
			},
			errorIsExist:             true,
			expectedError:            errno.DatabaseCallErrorMsg,
			mockQueryUserErrorReturn: errno.DatabaseCallError,
			mockQueryFollowReturn:    follow,
		},
		{
			name: "AccessTokenFail",
			req: &relation.RelationFollowListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
			mockQueryFollowReturn:      follow,
			mockQueryUserReturn:        new(model.User),
		},
		{
			name: "ConvertFail",
			req: &relation.RelationFollowListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &relation.RelationFollowListReq{
				UserID: "111",
			},
			errorIsExist: false,
			expectedResult: &relation.RelationFollowListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: new([]*base.User),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryFollowingByUserIdPaged).Return(tc.mockQueryFollowReturn, tc.mockQueryFollowCountReturn, tc.mockQueryFollowErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.UserListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()
			mockey.Mock(exquery.QueryUserByID).Return(tc.mockQueryUserReturn, tc.mockQueryUserErrorReturn).Build()

			result, err := relationService.NewFollowListEvent(tc.req)

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

func TestNewFollowerListEvent(t *testing.T) {
	type testCase struct {
		name                         string
		req                          *relation.RelationFollowerListReq
		errorIsExist                 bool
		expectedError                string
		mockQueryFollowerErrorReturn error
		mockQueryFollowerReturn      []*model.Follow
		mockQueryFollowerCountReturn int64
		mockQueryUserErrorReturn     error
		mockQueryUserReturn          *model.User
		mockAccessTokenErrorReturn   error
		mockConvertErrorReturn       error
		mockConvertResultReturn      *[]*base.User
		expectedResult               *relation.RelationFollowerListRespData
	}

	follower := make([]*model.Follow, 0)
	follower = append(follower, &model.Follow{
		FollowedID: 0,
		FollowerID: 0,
		CreatedAt:  0,
		DeletedAt:  0,
	})
	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: &relation.RelationFollowerListReq{
				UserID: "aaa",
			},
			errorIsExist:  true,
			expectedError: "用户ID错误",
		},
		{
			name: "QueryFollowerFail",
			req: &relation.RelationFollowerListReq{
				UserID: "111",
			},
			errorIsExist:                 true,
			expectedError:                errno.DatabaseCallErrorMsg,
			mockQueryFollowerErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "QueryUserFail",
			req: &relation.RelationFollowerListReq{
				UserID: "111",
			},
			errorIsExist:             true,
			expectedError:            errno.DatabaseCallErrorMsg,
			mockQueryUserErrorReturn: errno.DatabaseCallError,
			mockQueryFollowerReturn:  follower,
		},
		{
			name: "AccessTokenFail",
			req: &relation.RelationFollowerListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
			mockQueryFollowerReturn:    follower,
			mockQueryUserReturn:        new(model.User),
		},
		{
			name: "ConvertFail",
			req: &relation.RelationFollowerListReq{
				UserID:      "111",
				AccessToken: new(string),
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: &relation.RelationFollowerListReq{
				UserID: "111",
			},
			errorIsExist: false,
			expectedResult: &relation.RelationFollowerListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: new([]*base.User),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryFollowerByUserIdPaged).Return(tc.mockQueryFollowerReturn, tc.mockQueryFollowerCountReturn, tc.mockQueryFollowerErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ExtractPayloadFromToken).Return("111", tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.UserListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()
			mockey.Mock(exquery.QueryUserByID).Return(tc.mockQueryUserReturn, tc.mockQueryUserErrorReturn).Build()

			result, err := relationService.NewFollowerListEvent(tc.req)

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

func TestNewFriendListEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        *relation.RelationFriendListReq
		errorIsExist               bool
		expectedError              string
		mockQueryFriendErrorReturn error
		mockQueryFriendReturn      []int64
		mockQueryFriendCountReturn int64
		mockQueryUserErrorReturn   error
		mockQueryUserReturn        *model.User
		mockConvertErrorReturn     error
		mockConvertResultReturn    *[]*base.User
		mockAccessTokenErrorReturn error
		expectedResult             *relation.RelationFriendListRespData
	}

	friend := make([]int64, 0)
	friend = append(friend, 1)
	testCases := []testCase{
		{
			name: "AccessTokenFail",
			req: &relation.RelationFriendListReq{
				AccessToken: "1111",
			},
			errorIsExist:               true,
			expectedError:              errno.AccessTokenInvalidErrorMsg,
			mockAccessTokenErrorReturn: errno.AccessTokenInvalid,
		},
		{
			name: "QueryFollowFail",
			req: &relation.RelationFriendListReq{
				AccessToken: "1111",
			},
			errorIsExist:               true,
			expectedError:              errno.DatabaseCallErrorMsg,
			mockQueryFriendErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "QueryUserFail",
			req: &relation.RelationFriendListReq{
				AccessToken: "1111",
			},
			errorIsExist:             true,
			expectedError:            errno.DatabaseCallErrorMsg,
			mockQueryUserErrorReturn: errno.DatabaseCallError,
			mockQueryFriendReturn:    friend,
		},
		{
			name: "ConvertFail",
			req: &relation.RelationFriendListReq{
				AccessToken: "1111",
			},
			errorIsExist:           true,
			expectedError:          errno.DatabaseCallErrorMsg,
			mockConvertErrorReturn: errno.DatabaseCallError,
			mockQueryFriendReturn:  friend,
			mockQueryUserReturn:    new(model.User),
		},
		{
			name: "Success",
			req: &relation.RelationFriendListReq{
				AccessToken: "1111",
			},
			errorIsExist: false,
			expectedResult: &relation.RelationFriendListRespData{
				Items:    nil,
				IsEnd:    true,
				PageSize: 1,
				PageNum:  0,
				Total:    0,
			},
			mockConvertResultReturn: new([]*base.User),
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryFriendByUserIDPaged).Return(tc.mockQueryFriendReturn, tc.mockQueryFriendCountReturn, tc.mockQueryFriendErrorReturn).Build()
			mockey.Mock((*jwt.JWTMiddleware).ConvertJWTPayloadToInt64).Return(111, tc.mockAccessTokenErrorReturn).Build()
			mockey.Mock(model_converter.UserListDal2Resp).Return(tc.mockConvertResultReturn, tc.mockConvertErrorReturn).Build()
			mockey.Mock(exquery.QueryUserByID).Return(tc.mockQueryUserReturn, tc.mockQueryUserErrorReturn).Build()

			result, err := relationService.NewFriendListEvent(tc.req)

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
