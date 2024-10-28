package service

import (
	"context"
	"encoding/json"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/oss"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/redis"
	"sfw/pkg/errno"
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

type _CallbackReq struct {
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
	req := _CallbackReq{}
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

func (service *OssService) NewCallbackVideoEvent(_ *oss.OssCallbackVideoReq) error {
	body := service.c.Request.Body()
	req := _CallbackReq{}
	json.Unmarshal(body, &req)
	req.Key = strings.ReplaceAll(req.Key, "%2F", "/")

	videoId, err := strconv.ParseInt(req.Oid, 10, 64)
	if err != nil {
		return err
	}

	stat, err := redis.VideoUploadInfoGet(req.Oid)
	if err != nil {
		return errno.InternalServerError
	}

	userId, err := strconv.ParseInt(stat["user_id"], 10, 64)
	if err != nil {
		return err
	}

	c := dal.Executor.Category
	category, err := c.WithContext(context.Background()).Where(c.CategoryName.Eq(stat["category"])).First()
	if err != nil {
		return errno.DatabaseCallError
	}

	v := dal.Executor.Video
	err = v.WithContext(context.Background()).Create(&model.Video{
		ID:          videoId,
		UserID:      userId,
		VideoURL:    req.Key,
		CoverURL:    strings.Replace(req.Key, "video.mp4", "cover.jpg", strings.LastIndex(req.Key, "/")),
		Title:       stat["title"],
		Description: stat["description"],
		CategoryID:  category.ID,
		VisitCount:  0,
		Status:      "review",
	})
	if err != nil {
		return errno.DatabaseCallError
	}

	labels := strings.Split(stat["labels"], "\t")
	l := dal.Executor.VideoLabel
	vLabels := make([]*model.VideoLabel, 0, len(labels))
	for _, v := range labels {
		vLabels = append(vLabels, &model.VideoLabel{
			VideoID:   videoId,
			LabelName: v,
		})
	}
	err = l.WithContext(context.Background()).Create(vLabels...)
	if err != nil {
		return errno.DatabaseCallError
	}

	go redis.VideoUploadInfoDel(req.Oid)
	go gorse.InsertVideo(req.Oid, stat["category"], labels)
	return nil
}
