package service

import (
	"context"
	"strconv"

	"sfw/biz/dal"
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
	rows, err := dal.DB.Raw("SELECT * FROM Activity WHERE user_id IN (SELECT followed_id FROM Follow WHERE follower_id = ?) OR user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", uid, uid, req.PageSize, (req.PageNum)*req.PageSize).Rows()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	defer rows.Close()
	activities := []*model.Activity{}
	for rows.Next() {
		var activity model.Activity
		err = dal.DB.ScanRows(rows, &activity)
		if err != nil {
			return nil, errno.DatabaseCallError.WithInnerError(err)
		}
		activities = append(activities, &activity)
	}
	row, err := dal.DB.Raw("SELECT COUNT(*) FROM Activity WHERE user_id IN (SELECT followed_id FROM Follow WHERE follower_id = ?)", uid).Rows()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	defer row.Close()
	var count int64
	row.Next()
	row.Scan(&count)

	return &activity.ActivityFeedRespData{
		Items:    *model_converter.ActivityListDal2Resp(&activities),
		IsEnd:    count <= req.PageNum*req.PageSize,
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
			return errno.CustomError.WithMessage("无效的内容ID").WithInnerError(err)
		}
		refCount++
	}
	if req.RefVideo != nil {
		activity.RefVideoID, err = strconv.ParseInt(*req.RefVideo, 10, 64)
		if err != nil {
			return errno.CustomError.WithMessage("无效的视频ID").WithInnerError(err)
		}
		refCount++
	}
	if refCount >= 2 {
		return errno.CustomError.WithMessage("只能引用一个内容")
	}

	activity.ID = generator.ActivityIDGenerator.Generate()
	activity.UserID = uid
	activity.Content = req.Content

	a := dal.Executor.Activity
	err = a.WithContext(context.Background()).Create(&activity)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *ActivityService) NewListEvent(req *activity.ActivityListReq) (*activity.ActivityListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.CustomError.WithMessage("无效的用户ID").WithInnerError(err)
	}

	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)

	a := dal.Executor.Activity
	activities, count, err := a.WithContext(context.Background()).
		Where(a.UserID.Eq(uid)).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &activity.ActivityListRespData{
		Items:    *model_converter.ActivityListDal2Resp(&activities),
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}
