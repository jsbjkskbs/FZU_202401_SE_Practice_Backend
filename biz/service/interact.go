package service

import (
	"context"
	"strconv"
	"time"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/interact"
	"sfw/biz/mw/gorse"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/synchronizer"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/logger"
	"sfw/pkg/utils/scheduler"

	"github.com/cloudwego/hertz/pkg/app"
)

type InteractService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewInteractService(ctx context.Context, c *app.RequestContext) *InteractService {
	return &InteractService{ctx: ctx, c: c}
}

func (service *InteractService) NewLikeVideoActionEvent(req *interact.InteractLikeVideoActionReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	switch req.ActionType {
	case common.ActionTypeOn:
		go redis.AppendVideoLikeInfo(req.VideoID, uid)
	case common.ActionTypeOff:
		go redis.RemoveVideoLikeInfo(req.VideoID, uid)
	default:
		return errno.ParamInvalid
	}

	scheduler.Schdeduler.Start("video_like/"+req.VideoID, 10*time.Second, func() {
		err := synchronizer.SynchronizeVideoLikeFromRedis2DB(req.VideoID)
		if err != nil {
			logger.RuntimeLogger.Error("synchronize video like from redis to db failed, video_id: ", req.VideoID)
		}
	})
	return nil
}

func (service *InteractService) NewLikeActivityActionEvent(req *interact.InteractLikeActivityActionReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	switch req.ActionType {
	case common.ActionTypeOn:
		go redis.AppendActivityLikeInfo(req.ActivityID, uid)
	case common.ActionTypeOff:
		go redis.RemoveActivityLikeInfo(req.ActivityID, uid)
	default:
		return errno.ParamInvalid
	}
	scheduler.Schdeduler.Start("activity_like/"+req.ActivityID, 10*time.Second, func() {
		err := synchronizer.SynchronizeActivityLikeFromRedis2DB(req.ActivityID)
		if err != nil {
			logger.RuntimeLogger.Error("synchronize activity like from redis to db failed, activity_id: ", req.ActivityID)
		}
	})
	return nil
}

func (service *InteractService) newLikeVideoCommentEvent(req *interact.InteractLikeCommentActionReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	switch req.ActionType {
	case common.ActionTypeOn:
		go redis.AppendVideoCommentLikeInfo(req.FromMediaID, req.CommentID, uid)
	case common.ActionTypeOff:
		go redis.RemoveVideoCommentLikeInfo(req.FromMediaID, req.CommentID, uid)
	default:
		return errno.ParamInvalid
	}
	scheduler.Schdeduler.Start("video_comment_like/"+req.CommentID, 10*time.Second, func() {
		err := synchronizer.SynchronizeVideoCommentLikeFromRedis2DB(req.FromMediaID, req.CommentID)
		if err != nil {
			logger.RuntimeLogger.Error("synchronize video comment like from redis to db failed, comment_id: ", req.CommentID)
		}
	})
	return nil
}

func (service *InteractService) newLikeActivityCommentEvent(req *interact.InteractLikeCommentActionReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	switch req.ActionType {
	case common.ActionTypeOn:
		go redis.AppendActivityCommentLikeInfo(req.FromMediaID, req.CommentID, uid)
	case common.ActionTypeOff:
		go redis.RemoveActivityCommentLikeInfo(req.FromMediaID, req.CommentID, uid)
	default:
		return errno.ParamInvalid
	}
	scheduler.Schdeduler.Start("activity_comment_like/"+req.CommentID, 10*time.Second, func() {
		err := synchronizer.SynchronizeActivityCommentLikeFromRedis2DB(req.FromMediaID, req.CommentID)
		if err != nil {
			logger.RuntimeLogger.Error("synchronize activity comment like from redis to db failed, comment_id: ", req.CommentID)
		}
	})
	return nil
}

func (service *InteractService) NewLikeCommentEvent(req *interact.InteractLikeCommentActionReq) error {
	switch req.CommentType {
	case common.CommentTypeVideo:
		return service.newLikeVideoCommentEvent(req)
	case common.CommentTypeActivity:
		return service.newLikeActivityCommentEvent(req)
	default:
		return errno.ParamInvalid
	}
}

