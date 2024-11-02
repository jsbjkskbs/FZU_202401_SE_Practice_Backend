package service

import (
	"context"
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/exquery"
	"sfw/biz/dal/model"
	"sfw/biz/model/api/relation"
	"sfw/biz/mw/jwt"
	"sfw/biz/service/common"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"strconv"

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
		return err
	}
	if fmt.Sprint(uid) == req.ToUserID {
		return errno.CustomError.WithMessage("不能关注自己")
	}

	toUserId, err := strconv.ParseInt(req.ToUserID, 10, 64)
	if err != nil {
		return errno.CustomError.WithMessage("无效的用户ID").WithInnerError(err)
	}

	exist, err := exquery.QueryUserExistByID(toUserId)
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	if !exist {
		return errno.CustomError.WithMessage("用户不存在")
	}

	f := dal.Executor.Follow
	switch req.ActionType {
	case common.ActionTypeOff:
		_, err = f.WithContext(context.Background()).Delete(&model.Follow{
			FollowedID: toUserId,
			FollowerID: uid,
		})
	case common.ActionTypeOn:
		count, err := f.WithContext(context.Background()).Where(f.FollowedID.Eq(toUserId), f.FollowerID.Eq(uid)).Count()
		if err != nil {
			return errno.DatabaseCallError.WithInnerError(err)
		}
		if count > 0 {
			return errno.CustomError.WithMessage("已关注")
		}
		err = f.WithContext(context.Background()).Create(&model.Follow{
			FollowedID: toUserId,
			FollowerID: uid,
		})
	}
	if err != nil {
		return errno.DatabaseCallError.WithInnerError(err)
	}
	return nil
}

func (service *RelationService) NewFollowListEvent(req *relation.RelationFollowListReq) (*relation.RelationFollowListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.CustomError.WithMessage("无效的用户ID").WithInnerError(err)
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	f := dal.Executor.Follow
	follows, count, err := f.WithContext(context.Background()).
		Where(f.FollowerID.Eq(uid)).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))
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

	return &relation.RelationFollowListRespData{
		Items:    *model_converter.UserListDal2Resp(&users),
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

func (service *RelationService) NewFollowerListEvent(req *relation.RelationFollowerListReq) (*relation.RelationFollowerListRespData, error) {
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, errno.CustomError.WithMessage("无效的用户ID").WithInnerError(err)
	}
	req.PageNum, req.PageSize = common.CorrectPageNumAndPageSize(req.PageNum, req.PageSize)
	f := dal.Executor.Follow
	followers, count, err := f.WithContext(context.Background()).
		Where(f.FollowedID.Eq(uid)).
		FindByPage(int(req.PageNum*req.PageSize), int(req.PageSize))
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

	return &relation.RelationFollowerListRespData{
		Items:    *model_converter.UserListDal2Resp(&users),
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
	rows, err := dal.DB.Raw("SELECT followed_id FROM Follow WHERE follower_id = ? AND followed_id IN (SELECT follower_id FROM Follow WHERE followed_id = ?) LIMIT ?, ?", uid, uid, int(req.PageNum*req.PageSize), int(req.PageSize)).Rows()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	defer rows.Close()
	friends := []int64{}
	for rows.Next() {
		var friendId int64
		rows.Scan(&friendId)
		friends = append(friends, friendId)
	}
	row, err := dal.DB.Raw("SELECT COUNT(*) FROM Follow WHERE follower_id = ? AND followed_id IN (SELECT follower_id FROM Follow WHERE followed_id = ?)", uid, uid).Rows()
	if err != nil {
		return nil, errno.DatabaseCallError.WithInnerError(err)
	}
	defer row.Close()
	var count int64
	row.Next()
	row.Scan(&count)

	users := []*model.User{}
	for _, item := range friends {
		user, err := exquery.QueryUserByID(item)
		if err != nil {
			return nil, errno.DatabaseCallError.WithInnerError(err)
		}
		users = append(users, user)
	}

	return &relation.RelationFriendListRespData{
		Items:    *model_converter.UserListDal2Resp(&users),
		IsEnd:    count <= (req.PageNum+1)*req.PageSize,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}
