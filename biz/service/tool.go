package service

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/tool"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/pkg/errno"
	"sfw/pkg/oss"
	"sfw/pkg/utils/generator"
	"strconv"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
)

type ToolService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewToolService(ctx context.Context, c *app.RequestContext) *ToolService {
	return &ToolService{
		ctx: ctx,
		c:   c,
	}
}

func (service *ToolService) NewDeleteVideoEvent(req *tool.ToolDeleteVideoReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	v := dal.Executor.Video
	exist, err := v.WithContext(context.Background()).Where(v.UserID.Eq(uid), v.ID.Eq(videoId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("视频不存在或者不属于你")
	}

	_, err = v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		err := redis.DeleteVideo(req.VideoID)
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		vc := dal.Executor.VideoComment
		_, err := vc.WithContext(context.Background()).Where(vc.VideoID.Eq(videoId)).Delete()
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return errno.InternalServerError.WithInnerError(err)
	default:
		return nil
	}
}

func (service *ToolService) NewDeleteActivityEvent(req *tool.ToolDeleteActivityReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	activityId, err := strconv.ParseInt(req.ActivityID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	a := dal.Executor.Activity
	exist, err := a.WithContext(context.Background()).Where(a.UserID.Eq(uid), a.ID.Eq(activityId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("动态不存在或者不属于你")
	}

	_, err = a.WithContext(context.Background()).Where(a.ID.Eq(activityId)).Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		err := redis.DeleteActivity(req.ActivityID)
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		ac := dal.Executor.ActivityComment
		_, err := ac.WithContext(context.Background()).Where(ac.ActivityID.Eq(activityId)).Delete()
		if err != nil {
			errs <- err
		}
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return errno.InternalServerError.WithInnerError(err)
	default:
		return nil
	}
}

func (service *ToolService) newDeleteVideoCommentEvent(req *tool.ToolDeleteCommentReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	commentId, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	vc := dal.Executor.VideoComment
	exist, err := vc.WithContext(context.Background()).Where(vc.UserID.Eq(uid), vc.ID.Eq(commentId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("不能删除别人的评论")
	}

	list := []model.VideoComment{}
	err = vc.WithContext(context.Background()).
		Where(vc.ID.Eq(commentId)).
		Or(vc.RootID.Eq(commentId)).
		Select(vc.ID, vc.VideoID).
		Scan(&list)
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	_, err = vc.WithContext(context.Background()).
		Where(vc.ID.Eq(commentId)).
		Or(vc.RootID.Eq(commentId)).
		Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	go func() {
		for _, v := range list {
			redis.DeleteVideoComment(fmt.Sprint(v.VideoID), fmt.Sprint(v.ID))
		}
	}()

	return nil
}

func (service *ToolService) newDeleteActivityCommentEvent(req *tool.ToolDeleteCommentReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	commentId, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	ac := dal.Executor.ActivityComment
	exist, err := ac.WithContext(context.Background()).Where(ac.UserID.Eq(uid), ac.ID.Eq(commentId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("不能删除别人的评论")
	}

	list := []model.ActivityComment{}
	err = ac.WithContext(context.Background()).
		Where(ac.ID.Eq(commentId)).
		Or(ac.RootID.Eq(commentId)).
		Select(ac.ID, ac.ActivityID).
		Scan(&list)
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	_, err = ac.WithContext(context.Background()).
		Where(ac.ID.Eq(commentId)).
		Or(ac.RootID.Eq(commentId)).
		Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	go func() {
		for _, v := range list {
			redis.DeleteActivityComment(fmt.Sprint(v.ActivityID), fmt.Sprint(v.ID))
		}
	}()

	return nil
}

func (tool *ToolService) NewDeleteCommentEvent(req *tool.ToolDeleteCommentReq) error {
	switch req.CommentType {
	case "video":
		return tool.newDeleteVideoCommentEvent(req)
	case "activity":
		return tool.newDeleteActivityCommentEvent(req)
	}
	return errno.ParamInvalid.WithMessage("评论类型错误")
}

func (service *ToolService) NewAdminDeleteVideoEvent(req *tool.AdminToolDeleteVideoReq) error {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	v := dal.Executor.Video
	exist, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("视频不存在")
	}

	_, err = v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		err := redis.DeleteVideo(req.VideoID)
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		vc := dal.Executor.VideoComment
		_, err := vc.WithContext(context.Background()).Where(vc.VideoID.Eq(videoId)).Delete()
		if err != nil {
			errs <- err
		}
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return errno.InternalServerError.WithInnerError(err)
	default:
		return nil
	}
}

func (service *ToolService) NewAdminDeleteActivityEvent(req *tool.AdminToolDeleteActivityReq) error {
	activityId, err := strconv.ParseInt(req.ActivityID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	a := dal.Executor.Activity
	exist, err := a.WithContext(context.Background()).Where(a.ID.Eq(activityId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("动态不存在")
	}

	_, err = a.WithContext(context.Background()).Where(a.ID.Eq(activityId)).Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	errs := make(chan error, 2)
	go func() {
		err := redis.DeleteActivity(req.ActivityID)
		if err != nil {
			errs <- err
		}
		wg.Done()
	}()
	go func() {
		ac := dal.Executor.ActivityComment
		_, err := ac.WithContext(context.Background()).Where(ac.ActivityID.Eq(activityId)).Delete()
		if err != nil {
			errs <- err
		}
	}()
	wg.Wait()
	select {
	case err := <-errs:
		return errno.InternalServerError.WithInnerError(err)
	default:
		return nil
	}
}

func (service *ToolService) newAdminDeleteVideoCommentEvent(req *tool.AdminToolDeleteCommentReq) error {
	commentId, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	vc := dal.Executor.VideoComment
	exist, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(commentId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("评论不存在")
	}

	list := []model.VideoComment{}
	err = vc.WithContext(context.Background()).
		Where(vc.ID.Eq(commentId)).
		Or(vc.RootID.Eq(commentId)).
		Select(vc.ID, vc.VideoID).
		Scan(&list)
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	_, err = vc.WithContext(context.Background()).
		Where(vc.ID.Eq(commentId)).
		Or(vc.RootID.Eq(commentId)).
		Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	go func() {
		for _, v := range list {
			redis.DeleteVideoComment(fmt.Sprint(v.VideoID), fmt.Sprint(v.ID))
		}
	}()

	return nil
}

func (service *ToolService) newAdminDeleteActivityCommentEvent(req *tool.AdminToolDeleteCommentReq) error {
	commentId, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithInnerError(err)
	}

	ac := dal.Executor.ActivityComment
	exist, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(commentId)).Count()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("评论不存在")
	}

	list := []model.ActivityComment{}
	err = ac.WithContext(context.Background()).
		Where(ac.ID.Eq(commentId)).
		Or(ac.RootID.Eq(commentId)).
		Select(ac.ID, ac.ActivityID).
		Scan(&list)
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	_, err = ac.WithContext(context.Background()).
		Where(ac.ID.Eq(commentId)).
		Or(ac.RootID.Eq(commentId)).
		Delete()
	if err != nil {
		return errno.InternalServerError.WithInnerError(err)
	}

	go func() {
		for _, v := range list {
			redis.DeleteActivityComment(fmt.Sprint(v.ActivityID), fmt.Sprint(v.ID))
		}
	}()

	return nil
}

