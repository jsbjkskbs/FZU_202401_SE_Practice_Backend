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
	"sfw/pkg/oss"
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
		c.JSON(consts.StatusOK, user.UserRegisterResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewUserService(ctx, c).NewRegisterEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserRegisterResp{
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
		c.JSON(consts.StatusOK, user.UserSecurityEmailCodeResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewUserService(ctx, c).NewSecurityEmailCodeEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserSecurityEmailCodeResp{
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
		c.JSON(consts.StatusOK, user.UserLoginResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	data, err := service.NewUserService(ctx, c).NewLoginEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserLoginResp{
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
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserInfoResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewUserService(ctx, c).NewInfoEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserInfoResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserInfoResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: &base.User{
			ID:        fmt.Sprint(resp.ID),
			Username:  resp.Username,
			AvatarURL: resp.AvatarURL,
			CreatedAt: resp.CreatedAt,
			UpdatedAt: resp.UpdatedAt,
			DeletedAt: resp.DeletedAt,
		},
	})
}

// FollowerCountMethod .
// @router /api/v1/user/follower_count [GET]
func FollowerCountMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserFollowerCountReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserFollowerCountResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewUserService(ctx, c).NewFollowerCountEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserFollowerCountResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserFollowerCountResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: &user.UserFollowerCountRespData{
			ID:            req.UserID,
			FollowerCount: resp,
		},
	})
}

// FollowingCountMethod .
// @router /api/v1/user/following_count [GET]
func FollowingCountMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserFollowingCountReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserFollowingCountResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewUserService(ctx, c).NewFollowingCountEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserFollowingCountResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserFollowingCountResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: &user.UserFollowingCountRespData{
			ID:             req.UserID,
			FollowingCount: resp,
		},
	})
}

// LikeCountMethod .
// @router /api/v1/user/like_count [GET]
func LikeCountMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLikeCountReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserLikeCountResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewUserService(ctx, c).NewLikeCountEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserLikeCountResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserLikeCountResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: &user.UserLikeCountRespData{
			ID:        req.UserID,
			LikeCount: resp,
		},
	})
}

// AvatarUploadMethod .
// @router /api/v1/user/avatar/upload [GET]
func AvatarUploadMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserAvatarUploadReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserAvatarUploadResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewUserService(ctx, c).NewAvatarUploadEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserAvatarUploadResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserAvatarUploadResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: &user.UserAvatarUploadData{
			UploadURL: oss.UploadUrl,
			Uptoken:   resp,
		},
	})
}

// MfaQrcodeMethod .
// @router /api/v1/user/mfa/qrcode [GET]
func MfaQrcodeMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserMfaQrcodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserMfaQrcodeResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, err := service.NewUserService(ctx, c).NewMfaQrcodeEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserMfaQrcodeResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserMfaQrcodeResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: resp,
	})
}

// MfaBindMethod .
// @router /api/v1/user/mfa/bind [POST]
func MfaBindMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserMfaBindReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserMfaBindResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewUserService(ctx, c).NewMfaBindEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserMfaBindResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}
	c.JSON(consts.StatusOK, user.UserMfaBindResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// SearchMethod .
// @router /api/v1/user/search [GET]
func SearchMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserSearchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserSearchResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	resp, isEnd, pn, ps, err := service.NewUserService(ctx, c).NewSearchEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserSearchResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserSearchResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
		Data: &user.UserSearchRespData{
			IsEnd:    isEnd,
			Items:    *resp,
			PageNum:  pn,
			PageSize: ps,
		},
	})
}

// PasswordRetrieveMethod .
// @router /api/v1/user/security/password/retrieve [POST]
func PasswordRetrieveMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserPasswordRetrieveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserPasswordRetrieveResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewUserService(ctx, c).NewSecurityPasswordRetrieve(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserPasswordRetrieveResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserPasswordRetrieveResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}

// PasswordResetMethod .
// @router /api/v1/user/password/reset [POST]
func PasswordResetMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserPasswordResetReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserPasswordResetResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	err = service.NewUserService(ctx, c).NewSecurityPasswordResetEvent(&req)
	if err != nil {
		resp := utils.CreateBaseHttpResponse(err)
		c.JSON(consts.StatusOK, user.UserPasswordResetResp{
			Code: resp.Code,
			Msg:  resp.Msg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.UserPasswordResetResp{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	})
}
