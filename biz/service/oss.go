package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/oss"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/redis"
	"sfw/biz/service/common"
	"sfw/pkg/errno"
	"sfw/pkg/utils/checker"

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
	err = exquery.UpdateUserWithId(&model.User{
		ID:        id,
		AvatarURL: req.Key,
	})
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

	categoryId, ok := checker.CategoryMap[stat["category"]]
	if !ok {
		return errors.New("category not found")
	}

	err = exquery.InsertVideo(&model.Video{
		ID:          videoId,
		UserID:      userId,
		VideoURL:    req.Key,
		CoverURL:    strings.Replace(req.Key, "video.mp4", "cover.jpg", strings.LastIndex(req.Key, "/")),
		Title:       stat["title"],
		Description: stat["description"],
		CategoryID:  categoryId,
		VisitCount:  0,
		Status:      common.VideoStatusReview,
	})
	if err != nil {
		return err
	}

	labels := strings.Split(stat["labels"], "\t")
	vLabels := make([]*model.VideoLabel, 0, len(labels))
	for _, v := range labels {
		vLabels = append(vLabels, &model.VideoLabel{
			VideoID:   videoId,
			LabelName: v,
		})
	}
	err = exquery.InsertVideoLabel(vLabels...)
	if err != nil {
		return err
	}

	go redis.VideoUploadInfoDel(req.Oid)
	go gorse.InsertVideo(req.Oid, stat["category"], labels)
	return nil
}

type _ImageCallbackReq struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Fsize  int64  `json:"fsize"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
	Otype  string `json:"otype"`
	Oid    string `json:"oid"`
	UserId string `json:"user_id"`
}

func (service *OssService) NewCallbackImageEvent(_ *oss.OssCallbackImageReq) error {
	body := service.c.Request.Body()
	req := _ImageCallbackReq{}
	json.Unmarshal(body, &req)
	req.Key = strings.ReplaceAll(req.Key, "%2F", "/")

	imageId, err := strconv.ParseInt(req.Oid, 10, 64)
	if err != nil {
		return err
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return err
	}
	exist, err := exquery.QueryImageExistById(imageId)
	if err != nil {
		return errno.DatabaseCallError
	}
	if exist {
		exquery.UpdateImageWithId(&model.Image{
			ID:       imageId,
			ImageURL: req.Key,
			UserID:   userId,
		})
	}
	return nil
}