func (service *ToolService) NewAdminDeleteCommentEvent(req *tool.AdminToolDeleteCommentReq) error {
	switch req.CommentType {
	case "video":
		return service.newAdminDeleteVideoCommentEvent(req)
	case "activity":
		return service.newAdminDeleteActivityCommentEvent(req)
	}
	return errno.ParamInvalid.WithMessage("评论类型错误")
}

func (service *ToolService) NewUploadImageEvent(req *tool.ToolUploadImageReq) (*tool.ToolUploadImageRespData, error) {
	uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}
	imageId := generator.ImageIDGenerator.Generate()
	uptoken, key, err := oss.UploadImage(fmt.Sprint(imageId), imageId, uid)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}
	return &tool.ToolUploadImageRespData{
		UploadURL: oss.UploadUrl,
		Uptoken:   uptoken,
		UploadKey: key,
		ImageID:   fmt.Sprint(imageId),
	}, nil
}

func (service *ToolService) NewGetImageEvent(req *tool.ToolGetImageReq) (*tool.ToolGetImageRespData, error) {
	imageId, err := strconv.ParseInt(req.ImageID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithInnerError(err)
	}
	img := dal.Executor.Image
	image, _ := img.WithContext(context.Background()).Where(img.ID.Eq(imageId)).First()
	if image == nil {
		return nil, errno.CustomError.WithMessage("图片不存在")
	}
	return &tool.ToolGetImageRespData{
		URL: oss.Key2Url(image.ImageURL),
	}, nil
}

func (service *ToolService) NewTokenRefreshEvent(req *tool.ToolTokenRefreshReq) (*tool.ToolTokenRefreshRespData, error) {
	payload, expire, err := jwt.RefreshTokenJwtMiddleware.GetBasicDataFromToken(req.RefreshToken)
	if err != nil {
		return nil, errno.RefreshTokenInvalid
	}
	expiredAt, err := redis.TokenExpireTimeGet(fmt.Sprint(payload))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if expiredAt >= expire.Unix() {
		return nil, errno.RefreshTokenForbidden
	}
	token := jwt.AccessTokenJwtMiddleware.GenerateToken(fmt.Sprint(payload))
	return &tool.ToolTokenRefreshRespData{
		AccessToken: token,
	}, nil
}
