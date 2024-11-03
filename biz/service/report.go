package service

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/report"
	"sfw/biz/mw/jwt"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"sfw/pkg/utils/checker"
	"sfw/pkg/utils/generator"
	"strconv"

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
		return errno.CustomError.WithMessage("视频ID错误")
	}

	v := dal.Executor.Video
	exist, err := v.WithContext(context.Background()).Where(v.ID.Eq(vid)).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("视频不存在")
	}

	vr := dal.Executor.VideoReport
	count, err := vr.WithContext(context.Background()).Where(vr.UserID.Eq(uid), vr.VideoID.Eq(vid)).Count()
	if count > 3 {
		return errno.CustomError.WithMessage("您已经举报过该视频多次，请耐心等待处理结果")
	}

	err = vr.WithContext(context.Background()).Create(&model.VideoReport{
		ID:      generator.VideoReportIDGenerator.Generate(),
		UserID:  uid,
		VideoID: vid,
		Reason:  req.Content,
		Label:   req.Label,
		Status:  "unsolved",
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
		return errno.CustomError.WithMessage("动态ID错误")
	}

	a := dal.Executor.Activity
	exist, err := a.WithContext(context.Background()).Where(a.ID.Eq(aid)).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("动态不存在")
	}

	ar := dal.Executor.ActivityReport
	count, err := ar.WithContext(context.Background()).Where(ar.UserID.Eq(uid), ar.ActivityID.Eq(aid)).Count()
	if count > 3 {
		return errno.CustomError.WithMessage("您已经举报过该动态多次，请耐心等待处理结果")
	}

	err = ar.WithContext(context.Background()).Create(&model.ActivityReport{
		ID:         generator.ActivityReportIDGenerator.Generate(),
		UserID:     uid,
		ActivityID: aid,
		Reason:     req.Content,
		Label:      req.Label,
		Status:     "unsolved",
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
		return errno.CustomError.WithMessage("视频ID错误")
	}

	cid, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.CustomError.WithMessage("评论ID错误")
	}

	c := dal.Executor.VideoComment
	exist, err := c.WithContext(context.Background()).Where(c.ID.Eq(cid), c.VideoID.Eq(vid)).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("评论不存在或视频与评论索引不匹配")
	}

	cr := dal.Executor.VideoCommentReport
	count, err := cr.WithContext(context.Background()).Where(cr.UserID.Eq(uid), cr.CommentID.Eq(cid)).Count()
	if count > 3 {
		return errno.CustomError.WithMessage("您已经举报过该评论多次，请耐心等待处理结果")
	}

	err = cr.WithContext(context.Background()).Create(&model.VideoCommentReport{
		ID:        generator.VideoCommentReportIDGenerator.Generate(),
		UserID:    uid,
		CommentID: cid,
		Reason:    req.Content,
		Label:     req.Label,
		Status:    "unsolved",
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
		return errno.CustomError.WithMessage("动态ID错误")
	}

	cid, err := strconv.ParseInt(req.CommentID, 10, 64)
	if err != nil {
		return errno.CustomError.WithMessage("评论ID错误")
	}

	c := dal.Executor.ActivityComment
	exist, err := c.WithContext(context.Background()).Where(c.ID.Eq(cid), c.ActivityID.Eq(aid)).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("评论不存在或动态与评论索引不匹配")
	}

	cr := dal.Executor.ActivityCommentReport
	count, err := cr.WithContext(context.Background()).Where(cr.UserID.Eq(uid), cr.CommentID.Eq(cid)).Count()
	if count > 3 {
		return errno.CustomError.WithMessage("您已经举报过该评论多次，请耐心等待处理结果")
	}

	err = cr.WithContext(context.Background()).Create(&model.ActivityCommentReport{
		ID:        generator.ActivityCommentReportIDGenerator.Generate(),
		UserID:    uid,
		CommentID: cid,
		Reason:    req.Content,
		Label:     req.Label,
		Status:    "unsolved",
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
	case "video":
		return service.newAdminVideoCommentListEvent(req)
	case "activity":
		return service.newAdminActivityCommentListEvent(req)
	}
	return nil, errno.CustomError.WithMessage("评论类型错误")
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
	vr := dal.Executor.VideoReport
	exist, err := vr.WithContext(context.Background()).Where(vr.ID.Eq(reportId), vr.Status.Eq("unsolved")).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case 0:
		status = "rejected"
	case 1:
		status = "solved"
	default:
		return errno.CustomError.WithMessage("操作类型错误")
	}

	_, err = vr.WithContext(context.Background()).Where(vr.ID.Eq(reportId)).Updates(&model.VideoReport{
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
		return errno.CustomError.WithMessage("举报ID错误")
	}
	ar := dal.Executor.ActivityReport
	exist, err := ar.WithContext(context.Background()).Where(ar.ID.Eq(reportId), ar.Status.Eq("unsolved")).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case 0:
		status = "rejected"
	case 1:
		status = "solved"
	default:
		return errno.CustomError.WithMessage("操作类型错误")
	}

	_, err = ar.WithContext(context.Background()).Where(ar.ID.Eq(reportId)).Updates(&model.ActivityReport{
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
		return errno.CustomError.WithMessage("举报ID错误")
	}

	cr := dal.Executor.VideoCommentReport
	exist, err := cr.WithContext(context.Background()).Where(cr.ID.Eq(reportId), cr.Status.Eq("unsolved")).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case 0:
		status = "rejected"
	case 1:
		status = "solved"
	default:
		return errno.CustomError.WithMessage("操作类型错误")
	}

	_, err = cr.WithContext(context.Background()).Where(cr.ID.Eq(reportId)).Updates(&model.VideoCommentReport{
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
		return errno.CustomError.WithMessage("举报ID错误")
	}

	cr := dal.Executor.ActivityCommentReport
	exist, err := cr.WithContext(context.Background()).Where(cr.ID.Eq(reportId), cr.Status.Eq("unsolved")).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("举报不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case 0:
		status = "rejected"
	case 1:
		status = "solved"
	default:
		return errno.CustomError.WithMessage("操作类型错误")
	}

	_, err = cr.WithContext(context.Background()).Where(cr.ID.Eq(reportId)).Updates(&model.ActivityCommentReport{
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
	case "video":
		return service.newAdminVideoCommentReportHandleEvent(req)
	case "activity":
		return service.newAdminActivityCommentReportHandleEvent(req)
	}
	return errno.CustomError.WithMessage("评论类型错误")
}

func (service *ReportService) NewAdminVideoListEvent(req *report.AdminVideoListReq) (*report.AdminVideoListRespData, error) {
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	v := dal.Executor.Video
	conditions := []gen.Condition{v.Status.Eq("review")}

	if req.Category != nil {
		categoryId, ok := checker.CategoryMap[*req.Category]
		if !ok {
			return nil, errno.CustomError.WithMessage("视频分区不存在")
		}
		conditions = append(conditions, v.CategoryID.Eq(categoryId))
	}

	items, count, err := v.WithContext(context.Background()).
		Where(conditions...).FindByPage((int(req.PageNum * req.PageSize)), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	citems, err := model_converter.VideoListDal2Resp(&items)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	return &report.AdminVideoListRespData{
		Items:    citems,
		Total:    count,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsEnd:    (req.PageNum+1)*req.PageSize >= count,
	}, nil
}

func (service *ReportService) NewAdminVideoHandleEvent(req *report.AdminVideoHandleReq) error {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return errno.CustomError.WithMessage("视频ID错误")
	}

	v := dal.Executor.Video
	exist, err := v.WithContext(context.Background()).Where(v.ID.Eq(videoId), v.Status.Eq("review")).Count()
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if exist == 0 {
		return errno.CustomError.WithMessage("视频不存在或已处理")
	}

	status := ""
	switch req.ActionType {
	case 0:
		status = "locked"
	case 1:
		status = "passed"
	default:
		return errno.CustomError.WithMessage("操作类型错误")
	}

	_, err = v.WithContext(context.Background()).Where(v.ID.Eq(videoId)).Updates(&model.Video{
		Status: status,
	})
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}
