package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/video"
	"sfw/biz/model/base"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/oss"
	"sfw/pkg/synchronizer"
	"sfw/pkg/utils/checker"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/scheduler"

	"github.com/cloudwego/hertz/pkg/app"
)

type VideoService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewVideoService(ctx context.Context, c *app.RequestContext) *VideoService {
	return &VideoService{
		ctx: ctx,
		c:   c,
	}
}

func (service *VideoService) NewPublishEvent(req *video.VideoPublishReq) (*video.VideoPublishRespData, error) {
	id, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid.WithInnerError(err)
	}

	err = checker.CheckVideoPublish(req.Title, req.Description, req.Category, req.Labels)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("视频信息不合法")
	}

	videoId := generator.VideoIDGenerator.Generate()
	kv := map[string]interface{}{
		"user_id":     fmt.Sprint(id),
		"title":       req.Title,
		"description": req.Description,
		"category":    req.Category,
		"labels":      strings.Join(req.Labels, "\t"),
	}

	err = redis.VideoUploadInfoStore(fmt.Sprint(videoId), kv, oss.VideoUploadTokenDeadline)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	uptoken, uploadKey, err := oss.UploadVideo(fmt.Sprint(videoId), videoId)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	return &video.VideoPublishRespData{
		UploadURL: oss.UploadUrl,
		UploadKey: uploadKey,
		Uptoken:   uptoken,
	}, nil
}

func (service *VideoService) NewCoverUploadEvent(req *video.VideoCoverUploadReq) (*video.VideoCoverUploadRespData, error) {
	userId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("视频ID错误")
	}

	exist, err := exquery.QueryVideoExistByIdAndUserId(videoId, userId)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return nil, errno.ResourceNotFound.WithMessage("视频不存在")
	}

	uptoken, key, err := oss.UploadVideoCover(req.VideoID, videoId)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}
	return &video.VideoCoverUploadRespData{
		UploadURL: oss.UploadUrl,
		UploadKey: key,
		Uptoken:   uptoken,
	}, nil
}

func (service *VideoService) NewFeedEvent(req *video.VideoFeedReq) (*video.VideoFeedRespData, error) {
	var (
		vids = []string{}
		err  error
	)

	req.Offset, req.N = common.CorrectPageNumAndPageSize(req.Offset, req.N)

	if req.Category != nil {
		vids, err = gorse.GetRecommendWithCategory("", *req.Category, int(req.N), int(req.Offset))
	} else {
		vids, err = gorse.GetRecommendWithCategory("", "*", int(req.N), int(req.Offset))
	}

	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	videos := make([]*model.Video, 0)
	for _, vid := range vids {
		videoId, err := strconv.ParseInt(vid, 10, 64)
		if err != nil {
			return nil, errno.InternalServerError.WithInnerError(err)
		}
		video, err := exquery.QueryVideoById(videoId)
		if err != nil {
			return nil, errno.DatabaseCallError.WithInnerError(err)
		}
		if video == nil {
			return nil, errno.ResourceNotFound.WithMessage("视频不存在")
		}
		videos = append(videos, video)
	}
	data, err := model_converter.VideoListDal2Resp(&videos, nil)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}
	return &video.VideoFeedRespData{
		Items:  data,
		Offset: req.Offset,
		N:      req.N,
	}, nil
}

func (service *VideoService) NewCustomFeedEvent(req *video.VideoCustomFeedReq) (*video.VideoCustomFeedRespData, error) {
	userId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.Offset, req.N = common.CorrectPageNumAndPageSize(req.Offset, req.N)

	vids := []string{}
	if req.Category != nil {
		vids, err = gorse.GetRecommendWithCategory(fmt.Sprint(userId), *req.Category, int(req.N), int(req.Offset))
	} else {
		vids, err = gorse.GetRecommendWithCategory(fmt.Sprint(userId), "*", int(req.N), int(req.Offset))
	}

	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	videos := make([]*model.Video, 0)
	for _, vid := range vids {
		videoId, err := strconv.ParseInt(vid, 10, 64)
		if err != nil {
			return nil, errno.InternalServerError.WithInnerError(err)
		}
		video, err := exquery.QueryVideoById(videoId)
		if err != nil {
			return nil, errno.DatabaseCallError
		}
		if video == nil {
			return nil, errno.ResourceNotFound.WithMessage("视频不存在")
		}
		videos = append(videos, video)
	}

	fromUser := fmt.Sprint(userId)

	data, err := model_converter.VideoListDal2Resp(&videos, &fromUser)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}
	return &video.VideoCustomFeedRespData{
		Items:  data,
		Offset: req.Offset,
		N:      req.N,
	}, nil
}

func (service *VideoService) NewNeighbourFeedEvent(req *video.VideoNeighbourFeedReq) (*video.VideoNeighbourFeedRespData, error) {
	req.Offset, req.N = common.CorrectPageNumAndPageSize(req.Offset, req.N)

	userId := ""
	if req.AccessToken != nil {
		uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(*req.AccessToken)
		if err != nil {
			return nil, errno.AccessTokenInvalid
		}
		userId = uid
	}

	vids, err := gorse.GetItemNeighbors(userId, req.VideoID, int(req.N), int(req.Offset))
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	videos := make([]*model.Video, 0)
	for _, vid := range vids {
		videoId, err := strconv.ParseInt(vid, 10, 64)
		if err != nil {
			return nil, errno.InternalServerError.WithInnerError(err)
		}
		video, err := exquery.QueryVideoById(videoId)
		if err != nil {
			return nil, errno.DatabaseCallError.WithInnerError(err)
		}
		if video == nil {
			return nil, errno.ResourceNotFound.WithMessage("视频不存在")
		}
		videos = append(videos, video)
	}

	data, err := model_converter.VideoListDal2Resp(&videos, nil)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}
	return &video.VideoNeighbourFeedRespData{
		Items:  data,
		Offset: req.Offset,
		N:      req.N,
	}, nil
}

