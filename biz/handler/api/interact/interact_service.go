// Code generated by hertz generator.

package interact

import (
	"context"

	interact "sfw/biz/model/api/interact"
	"sfw/biz/service"
	"sfw/pkg/errno"
	"sfw/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// InteractMessageSendMethod .
// @router /api/v1/interact/message/send [POST]
func InteractMessageSendMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractMessageSendReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.InteractMessageSendResp)

	c.JSON(consts.StatusOK, resp)
}

// InteractLikeVideoActionMethod .
// @router /api/v1/interact/like/video/action [POST]
func InteractLikeVideoActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractLikeVideoActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeVideoActionResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewInteractService(ctx, c).NewLikeVideoActionEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeVideoActionResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractLikeVideoActionResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// InteractLikeActivityActionMethod .
// @router /api/v1/interact/like/activity/action [POST]
func InteractLikeActivityActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractLikeActivityActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeActivityActionResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewInteractService(ctx, c).NewLikeActivityActionEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeActivityActionResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractLikeActivityActionResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// InteractLikeCommentActionMethod .
// @router /api/v1/interact/like/comment/action [POST]
func InteractLikeCommentActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractLikeCommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeCommentActionResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewInteractService(ctx, c).NewLikeCommentEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeCommentActionResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractLikeCommentActionResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// InteractLikeVideoListMethod .
// @router /api/v1/interact/like/video/list [GET]
func InteractLikeVideoListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractLikeVideoListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeVideoListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewInteractService(ctx, c).NewLikeVideoListEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractLikeVideoListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractLikeVideoListResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: resp,
	})
}

// InteractCommentVideoPublishMethod .
// @router /api/v1/interact/comment/video/publish [POST]
func InteractCommentVideoPublishMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractCommentVideoPublishReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentVideoPublishResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewInteractService(ctx, c).NewCommentVideoPublishEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentVideoPublishResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractCommentVideoPublishResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// InteractCommentActivityPublishMethod .
// @router /api/v1/interact/comment/activity/publish [POST]
func InteractCommentActivityPublishMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractCommentActivityPublishReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentActivityPublishResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewInteractService(ctx, c).NewCommentActivityPublishEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentActivityPublishResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractCommentActivityPublishResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// InteractCommentVideoListMethod .
// @router /api/v1/interact/comment/video/list [GET]
func InteractCommentVideoListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractCommentVideoListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentVideoListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewInteractService(ctx, c).NewCommentVideoListEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentVideoListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractCommentVideoListResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: resp,
	})
}

// InteractCommentActivityListMethod .
// @router /api/v1/interact/comment/activity/list [GET]
func InteractCommentActivityListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractCommentActivityListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentActivityListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewInteractService(ctx, c).NewCommentActivityListEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractCommentActivityListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractCommentActivityListResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: resp,
	})
}

// InteractVideoChildCommentListMethod .
// @router /api/v1/interact/video/child_comment/list [GET]
func InteractVideoChildCommentListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractVideoChildCommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractVideoChildCommentListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewInteractService(ctx, c).NewChildCommentVideoListEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractVideoChildCommentListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractVideoChildCommentListResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: resp,
	})
}

// InteractActivityChildCommentListMethod .
// @router /api/v1/interact/activity/child_comment/list [GET]
func InteractActivityChildCommentListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.InteractActivityChildCommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractActivityChildCommentListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewInteractService(ctx, c).NewChildCommentActivityListEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, interact.InteractActivityChildCommentListResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, interact.InteractActivityChildCommentListResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: resp,
	})
}
