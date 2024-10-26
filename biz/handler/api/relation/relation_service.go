// Code generated by hertz generator.

package relation

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	relation "sfw/biz/model/api/relation"
)

// FollowActionMethod .
// @router /api/v1/relation/follow/action [POST]
func FollowActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationFollowActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.RelationFollowActionResp)

	c.JSON(consts.StatusOK, resp)
}

// FollowListMethod .
// @router /api/v1/relation/follow/list [GET]
func FollowListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationFollowListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.RelationFollowListResp)

	c.JSON(consts.StatusOK, resp)
}

// FollowedListMethod .
// @router /api/v1/relation/followed/list [GET]
func FollowedListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationFollowedListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.RelationFollowedListResp)

	c.JSON(consts.StatusOK, resp)
}

// FriendListMethod .
// @router /api/v1/relation/friend/list [GET]
func FriendListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationFriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.RelationFriendListResp)

	c.JSON(consts.StatusOK, resp)
}