func (service *InteractService) NewLikeVideoListEvent(req *interact.InteractLikeVideoListReq) (*interact.InteractLikeVideoListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("无效的用户ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	videos, count, err := exquery.QueryVideoLikedByUserIdPaged(uid, int(req.PageNum), int(req.PageSize))
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

	items, err := model_converter.VideoListDal2Resp(&videos, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &interact.InteractLikeVideoListRespData{
		Items:    items,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		Total:    count,
	}, nil
}

func (service *InteractService) NewCommentVideoPublishEvent(req *interact.InteractCommentVideoPublishReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	rid := int64(0)
	pid := int64(0)
	if req.RootID != nil {
		rid, err = strconv.ParseInt(*req.RootID, 10, 64)
		if err != nil {
			return errno.ParamInvalid.WithMessage("无效的根评论ID")
		}
		if rid == 0 {
			return errno.ParamInvalid.WithMessage("无效的根评论ID")
		}
		exist, err := exquery.QueryVideoCommentExistByIdParentIdAndRootId(rid, 0, 0)
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if !exist {
			return errno.CustomError.WithMessage("根评论不存在")
		}
	}
	if req.ParentID != nil {
		if req.RootID == nil {
			return errno.ParamInvalid.WithMessage("父评论ID必须与根评论ID同时存在")
		}
		pid, err = strconv.ParseInt(*req.ParentID, 10, 64)
		if err != nil {
			return errno.ParamInvalid.WithMessage("无效的父评论ID")
		}
		if pid == 0 {
			return errno.ParamInvalid.WithMessage("无效的父评论ID")
		}
		var exist bool
		if pid == rid {
			exist, err = exquery.QueryVideoCommentExistById(pid)
		} else {
			exist, err = exquery.QueryVideoCommentExistByIdAndRootId(pid, rid)
		}
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if !exist {
			return errno.ResourceNotFound.WithMessage("父评论不存在")
		}
	}
	if pid == 0 {
		pid = rid
	}

	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("无效的视频ID")
	}
	exist, err := exquery.QueryVideoExistById(videoId)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("视频不存在")
	}

	err = exquery.InsertVideoComment(&model.VideoComment{
		ID:       generator.VideoCommentIDGenerator.Generate(),
		VideoID:  videoId,
		UserID:   uid,
		RootID:   rid,
		ParentID: pid,
		Content:  req.Content,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *InteractService) NewCommentActivityPublishEvent(req *interact.InteractCommentActivityPublishReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	rid := int64(0)
	pid := int64(0)
	if req.RootID != nil {
		rid, err = strconv.ParseInt(*req.RootID, 10, 64)
		if err != nil {
			return errno.ParamInvalid.WithMessage("无效的根评论ID")
		}
		if rid == 0 {
			return errno.ParamInvalid.WithMessage("无效的根评论ID")
		}
		exist, err := exquery.QueryActivityCommentExistByIdParentIdAndRootId(rid, 0, 0)
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if !exist {
			return errno.CustomError.WithMessage("根评论不存在")
		}
	}
	if req.ParentID != nil {
		pid, err = strconv.ParseInt(*req.ParentID, 10, 64)
		if err != nil {
			return errno.ParamInvalid.WithMessage("无效的父评论ID")
		}
		if pid == 0 {
			return errno.ParamInvalid.WithMessage("无效的父评论ID")
		}
		var exist bool
		if pid == rid {
			exist, err = exquery.QueryActivityCommentExistById(pid)
		} else {
			exist, err = exquery.QueryActivityCommentExistByIdAndRootId(pid, rid)
		}
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if !exist {
			return errno.CustomError.WithMessage("父评论不存在")
		}
	}
	if pid == 0 {
		pid = rid
	}

	activityId, err := strconv.ParseInt(req.ActivityID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("无效的动态ID")
	}
	exist, err := exquery.QueryActivityExistById(activityId)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("动态不存在")
	}

	err = exquery.InsertActivityComment(&model.ActivityComment{
		ID:         generator.ActivityIDGenerator.Generate(),
		ActivityID: activityId,
		UserID:     uid,
		RootID:     rid,
		ParentID:   pid,
		Content:    req.Content,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *InteractService) NewCommentVideoListEvent(req *interact.InteractCommentVideoListReq) (*interact.InteractCommentVideoListRespData, error) {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("无效的视频ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	exist, err := exquery.QueryVideoExistById(videoId)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return nil, errno.CustomError.WithMessage("视频不存在")
	}

	comments, count, err := exquery.QueryVideoRootCommentByVideoIdPaged(videoId, int(req.PageNum), int(req.PageSize))
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

	items, err := model_converter.VideoCommentDal2Resp(&comments, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &interact.InteractCommentVideoListRespData{
		Items:    *items,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		Total:    count,
	}, nil
}

func (service *InteractService) NewCommentActivityListEvent(req *interact.InteractCommentActivityListReq) (*interact.InteractCommentActivityListRespData, error) {
	activityId, err := strconv.ParseInt(req.ActivityID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("无效的动态ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	exist, err := exquery.QueryActivityExistById(activityId)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return nil, errno.CustomError.WithMessage("动态不存在")
	}

	comments, count, err := exquery.QueryActivityRootCommentByActivityIdPaged(activityId, int(req.PageNum), int(req.PageSize))
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

	items, err := model_converter.ActivityCommentDal2Resp(&comments, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &interact.InteractCommentActivityListRespData{
		Items:    *items,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		Total:    count,
	}, nil
}

func (service *InteractService) NewChildCommentVideoListEvent(req *interact.InteractVideoChildCommentListReq) (*interact.InteractVideoChildCommentListRespData, error) {
	commentId, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("无效的评论ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	exist, err := exquery.QueryVideoCommentExistById(commentId)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return nil, errno.CustomError.WithMessage("评论不存在")
	}

	comments, count, err := exquery.QueryVideoChildCommentByRootIdPaged(commentId, int(req.PageNum), int(req.PageSize))
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

	items, err := model_converter.VideoCommentDal2Resp(&comments, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &interact.InteractVideoChildCommentListRespData{
		Items:    *items,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		Total:    count,
	}, nil
}

func (service *InteractService) NewChildCommentActivityListEvent(req *interact.InteractActivityChildCommentListReq) (*interact.InteractActivityChildCommentListRespData, error) {
	commentId, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("无效的评论ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	exist, err := exquery.QueryActivityCommentExistById(commentId)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return nil, errno.CustomError.WithMessage("评论不存在")
	}

	comments, count, err := exquery.QueryActivityChildCommentByRootIdPaged(commentId, int(req.PageNum), int(req.PageSize))
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

	items, err := model_converter.ActivityCommentDal2Resp(&comments, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &interact.InteractActivityChildCommentListRespData{
		Items:    *items,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		Total:    count,
	}, nil
}

func (service *InteractService) NewVideoDislikeEvent(req *interact.InteractVideoDislikeReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	vid, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("无效的视频ID")
	}

	exist, err := exquery.QueryVideoExistById(vid)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.ResourceNotFound.WithMessage("视频不存在")
	}

	go gorse.PutFeedback(uid, req.VideoID, common.GorseFeedbackDislike)
	return nil
}
