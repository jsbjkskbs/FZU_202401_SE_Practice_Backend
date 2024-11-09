package service

import (
	"encoding/json"
	"github.com/bytedance/mockey"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/stretchr/testify/assert"
	"sfw/biz/dal/exquery"
	"sfw/biz/model/api/oss"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/redis"
	"sfw/pkg/errno"
	"sfw/pkg/utils/checker"
	"sfw/pkg/utils/generator"
	"testing"
	"time"
)

var ossService = NewOssService(nil, new(app.RequestContext))

func TestNewCallbackAvatarEvent(t *testing.T) {
	type testCase struct {
		name                  string
		req                   _CallbackReq
		errorIsExist          bool
		mockUpdateErrorReturn error
	}

	testCases := []testCase{
		{
			name: "ParamInvalid",
			req: _CallbackReq{
				Oid: "aaa",
			},
			errorIsExist: true,
		},
		{
			name: "Success",
			req: _CallbackReq{
				Oid: "111",
			},
			errorIsExist: false,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.UpdateUserWithId).Return(tc.mockUpdateErrorReturn).Build()
			body, _ := json.Marshal(tc.req)
			mockey.Mock((*protocol.Request).Body).Return(body).Build()

			err := ossService.NewCallbackAvatarEvent(new(oss.OssCallbackAvatarReq))

			if tc.errorIsExist {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewCallbackVideoEvent(t *testing.T) {
	type testCase struct {
		name                       string
		req                        _CallbackReq
		errorIsExist               bool
		mockUploadErrorReturn      error
		mockUploadResultReturn     map[string]string
		mockInsertVideoErrorReturn error
		mockInsertLabelErrorReturn error
	}

	stat := make(map[string]string)
	stat["user_id"] = "111"
	stat["category"] = "运动"
	stat["labels"] = "1\t2\t3"
	wrongStat := make(map[string]string)
	wrongStat["user_id"] = "111"
	wrongStat["category"] = "111"
	testCases := []testCase{
		{
			name: "OidParamInvalid",
			req: _CallbackReq{
				Oid: "aaa",
			},
			errorIsExist: true,
		},
		{
			name: "UploadFail",
			req: _CallbackReq{
				Oid: "111",
			},
			errorIsExist:          true,
			mockUploadErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "StatInvalid",
			req: _CallbackReq{
				Oid: "111",
			},
			errorIsExist:           true,
			mockUploadResultReturn: make(map[string]string),
		},
		{
			name: "CategoryNotFound",
			req: _CallbackReq{
				Oid: "111",
			},
			errorIsExist:           true,
			mockUploadResultReturn: wrongStat,
		},
		{
			name: "InsertVideoFail",
			req: _CallbackReq{
				Oid: "111",
			},
			errorIsExist:               true,
			mockUploadResultReturn:     stat,
			mockInsertVideoErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "InsertLabelFail",
			req: _CallbackReq{
				Oid: "111",
			},
			errorIsExist:               true,
			mockUploadResultReturn:     stat,
			mockInsertLabelErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: _CallbackReq{
				Oid: "111",
			},
			errorIsExist:           false,
			mockUploadResultReturn: stat,
		},
	}

	generator.VideoIDGenerator, _ = generator.NewSnowflake(5)
	checker.CategoryMap["运动"] = 1
	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(redis.VideoUploadInfoGet).Return(tc.mockUploadResultReturn, tc.mockUploadErrorReturn).Build()
			mockey.Mock(exquery.InsertVideo).Return(tc.mockInsertVideoErrorReturn).Build()
			mockey.Mock(exquery.InsertVideoLabel).Return(tc.mockInsertLabelErrorReturn).Build()
			mockey.Mock(redis.VideoUploadInfoDel).Return(nil).Build()
			mockey.Mock(gorse.InsertVideo).Return(nil).Build()
			body, _ := json.Marshal(tc.req)
			mockey.Mock((*protocol.Request).Body).Return(body).Build()

			err := ossService.NewCallbackVideoEvent(new(oss.OssCallbackVideoReq))

			if tc.errorIsExist {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				time.Sleep(1 * time.Second)
			}
		})
	}
}

func TestNewCallbackImageEvent(t *testing.T) {
	type testCase struct {
		name                 string
		req                  _ImageCallbackReq
		errorIsExist         bool
		mockQueryErrorReturn error
	}

	testCases := []testCase{
		{
			name: "OidParamInvalid",
			req: _ImageCallbackReq{
				Oid: "aaa",
			},
			errorIsExist: true,
		},
		{
			name: "UserIdParamInvalid",
			req: _ImageCallbackReq{
				Oid:    "111",
				UserId: "aaa",
			},
			errorIsExist: true,
		},
		{
			name: "QueryFail",
			req: _ImageCallbackReq{
				Oid:    "111",
				UserId: "111",
			},
			errorIsExist:         true,
			mockQueryErrorReturn: errno.DatabaseCallError,
		},
		{
			name: "Success",
			req: _ImageCallbackReq{
				Oid:    "111",
				UserId: "111",
			},
			errorIsExist: false,
		},
	}

	defer mockey.UnPatchAll()

	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			t.Logf("%s  :  %s", t.Name(), tc.name)

			mockey.Mock(exquery.QueryImageExistById).Return(true, tc.mockQueryErrorReturn).Build()
			mockey.Mock(exquery.UpdateImageWithId).Return(nil).Build()
			body, _ := json.Marshal(tc.req)
			mockey.Mock((*protocol.Request).Body).Return(body).Build()

			err := ossService.NewCallbackImageEvent(new(oss.OssCallbackImageReq))

			if tc.errorIsExist {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				time.Sleep(1 * time.Second)
			}
		})
	}
}
