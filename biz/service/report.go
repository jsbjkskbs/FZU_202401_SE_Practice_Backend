package service

import (
	"context"
	"fmt"
	"strconv"

	"sfw/biz/dal"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/report"
	"sfw/biz/mw/jwt"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/utils/checker"
	"sfw/pkg/utils/generator"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gen"
)

type ReportService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewReportService(ctx context.Context, c *app.RequestContext) *ReportService {
	return &ReportService{
		ctx: ctx,
		c:   c,
	}
}

func (service *ReportService) NewReportVideoEvent(req *report.ReportVideoReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	vid, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("视频ID错误")
	}

	exist, err := exquery.QueryVideoExistById(vid)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("视频不存在")
	}

	count, err := exquery.QueryVideoReportCountByUserIdAndVideoId(uid, vid)
	if count >= common.ReportLimit {
		return errno.CustomError.WithMessage("您已经举报过该视频多次，请耐心等待处理结果")
	}

	err = exquery.InsertVideoReport(&model.VideoReport{
		ID:      generator.VideoReportIDGenerator.Generate(),
		UserID:  uid,
		VideoID: vid,
		Reason:  req.Content,
		Label:   req.Label,
		Status:  common.ReportUnresolved,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) NewReportActivityEvent(req *report.ReportActivityReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	aid, err := strconv.ParseInt(req.ActivityID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("动态ID错误")
	}

	exist, err := exquery.QueryActivityExistById(aid)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("动态不存在")
	}

	count, err := exquery.QueryActivityReportCountByUserIdAndActivityId(uid, aid)
	if count >= common.ReportLimit {
		return errno.CustomError.WithMessage("您已经举报过该动态多次，请耐心等待处理结果")
	}

	err = exquery.InsertActivityReport(&model.ActivityReport{
		ID:         generator.ActivityReportIDGenerator.Generate(),
		UserID:     uid,
		ActivityID: aid,
		Reason:     req.Content,
		Label:      req.Label,
		Status:     common.ReportUnresolved,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) newReportVideoCommentEvent(req *report.ReportCommentReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	vid, err := strconv.ParseInt(req.FromMediaID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("视频ID错误")
	}

	cid, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("评论ID错误")
	}

	exist, err := exquery.QueryVideoCommentExistByIdAndVideoId(cid, vid)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("评论不存在或视频与评论索引不匹配")
	}

	count, err := exquery.QueryVideoCommentReportCountByUserIdAndCommentId(uid, cid)
	if count >= common.ReportLimit {
		return errno.CustomError.WithMessage("您已经举报过该评论多次，请耐心等待处理结果")
	}

	err = exquery.InsertVideoCommentReport(&model.VideoCommentReport{
		ID:        generator.VideoCommentReportIDGenerator.Generate(),
		UserID:    uid,
		CommentID: cid,
		Reason:    req.Content,
		Label:     req.Label,
		Status:    common.ReportUnresolved,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) newReportActivityCommentEvent(req *report.ReportCommentReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}

	aid, err := strconv.ParseInt(req.FromMediaID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("动态ID错误")
	}

	cid, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("评论ID错误")
	}

	exist, err := exquery.QueryActivityCommentExistByIdAndActivityId(cid, aid)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("评论不存在或动态与评论索引不匹配")
	}

	count, err := exquery.QueryActivityCommentReportCountByUserIdAndCommentId(uid, cid)
	if count >= common.ReportLimit {
		return errno.CustomError.WithMessage("您已经举报过该评论多次，请耐心等待处理结果")
	}

	err = exquery.InsertActivityCommentReport(&model.ActivityCommentReport{
		ID:        generator.ActivityCommentReportIDGenerator.Generate(),
		UserID:    uid,
		CommentID: cid,
		Reason:    req.Content,
		Label:     req.Label,
		Status:    common.ReportUnresolved,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) NewReportCommentEvent(req *report.ReportCommentReq) error {
	switch req.CommentType {
	case "video":
		return service.newReportVideoCommentEvent(req)
	case "activity":
		return service.newReportActivityCommentEvent(req)
	}
	return errno.CustomError.WithMessage("评论类型错误")
}

func (service *ReportService) NewAdminVideoReportListEvent(req *report.AdminVideoReportListReq) (*report.AdminVideoReportListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	vr := dal.Executor.VideoReport
	conditions := []gen.Condition{}
	if req.Status != nil {
		conditions = append(conditions, vr.Status.Eq(*req.Status))
	}
	if req.Keyword != nil {
		conditions = append(conditions, vr.Reason.Like(fmt.Sprint("%", *req.Keyword, "%")))
	}
	if req.UserID != nil {
		userId, err := strconv.ParseInt(*req.UserID, 10, 64)
		if err != nil {
			return nil, errno.CustomError.WithMessage("用户ID错误")
		}
		conditions = append(conditions, vr.UserID.Eq(userId))
	}
	if req.Label != nil {
		conditions = append(conditions, vr.Label.Eq(*req.Label))
	}
	// 此处代码不必提取至exquery
	items, count, err := vr.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(req.PageNum * req.PageSize)), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &report.AdminVideoReportListRespData{
		Items:    model_converter.VideoReportDal2Resp(&items),
		Total:    count,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    (req.PageNum+1)*req.PageSize >= count,
	}, nil
}

