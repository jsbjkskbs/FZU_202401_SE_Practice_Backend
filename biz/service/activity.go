package service

import (
	"context"

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

/*
func (service *ActivityService) NewPublishEvent(req *activity.ActivityPublishReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return err
	}
	var refCount int64
	if req.RefActivity != nil {
		refCount++
	}
	if req.RefVideo != nil {
		refCount++
	}
	if refCount >= 2 {
		return errno.CustomError.WithMessage("只能引用一个内容")
	}

	a := dal.Executor.Activity
	_, err = a.WithContext(context.Background()).Create(&model.Activity{
		ID: generator.ActivityIDGenerator.Generate(),
		UserID: uid,
		Content: req.Content,

	}

}
*/
