package service

import (
	"context"
	"fmt"
	"strconv"

	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/relation"
	"sfw/biz/mw/jwt"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

type RelationService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewRelationService(ctx context.Context, c *app.RequestContext) *RelationService {
	return &RelationService{
		ctx: ctx,
		c:   c,
	}
}

func (service *RelationService) NewFollowActionEvent(req *relation.RelationFollowActionReq) error {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	if fmt.Sprint(uid) == req.ToUserID {
		return errno.CustomError.WithMessage("不能关注自己")
	}

	toUserId, err := strconv.ParseInt(req.ToUserID, 10, 64)
	if err != nil {
		return errno.ParamInvalid.WithMessage("用户ID不合法")
	}

	exist, err := exquery.QueryUserExistByID(toUserId)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.ResourceNotFound.WithMessage("用户不存在")
	}

	switch req.ActionType {
	case common.ActionTypeOff:
		err = exquery.DeleteFollowRecord(uid, toUserId)
	case common.ActionTypeOn:
		exist, err = exquery.QueryFollowExistByFollowerIDAndFollowedID(uid, toUserId)
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if exist {
			return errno.CustomError.WithMessage("已关注")
		}
		err = exquery.InsertFollowRecord(&model.Follow{
			FollowedID: toUserId,
			FollowerID: uid,
		})
	default:
		return errno.ParamInvalid.WithMessage("操作类型不合法")
	}
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *RelationService) NewFollowListEvent(req *relation.RelationFollowListReq) (*relation.RelationFollowListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("用户ID错误")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	follows, count, err := exquery.QueryFollowingByUserIdPaged(uid, req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	users := []*model.User{}
	for _, item := range follows {
		user, err := exquery.QueryUserByID(item.FollowedID)
		if err != nil {
			return nil, errno.DatabaseCallError.WithInnerError(err)
		}
		users = append(users, user)
	}

	var fromUser *string
	if req.AccessToken != nil {
		uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(*req.AccessToken)
		if err != nil {
			return nil, errno.AccessTokenInvalid
		}
		fromUser = &uid
	}

	items, err := model_converter.UserListDal2Resp(&users, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &relation.RelationFollowListRespData{
		Items:    *items,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *RelationService) NewFollowerListEvent(req *relation.RelationFollowerListReq) (*relation.RelationFollowerListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.ParamInvalid.WithMessage("用户ID错误")
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	followers, count, err := exquery.QueryFollowerByUserIdPaged(uid, req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	users := []*model.User{}
	for _, item := range followers {
		user, err := exquery.QueryUserByID(item.FollowerID)
		if err != nil {
			return nil, errno.DatabaseCallError.WithInnerError(err)
		}
		users = append(users, user)
	}

	var fromUser *string
	if req.AccessToken != nil {
		uid, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(*req.AccessToken)
		if err != nil {
			return nil, errno.AccessTokenInvalid
		}
		fromUser = &uid
	}

	items, err := model_converter.UserListDal2Resp(&users, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &relation.RelationFollowerListRespData{
		Items:    *items,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *RelationService) NewFriendListEvent(req *relation.RelationFriendListReq) (*relation.RelationFriendListRespData, error) {
	uid, err := jwt.AccessTokenJwtMiddleware.ConvertJWTPayloadToInt64(req.AccessToken)
	if err != nil {
		return nil, errno.AccessTokenInvalid
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	friends, count, err := exquery.QueryFriendByUserIDPaged(uid, req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	users := []*model.User{}
	for _, item := range friends {
		user, err := exquery.QueryUserByID(item)
		if err != nil {
			return nil, errno.DatabaseCallError.WithInnerError(err)
		}
		users = append(users, user)
	}

	var fromUser *string
	userId := fmt.Sprint(uid)
	fromUser = &userId

	items, err := model_converter.UserListDal2Resp(&users, fromUser)
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}

	return &relation.RelationFriendListRespData{
		Items:    *items,
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}