func (service *VideoService) NewCategoriesEvent(req *video.VideoCategoriesReq) ([]string, error) {
	return checker.Categories, nil
}

func (service *VideoService) NewInfoEvent(req *video.VideoInfoReq) (*base.Video, error) {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("视频ID错误")
	}

	video, err := exquery.QueryVideoById(videoId)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if video == nil {
		return nil, errno.ResourceNotFound.WithMessage("视频不存在")
	}
	if video.Status != common.VideoStatusPassed {
		return nil, errno.ResourceNotFound.WithMessage("视频不存在")
	}
	go func() {
		visited, err := redis.IsIPVisited(req.VideoID, service.c.ClientIP())
		if err != nil {
			return
		}
		if !visited {
			err := redis.IncrVideoVisitCount(req.VideoID)
			if err != nil {
				return
			}
		}
		scheduler.Schdeduler.Start(
			strings.Join([]string{"video_visit_count", req.VideoID}, "/"),
			common.SyncInterval, func() {
				synchronizer.SynchronizeVideoVisitInfoRedis2DB(req.VideoID)
			},
		)

		err = redis.PutIPVisitInfo(req.VideoID, service.c.ClientIP())
		if err != nil {
			return
		}
		scheduler.Schdeduler.Start(
			strings.Join([]string{"video_visit", req.VideoID, service.c.ClientIP()}, "/"),
			common.VideoVisitInterval, func() {
				redis.DelIPVisitInfo(req.VideoID, service.c.ClientIP())
			},
		)
	}()

	var fromUser *string
	if req.AccessToken != nil {
		uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(*req.AccessToken)
		if err != nil {
			return nil, errno.AccessTokenInvalid
		}
		fromUser = &uid
		go gorse.PutFeedback(uid, fmt.Sprint(videoId), common.GorseFeedbackVisit)
	}

	resp, err := model_converter.VideoDal2Resp(video, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return resp, nil
}

func (service *VideoService) NewListEvent(req *video.VideoListReq) (*video.VideoListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("用户ID错误")
	}
	result, count, err := exquery.QueryVideoByUserIdAndStatusPaged(userId, int(req.PageNum), int(req.PageSize), common.VideoStatusPassed)
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

	items, err := model_converter.VideoListDal2Resp(&result, fromUser)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}
	return &video.VideoListRespData{
		Items:    items,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *VideoService) NewSubmitAllEvent(req *video.VideoSubmitAllReq) (*video.VideoSubmitAllRespData, error) {
	userId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	videos, count, err := exquery.QueryVideoByUserIdPaged(userId, int(req.PageNum), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	items, err := model_converter.VideoListDal2Resp(&videos, nil)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	return &video.VideoSubmitAllRespData{
		Items:    items,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *VideoService) NewSubmitReviewEvent(req *video.VideoSubmitReviewReq) (*video.VideoSubmitReviewRespData, error) {
	userId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	videos, count, err := exquery.QueryVideoByUserIdAndStatusPaged(
		userId, int(req.PageNum), int(req.PageSize), common.VideoStatusReview)
	if err != nil {
		return nil, err
	}

	items, err := model_converter.VideoListDal2Resp(&videos, nil)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &video.VideoSubmitReviewRespData{
		Items:    items,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *VideoService) NewSubmitLockedEvent(req *video.VideoSubmitLockedReq) (*video.VideoSubmitLockedRespData, error) {
	userId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	videos, count, err := exquery.QueryVideoByUserIdAndStatusPaged(
		userId, int(req.PageNum), int(req.PageSize), common.VideoStatusLocked)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	items, err := model_converter.VideoListDal2Resp(&videos, nil)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	return &video.VideoSubmitLockedRespData{
		Items:    items,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *VideoService) NewSumitPassedEvent(req *video.VideoSubmitPassedReq) (*video.VideoSubmitPassedRespData, error) {
	userId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	videos, count, err := exquery.QueryVideoByUserIdAndStatusPaged(
		userId, int(req.PageNum), int(req.PageSize), common.VideoStatusPassed)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	items, err := model_converter.VideoListDal2Resp(&videos, nil)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	return &video.VideoSubmitPassedRespData{
		Items:    items,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *VideoService) NewSearchEvent(req *video.VideoSearchReq) (*video.VideoSearchRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	result, count, err := exquery.QueryVideoFuzzyByKeywordPaged(req.Keyword, int(req.PageNum), int(req.PageSize), req.FromDate, req.ToDate)
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

	items, err := model_converter.VideoListDal2Resp(&result, fromUser)
	if err != nil {
		return nil, errno.InternalServerError.WithInnerError(err)
	}

	return &video.VideoSearchRespData{
		Items:    items,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}
