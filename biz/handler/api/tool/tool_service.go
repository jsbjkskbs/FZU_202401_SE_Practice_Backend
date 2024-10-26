// Code generated by hertz generator.

package tool

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	tool "sfw/biz/model/api/tool"
)

// ToolDelete .
// @router /api/v1/tool/delete [DELETE]
func ToolDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tool.ToolDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tool.ToolDeleteResp)

	c.JSON(consts.StatusOK, resp)
}

// AdminToolDelete .
// @router /api/v1/admin/tool/delete [DELETE]
func AdminToolDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tool.AdminToolDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tool.AdminToolDeleteResp)

	c.JSON(consts.StatusOK, resp)
}

// ToolUploadImage .
// @router /api/v1/tool/upload/image [GET]
func ToolUploadImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tool.ToolUploadImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tool.ToolUploadImageResp)

	c.JSON(consts.StatusOK, resp)
}

// ToolGetImage .
// @router /api/v1/tool/get/image [GET]
func ToolGetImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tool.ToolGetImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tool.ToolGetImageResp)

	c.JSON(consts.StatusOK, resp)
}

// UserRefresh .
// @router /api/v1/user/refresh [GET]
func UserRefresh(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tool.UserRefreshReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tool.UserRefreshResp)

	c.JSON(consts.StatusOK, resp)
}