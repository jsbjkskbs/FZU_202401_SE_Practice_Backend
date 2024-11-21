package service

import (
	"context"
	"fmt"
	"strconv"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/activity"
	"sfw/biz/mw/jwt"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/utils/generator"

	"github.com/cloudwego/hertz/pkg/app"
)

type ActivityService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewActivityService(ctx context.Context, c *app.RequestContext) *ActivityService {
	return &ActivityService{
		ctx: ctx,
		c:   c,
	}
}

func (service *ActivityService) NewFeedEvent(req *activity.ActivityFeedReq) (*activity.ActivityFeedRespData, error) {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	activities, count, err := exquery.QueryActivityByFollowedIdPaged(uid, int(req.PageNum), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	fromUser := fmt.Sprint(uid)
	items, err := model_converter.ActivityListDal2Resp(&activities, &fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &activity.ActivityFeedRespData{
		Items:    *items,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *ActivityService) NewPublishEvent(req *activity.ActivityPublishReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return err
	}
	var refCount int64
	var activity model.Activity
	if req.RefActivity != nil {
		activity.RefActivityID, err = strconv.ParseInt(*req.RefActivity, 10, 64)
		if err != nil {
			return errno.ParamInvalid.WithMessage("无效的动态ID")
		}
		refCount++
	}
	if req.RefVideo != nil {
		activity.RefVideoID, err = strconv.ParseInt(*req.RefVideo, 10, 64)
		if err != nil {
			return errno.ParamInvalid.WithMessage("无效的视频ID")
		}
		refCount++
	}
	if refCount >= 2 {
		return errno.ParamInvalid.WithMessage("只能引用一个视频或一个动态")
	}

	if len(req.Image) > 0 {
		for _, image := range req.Image {
			imageId, err := strconv.ParseInt(image, 10, 64)
			if err != nil {
				return errno.ParamInvalid.WithMessage("无效的图片ID: " + image).WithInnerError(err)
			}
			exist, err := exquery.QueryImageExistById(imageId)
			if err != nil {
				return errno.DatabaseCallError.WithInnerError(err)
			}
			if !exist {
				return errno.ParamInvalid.WithMessage("图片不存在: " + image)
			}
		}
	}

	activity.ID = generator.ActivityIDGenerator.Generate()
	activity.UserID = uid
	activity.Content = req.Content

	err = exquery.InsertActivity(&activity)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}

	if len(req.Image) > 0 {
		imgs := make([]*model.ActivityImage, 0, len(req.Image))
		for _, image := range req.Image {
			imageId, _ := strconv.ParseInt(image, 10, 64)
			imgs = append(imgs, &model.ActivityImage{
				ActivityID: activity.ID,
				ImageID:    imageId,
			})
		}
		err = exquery.InsertActivityImage(imgs...)
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
	}

	return nil
}

func (service *ActivityService) NewListEvent(req *activity.ActivityListReq) (*activity.ActivityListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("无效的用户ID")
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	activities, count, err := exquery.QueryActivityByUserIdPaged(uid, int(req.PageNum), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	var fromUser *string
	if req.AccessToken != nil {
		uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(*req.AccessToken)
		if err != nil {
			return nil, errno.AccessTokenInvalid
		}
		fromUser = &uid
	}

	items, err := model_converter.ActivityListDal2Resp(&activities, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &activity.ActivityListRespData{
		Items:    *items,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}