func (service *ReportService) NewAdminActivityReportListEvent(req *report.AdminActivityReportListReq) (*report.AdminActivityReportListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	ar := dal.Executor.ActivityReport
	conditions := []gen.Condition{}
	if req.Status != nil {
		conditions = append(conditions, ar.Status.Eq(*req.Status))
	}
	if req.Keyword != nil {
		conditions = append(conditions, ar.Reason.Like(fmt.Sprint("%", *req.Keyword, "%")))
	}
	if req.UserID != nil {
		userId, err := strconv.ParseInt(*req.UserID, 10, 64)
		if err != nil {
			return nil, errno.CustomError.WithMessage("用户ID错误")
		}
		conditions = append(conditions, ar.UserID.Eq(userId))
	}
	if req.Label != nil {
		conditions = append(conditions, ar.Label.Eq(*req.Label))
	}
	// 此处代码不必提取至exquery
	items, count, err := ar.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(req.PageNum * req.PageSize)), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &report.AdminActivityReportListRespData{
		Items:    model_converter.ActivityReportDal2Resp(&items),
		Total:    count,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    (req.PageNum+1)*req.PageSize >= count,
	}, nil
}

func (service *ReportService) newAdminVideoCommentListEvent(req *report.AdminCommentReportListReq) (*report.AdminCommentReportListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	cr := dal.Executor.VideoCommentReport
	conditions := []gen.Condition{}
	if req.Status != nil {
		conditions = append(conditions, cr.Status.Eq(*req.Status))
	}
	if req.Keyword != nil {
		conditions = append(conditions, cr.Reason.Like(fmt.Sprint("%", *req.Keyword, "%")))
	}
	if req.UserID != nil {
		userId, err := strconv.ParseInt(*req.UserID, 10, 64)
		if err != nil {
			return nil, errno.CustomError.WithMessage("用户ID错误")
		}
		conditions = append(conditions, cr.UserID.Eq(userId))
	}
	if req.Label != nil {
		conditions = append(conditions, cr.Label.Eq(*req.Label))
	}
	// 此处代码不必提取至exquery
	items, count, err := cr.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(req.PageNum * req.PageSize)), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &report.AdminCommentReportListRespData{
		Items:    model_converter.VideoCommentReportDal2Resp(&items),
		Total:    count,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    (req.PageNum+1)*req.PageSize >= count,
	}, nil
}

