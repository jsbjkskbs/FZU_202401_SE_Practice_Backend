package service

import (
	"context"
	"fmt"
	"sfw/biz/dal"
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
	"sfw/pkg/utils/checker"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/scheduler"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gen"
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
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	err = checker.CheckVideoPublish(req.Title, req.Description, req.Category, req.Labels)
	if err != nil {
		return nil, err
	}

	videoId := generator.VideoIDGenerator.Generate()
	kv := map[string]interface{}{
		"user_id":     uid,
		"title":       req.Title,
		"description": req.Description,
		"category":    req.Category,
		"labels":      strings.Join(req.Labels, "\t"),
	}

	err = redis.VideoUploadInfoStore(fmt.Sprint(videoId), kv)
	if err != nil {
		return nil, errno.DatabaseCallError
	}
	defer scheduler.Schdeduler.Start(
		"video:"+fmt.Sprint(videoId),
		oss.VideoUploadTokenDeadline,
		func() {
			redis.VideoUploadInfoDel(fmt.Sprint(videoId))
		},
	)

	uptoken, uploadKey, err := oss.UploadVideo(fmt.Sprint(videoId), videoId)
	if err != nil {
		return nil, errno.InternalServerError
	}

	return &video.VideoPublishRespData{
		UploadURL: oss.UploadUrl,
		UploadKey: uploadKey,
		Uptoken:   uptoken,
	}, nil
}

func (service *VideoService) NewCoverUploadEvent(req *video.VideoCoverUploadReq) (*video.VideoCoverUploadRespData, error) {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}
	userId, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid
	}
	v := dal.Executor.Video
	count, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId), v.UserID.Eq(userId)).Count()
	if err != nil {
		return nil, errno.DatabaseCallError
	}
	if count == 0 {
		return nil, errno.ResourceNotFound
	}

	uptoken, key, err := oss.UploadVideoCover(req.VideoID, videoId)
	if err != nil {
		return nil, errno.InternalServerError
	}
	return &video.VideoCoverUploadRespData{
		UploadURL: oss.UploadUrl,
		UploadKey: key,
		Uptoken:   uptoken,
	}, nil
}

func (service *VideoService) NewFeedEvent(req *video.VideoFeedReq) ([]*base.Video, error) {
	var (
		vids = []string{}
		err  error
	)

	if req.Category != nil {
		vids, err = gorse.GetRecommendWithCategory("", *req.Category, 10)
	} else {
		vids, err = gorse.GetRecommend("", 10)
	}

	if err != nil {
		return nil, errno.InternalServerError
	}

	v := dal.Executor.Video
	videos := make([]*model.Video, 0)
	for _, vid := range vids {
		videoId, err := strconv.ParseInt(vid, 10, 64)
		if err != nil {
			return nil, errno.InternalServerError
		}
		video, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).First()
		if err != nil {
			return nil, errno.DatabaseCallError
		}
		videos = append(videos, video)
	}
	return model_converter.VideoListDal2Resp(&videos)
}

func (service *VideoService) NewCustomFeedEvent(req *video.VideoCustomFeedReq) ([]*base.Video, error) {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	vids := []string{}
	if req.Category != nil {
		vids, err = gorse.GetRecommendWithCategory(uid, *req.Category, 10)
	} else {
		vids, err = gorse.GetRecommend(uid, 10)
	}

	v := dal.Executor.Video
	videos := make([]*model.Video, 0)
	for _, vid := range vids {
		videoId, err := strconv.ParseInt(vid, 10, 64)
		if err != nil {
			return nil, errno.InternalServerError
		}
		video, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).First()
		if err != nil {
			return nil, errno.DatabaseCallError
		}
		videos = append(videos, video)
	}
	return model_converter.VideoListDal2Resp(&videos)
}

func (service *VideoService) NewCategoriesEvent(req *video.VideoCategoriesReq) ([]string, error) {
	return checker.Categories, nil
}

func (service *VideoService) NewInfoEvent(req *video.VideoInfoReq) (*base.Video, error) {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid
	}

	v := dal.Executor.Video
	video, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).First()
	if err != nil {
		return nil, errno.DatabaseCallError
	}
	return model_converter.VideoDal2Resp(video), nil
}

func (service *VideoService) NewListEvent(req *video.VideoListReq) (*video.VideoListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid
	}
	v := dal.Executor.Video
	result, count, err := v.WithContext(context.Background()).
		Where(v.UserID.Eq(userId), v.Status.Eq("passed")).
		FindByPage(int(req.PageNum), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError
	}
	items, err := model_converter.VideoListDal2Resp(&result)
	if err != nil {
		return nil, errno.InternalServerError
	}
	return &video.VideoListRespData{
		Items:    items,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}

func (service *VideoService) NewSubmitAllEvent(req *video.VideoSubmitAllReq) (*video.VideoSubmitAllRespData, error) {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	resp, count, err := common.QueryVideoSubmit(uid, "", req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	return &video.VideoSubmitAllRespData{
		Items:    *resp,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}

func (service *VideoService) NewSubmitReviewEvent(req *video.VideoSubmitReviewReq) (*video.VideoSubmitReviewRespData, error) {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	resp, count, err := common.QueryVideoSubmit(uid, "review", req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	return &video.VideoSubmitReviewRespData{
		Items:    *resp,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}

func (service *VideoService) NewSubmitLockedEvent(req *video.VideoSubmitLockedReq) (*video.VideoSubmitLockedRespData, error) {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	resp, count, err := common.QueryVideoSubmit(uid, "locked", req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	return &video.VideoSubmitLockedRespData{
		Items:    *resp,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}

func (service *VideoService) NewSumitPassedEvent(req *video.VideoSubmitPassedReq) (*video.VideoSubmitPassedRespData, error) {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	resp, count, err := common.QueryVideoSubmit(uid, "passed", req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	return &video.VideoSubmitPassedRespData{
		Items:    *resp,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}

func (service *VideoService) NewSearchEvent(req *video.VideoSearchReq) (*video.VideoSearchRespData, error) {
	v := dal.Executor.Video
	vd := v.WithContext(context.Background())
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	conditions := []gen.Condition{}

	if req.FromDate != nil {
		conditions = append(conditions, v.CreatedAt.Gte(*req.FromDate))
	}
	if req.ToDate != nil {
		conditions = append(conditions, v.CreatedAt.Lte(*req.ToDate))
	}
	conditions = append(conditions, v.Status.Eq("passed"))
	result, count, err := vd.Where(conditions...).
		Where(vd.Where(v.Title.Like("%"+req.Keyword+"%")).Or(v.Description.Like("%"+req.Keyword+"%"))).
		FindByPage(int(req.PageNum), int(req.PageSize))
	/*
		SELECT *
			FROM video
			WHERE
				(video.title LIKE '%keyword%' OR video.description LIKE '%keyword%')
				AND
				video.status = 'passed'
				AND
				video.created_at >= fromDate
				AND
				video.created_at <= toDate
				LIMIT pageSize OFFSET pageNum
	*/
	if err != nil {
		return nil, errno.DatabaseCallError
	}

	resp, err := model_converter.VideoListDal2Resp(&result)
	if err != nil {
		return nil, errno.InternalServerError
	}

	return &video.VideoSearchRespData{
		Items:    resp,
		IsEnd:    count <= req.PageSize*(req.PageNum+1),
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
