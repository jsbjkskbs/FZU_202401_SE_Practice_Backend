package service

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/interact"
	"sfw/biz/mw/jwt"
	"sfw/biz/mw/redis"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/synchronizer"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/scheduler"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
		return errno.CustomError.WithMessage("无效的操作类型")
	}

	scheduler.Schdeduler.Start("video_like/"+req.VideoID, 3*time.Minute, func() {
		err := synchronizer.SynchronizeVideoLikeFromRedis2DB(req.VideoID)
		if err != nil {
			hlog.Info("synchronize video like from redis to db failed, video_id: ", req.VideoID)
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
		return errno.CustomError.WithMessage("无效的操作类型")
	}
	scheduler.Schdeduler.Start("activity_like/"+req.ActivityID, 3*time.Minute, func() {
		err := synchronizer.SynchronizeActivityLikeFromRedis2DB(req.ActivityID)
		if err != nil {
			hlog.Info("synchronize activity like from redis to db failed, activity_id: ", req.ActivityID)
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
		go redis.AppendVideoCommentLikeInfo(req.CommentID, uid)
	case common.ActionTypeOff:
		go redis.RemoveVideoCommentLikeInfo(req.CommentID, uid)
	default:
		return errno.CustomError.WithMessage("无效的操作类型")
	}
	scheduler.Schdeduler.Start("video_comment_like/"+req.CommentID, 3*time.Minute, func() {
		err := synchronizer.SynchronizeVideoCommentLikeFromRedis2DB(req.CommentID)
		if err != nil {
			hlog.Info("synchronize video comment like from redis to db failed, comment_id: ", req.CommentID)
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
		go redis.AppendActivityCommentLikeInfo(req.CommentID, uid)
	case common.ActionTypeOff:
		go redis.RemoveActivityCommentLikeInfo(req.CommentID, uid)
	default:
		return errno.CustomError.WithMessage("无效的操作类型")
	}
	scheduler.Schdeduler.Start("activity_comment_like/"+req.CommentID, 3*time.Minute, func() {
		err := synchronizer.SynchronizeActivityCommentLikeFromRedis2DB(req.CommentID)
		if err != nil {
			hlog.Info("synchronize activity comment like from redis to db failed, comment_id: ", req.CommentID)
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
		return errno.CustomError.WithMessage("无效的评论类型")
	}
}

func (service *InteractService) NewLikeVideoListEvent(req *interact.InteractLikeVideoListReq) (*interact.InteractLikeVideoListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	rows, err := dal.DB.Raw(
		`SELECT v.*  
		FROM Video v  
		JOIN (  
			SELECT video_id, created_at  
			FROM VideoLike  
			WHERE user_id = ?  
			ORDER BY created_at DESC  
			LIMIT ?, ?  
		) vl ON v.id = vl.video_id;`,
		uid, req.PageNum*req.PageSize, req.PageSize,
	).Rows()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	defer rows.Close()
	row, err := dal.DB.Raw(
		`SELECT COUNT(*) FROM Video WHERE id IN (SELECT video_id FROM VideoLike WHERE user_id = ?)`,
		uid,
	).Rows()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	defer row.Close()

	videos := make([]*model.Video, 0, req.PageSize)
	for rows.Next() {
		var video model.Video
		dal.DB.ScanRows(rows, &video)
		videos = append(videos, &video)
	}
	items, err := model_converter.VideoListDal2Resp(&videos)
	if err != nil {
		return nil, err
	}

	var count int64
	row.Next()
	row.Scan(&count)

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

	vc := dal.Executor.VideoComment
	rid := int64(0)
	pid := int64(0)
	if req.RootID != nil {
		rid, err = strconv.ParseInt(*req.RootID, 10, 64)
		if err != nil {
			return errno.CustomError.WithMessage("无效的根评论ID")
		}
		if rid == 0 {
			return errno.CustomError.WithMessage("无效的根评论ID")
		}
		exist, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(rid), vc.RootID.Eq(0)).Count()
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if exist == 0 {
			return errno.CustomError.WithMessage("根评论不存在")
		}
	}
	if req.ParentID != nil {
		if req.RootID == nil {
			return errno.CustomError.WithMessage("父评论ID必须与根评论ID同时存在")
		}
		pid, err = strconv.ParseInt(*req.ParentID, 10, 64)
		if err != nil {
			return errno.CustomError.WithMessage("无效的父评论ID")
		}
		if pid == 0 {
			return errno.CustomError.WithMessage("无效的父评论ID")
		}
		exist, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(pid), vc.RootID.Eq(rid)).Count()
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if exist == 0 {
			return errno.CustomError.WithMessage("父评论不存在")
		}
	}
	if pid == 0 {
		pid = rid
	}

	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.CustomError.WithMessage("无效的视频ID")
	}
	v := dal.Executor.Video
	exist, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("视频不存在")
	}

	err = vc.WithContext(context.Background()).Create(&model.VideoComment{
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

	ac := dal.Executor.ActivityComment
	rid := int64(0)
	pid := int64(0)
	if req.RootID != nil {
		rid, err = strconv.ParseInt(*req.RootID, 10, 64)
		if err != nil {
			return errno.CustomError.WithMessage("无效的根评论ID")
		}
		if rid == 0 {
			return errno.CustomError.WithMessage("无效的根评论ID")
		}
		exist, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(rid), ac.RootID.Eq(0), ac.ParentID.Eq(0)).Count()
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if exist == 0 {
			return errno.CustomError.WithMessage("根评论不存在")
		}
	}
	if req.ParentID != nil {
		pid, err = strconv.ParseInt(*req.ParentID, 10, 64)
		if err != nil {
			return errno.CustomError.WithMessage("无效的父评论ID")
		}
		if pid == 0 {
			return errno.CustomError.WithMessage("无效的父评论ID")
		}
		exist, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(pid), ac.RootID.Eq(rid)).Count()
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if exist == 0 {
			return errno.CustomError.WithMessage("父评论不存在")
		}
	}
	if pid == 0 {
		pid = rid
	}

	activityId, err := strconv.ParseInt(req.ActivityID, 10, 64)
	if err != nil {
		return errno.CustomError.WithMessage("无效的活动ID")
	}
	a := dal.Executor.Activity
	exist, err := a.WithContext(context.Background()).Where(a.ID.Eq(activityId)).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("活动不存在")
	}

	err = ac.WithContext(context.Background()).Create(&model.ActivityComment{
		ID:         generator.ActivityIDGenerator.Generate(),
		ActivityID: activityId,
		UserID:     uid,
		RootID:     rid,
		ParentID:   pid,
		Content:    req.Content,
	})
	return nil
}

func (service *InteractService) NewCommentVideoListEvent(req *interact.InteractCommentVideoListReq) (*interact.InteractCommentVideoListRespData, error) {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, errno.CustomError.WithMessage("无效的视频ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	v := dal.Executor.Video
	exist, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).Count()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return nil, errno.CustomError.WithMessage("视频不存在")
	}

	vc := dal.Executor.VideoComment
	comments, count, err := vc.WithContext(context.Background()).
		Where(vc.VideoID.Eq(videoId), vc.RootID.Eq(0), vc.ParentID.Eq(0)).
		Order(vc.CreatedAt.Desc()).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))

	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	items, err := model_converter.VideoCommentDal2Resp(&comments)
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
		return nil, errno.CustomError.WithMessage("无效的动态ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	a := dal.Executor.Activity
	exist, err := a.WithContext(context.Background()).Where(a.ID.Eq(activityId)).Count()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return nil, errno.CustomError.WithMessage("动态不存在")
	}

	ac := dal.Executor.ActivityComment
	comments, count, err := ac.WithContext(context.Background()).
		Where(ac.ActivityID.Eq(activityId), ac.RootID.Eq(0), ac.ParentID.Eq(0)).
		Order(ac.CreatedAt.Desc()).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))

	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	items, err := model_converter.ActivityCommentDal2Resp(&comments)
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
		return nil, errno.CustomError.WithMessage("无效的评论ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	vc := dal.Executor.VideoComment
	exist, err := vc.WithContext(context.Background()).Where(vc.ID.Eq(commentId)).Count()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return nil, errno.CustomError.WithMessage("评论不存在")
	}

	comments, count, err := vc.WithContext(context.Background()).
		Where(vc.RootID.Eq(commentId)).
		Order(vc.CreatedAt.Desc()).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))

	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	items, err := model_converter.VideoCommentDal2Resp(&comments)
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
		return nil, errno.CustomError.WithMessage("无效的评论ID")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	ac := dal.Executor.ActivityComment
	exist, err := ac.WithContext(context.Background()).Where(ac.ID.Eq(commentId)).Count()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return nil, errno.CustomError.WithMessage("评论不存在")
	}

	comments, count, err := ac.WithContext(context.Background()).
		Where(ac.RootID.Eq(commentId)).
		Order(ac.CreatedAt.Desc()).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))

	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	items, err := model_converter.ActivityCommentDal2Resp(&comments)
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
