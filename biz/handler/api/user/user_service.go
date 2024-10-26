// Code generated by hertz generator.

package user

import (
	"context"
	"fmt"

	user "sfw/biz/model/api/user"
	"sfw/biz/model/base"
	"sfw/biz/mw/jwt"
	"sfw/biz/service"
	"sfw/pkg/errno"
	"sfw/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RegisterMethod .
// @router /api/v1/user/register [POST]
func RegisterMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusBadRequest, user.UserRegisterResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewUserService(ctx, c).NewRegisterEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusBadRequest, user.UserRegisterResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserRegisterResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// SecurityEmailCodeMethod .
// @router /api/v1/user/security/email/code [POST]
func SecurityEmailCodeMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserSecurityEmailCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusBadRequest, user.UserSecurityEmailCodeResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewUserService(ctx, c).NewSecurityEmailCodeEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusBadRequest, user.UserSecurityEmailCodeResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserSecurityEmailCodeResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// LoginMethod .
// @router /api/v1/user/login [POST]
func LoginMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusBadRequest, user.UserLoginResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	data, err := service.NewUserService(ctx, c).NewLoginEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusBadRequest, user.UserLoginResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	jwt.AccessTokenJwtMiddleware.LoginHandler(ctx, c)
	jwt.RefreshTokenJwtMiddleware.LoginHandler(ctx, c)

	c.JSON(consts.StatusOK, user.UserLoginResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: &base.UserWithToken{
			ID:           fmt.Sprint(data.ID),
			Username:     data.Username,
			AvatarURL:    data.AvatarURL,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
			DeletedAt:    data.DeletedAt,
			AccessToken:  c.GetString("Access-Token"),
			RefreshToken: c.GetString("Refresh-Token"),
		},
	})
}

// InfoMethod .
// @router /api/v1/user/info [GET]
func InfoMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserInfoResp)

	c.JSON(consts.StatusOK, resp)
}

// FollowerCountMethod .
// @router /api/v1/user/follower_count [GET]
func FollowerCountMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserFollowerCountReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserFollowerCountResp)

	c.JSON(consts.StatusOK, resp)
}

// FollowingCountMethod .
// @router /api/v1/user/following_count [GET]
func FollowingCountMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserFollowingCountReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserFollowingCountResp)

	c.JSON(consts.StatusOK, resp)
}

// LikeCountMethod .
// @router /api/v1/user/like_count [GET]
func LikeCountMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLikeCountReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserLikeCountResp)

	c.JSON(consts.StatusOK, resp)
}

// AvatarUploadMethod .
// @router /api/v1/user/avatar/upload [GET]
func AvatarUploadMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserAvatarUploadReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserAvatarUploadResp)

	c.JSON(consts.StatusOK, resp)
}

// MfaQrcodeMethod .
// @router /api/v1/user/mfa/qrcode [GET]
func MfaQrcodeMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserMfaQrcodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserMfaQrcodeResp)

	c.JSON(consts.StatusOK, resp)
}

// MfaBindMethod .
// @router /api/v1/user/mfa/bind [POST]
func MfaBindMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserMfaBindReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserMfaBindResp)

	c.JSON(consts.StatusOK, resp)
}

// SearchMethod .
// @router /api/v1/user/search [GET]
func SearchMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserSearchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserSearchResp)

	c.JSON(consts.StatusOK, resp)
}

// PasswordRetriveMethod .
// @router /api/v1/user/password/retrive [POST]
func PasswordRetriveMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserPasswordRetriveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserPasswordRetriveResp)

	c.JSON(consts.StatusOK, resp)
}

// PasswordResetMethod .
// @router /api/v1/user/password/reset [POST]
func PasswordResetMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserPasswordResetReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserPasswordResetResp)

	c.JSON(consts.StatusOK, resp)
}