func (service *ReportService) newAdminActivityCommentListEvent(req *report.AdminCommentReportListReq) (*report.AdminCommentReportListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	cr := dal.Executor.ActivityCommentReport
	conditions := []gen.Condition{}
	if req.Status != nil {
		conditions = append(conditions, cr.Status.Eq(*req.Status))
	}
	if req.Keyword != nil {
		conditions = append(conditions, cr.Reason.Like(fmt.Sprint("%", *req.Keyword, "%")))
	}
	if req.UserID != nil {
		userId, err := strconv.ParseInt(*req.UserID, 10, 64)
		if err != nil {
			return nil, errno.CustomError.WithMessage("用户ID错误")
		}
		conditions = append(conditions, cr.UserID.Eq(userId))
	}
	if req.Label != nil {
		conditions = append(conditions, cr.Label.Eq(*req.Label))
	}
	// 此处代码不必提取至exquery
	items, count, err := cr.WithContext(context.Background()).
		Where(conditions...).
		FindByPage((int(req.PageNum * req.PageSize)), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &report.AdminCommentReportListRespData{
		Items:    model_converter.ActivityCommentReportDal2Resp(&items),
		Total:    count,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    (req.PageNum+1)*req.PageSize >= count,
	}, nil
}

func (service *ReportService) NewAdminCommentReportListEvent(req *report.AdminCommentReportListReq) (*report.AdminCommentReportListRespData, error) {
	switch req.CommentType {
	case common.CommentTypeVideo:
		return service.newAdminVideoCommentListEvent(req)
	case common.CommentTypeActivity:
		return service.newAdminActivityCommentListEvent(req)
	}
	return nil, errno.ParamInvalid.WithMessage("评论类型错误")
}

func (service *ReportService) NewAdminVideoReportHandleEvent(req *report.AdminVideoReportHandleReq) error {
	adminId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	reportId, err := strconv.ParseInt(req.ReportID, 10, 64)
	if err != nil {
		return errno.CustomError.WithMessage("举报ID错误")
	}
	exist, err := exquery.QueryVideoReportExistByIdAndStatus(reportId, common.ReportUnresolved)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case common.ActionTypeOff:
		status = common.ReportRejected
	case common.ActionTypeOn:
		status = common.ReportResolved
	default:
		return errno.ParamInvalid.WithMessage("操作类型错误")
	}

	err = exquery.UpdateVideoReportById(&model.VideoReport{
		Status:  status,
		AdminID: adminId,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) NewAdminActivityReportHandleEvent(req *report.AdminActivityReportHandleReq) error {
	adminId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	reportId, err := strconv.ParseInt(req.ReportID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("举报ID错误")
	}
	exist, err := exquery.QueryActivityReportExistByIdAndStatus(reportId, common.ReportUnresolved)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case common.ActionTypeOff:
		status = common.ReportRejected
	case common.ActionTypeOn:
		status = common.ReportResolved
	default:
		return errno.ParamInvalid.WithMessage("操作类型错误")
	}

	err = exquery.UpdateActivityReportById(&model.ActivityReport{
		Status:  status,
		AdminID: adminId,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) newAdminVideoCommentReportHandleEvent(req *report.AdminCommentReportHandleReq) error {
	adminId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	reportId, err := strconv.ParseInt(req.ReportID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("举报ID错误")
	}

	exist, err := exquery.QueryVideoCommentReportExistByIdAndStatus(reportId, common.ReportUnresolved)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case common.ActionTypeOff:
		status = common.ReportRejected
	case common.ActionTypeOn:
		status = common.ReportResolved
	default:
		return errno.ParamInvalid.WithMessage("操作类型错误")
	}

	err = exquery.UpdateVideoCommentReportById(&model.VideoCommentReport{
		Status:  status,
		AdminID: adminId,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) newAdminActivityCommentReportHandleEvent(req *report.AdminCommentReportHandleReq) error {
	adminId, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	reportId, err := strconv.ParseInt(req.ReportID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("举报ID错误")
	}

	exist, err := exquery.QueryActivityCommentReportExistByIdAndStatus(reportId, common.ReportUnresolved)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case common.ActionTypeOff:
		status = common.ReportRejected
	case common.ActionTypeOn:
		status = common.ReportResolved
	default:
		return errno.ParamInvalid.WithMessage("操作类型错误")
	}

	err = exquery.UpdateActivityCommentReportById(&model.ActivityCommentReport{
		Status:  status,
		AdminID: adminId,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ReportService) NewAdminCommentReportHandleEvent(req *report.AdminCommentReportHandleReq) error {
	switch req.CommentType {
	case common.CommentTypeVideo:
		return service.newAdminVideoCommentReportHandleEvent(req)
	case common.CommentTypeActivity:
		return service.newAdminActivityCommentReportHandleEvent(req)
	}
	return errno.ParamInvalid.WithMessage("评论类型错误")
}

func (service *ReportService) NewAdminVideoListEvent(req *report.AdminVideoListReq) (*report.AdminVideoListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	v := dal.Executor.Video
	conditions := []gen.Condition{v.Status.Eq(common.VideoStatusReview)}

	if req.Category != nil {
		categoryId, ok := checker.CategoryMap[*req.Category]
		if !ok {
			return nil, errno.ParamInvalid.WithMessage("视频分区不存在")
		}
		conditions = append(conditions, v.CategoryID.Eq(categoryId))
	}

	// 此处代码不必提取至exquery
	videos, count, err := v.WithContext(context.Background()).
		Where(conditions...).FindByPage((int(req.PageNum * req.PageSize)), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	items, err := model_converter.VideoListDal2Resp(&videos)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &report.AdminVideoListRespData{
		Items:    items,
		Total:    count,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    (req.PageNum+1)*req.PageSize >= count,
	}, nil
}

func (service *ReportService) NewAdminVideoHandleEvent(req *report.AdminVideoHandleReq) error {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("视频ID错误")
	}

	exist, err := exquery.QueryVideoExistByIdAndStatus(videoId, common.VideoStatusReview)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("视频不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case common.ActionTypeOff:
		status = common.VideoStatusLocked
	case common.ActionTypeOn:
		status = common.VideoStatusPassed
	default:
		return errno.ParamInvalid.WithMessage("操作类型错误")
	}

	err = exquery.UpdateVideoWithId(&model.Video{
		ID:     videoId,
		Status: status,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}
