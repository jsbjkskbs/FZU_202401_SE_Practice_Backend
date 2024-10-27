package service

import (
	"context"
	"encoding/json"
	"sfw/biz/dal"
	"sfw/biz/model/api/oss"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

type OssService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewOssService(ctx context.Context, c *app.RequestContext) *OssService {
	return &OssService{
		ctx: ctx,
		c:   c,
	}
}

type _CallbackAvatarReq struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Fsize  int64  `json:"fsize"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
	Otype  string `json:"otype"`
	Oid    string `json:"oid"`
}

func (service *OssService) NewCallbackAvatarEvent(_ *oss.OssCallbackAvatarReq) error {
	body := service.c.Request.Body()
	req := _CallbackAvatarReq{}
	json.Unmarshal(body, &req)
	req.Key = strings.ReplaceAll(req.Key, "%2F", "/")

	id, err := strconv.ParseInt(req.Oid, 10, 64)
	if err != nil {
		return err
	}
	u := dal.Executor.User
	_, err = u.WithContext(context.Background()).Where(u.ID.Eq(id)).Update(u.AvatarURL, req.Key)
	return err
}
