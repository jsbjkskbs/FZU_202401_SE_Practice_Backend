package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryFollowerCountByUserID(uid int64) (int64, error) {
	f := dal.Executor.Follow
	count, err := f.WithContext(context.Background()).Where(f.FollowedID.Eq(uid)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryFollowingCountByUserID(uid int64) (int64, error) {
	f := dal.Executor.Follow
	count, err := f.WithContext(context.Background()).Where(f.FollowerID.Eq(uid)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryFollowExistByFollowerIDAndFollowedID(followerID, followedID int64) (bool, error) {
	f := dal.Executor.Follow
	count, err := f.WithContext(context.Background()).Where(f.FollowerID.Eq(followerID), f.FollowedID.Eq(followedID)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryFollowingByUserIdPaged(uid, pageNum, pageSize int64) ([]*model.Follow, int64, error) {
	f := dal.Executor.Follow
	follows, count, err := f.WithContext(context.Background()).Where(f.FollowerID.Eq(uid)).FindByPage(int(pageNum*pageSize), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return follows, count, nil
}

func QueryFollowerByUserIdPaged(uid, pageNum, pageSize int64) ([]*model.Follow, int64, error) {
	f := dal.Executor.Follow
	followers, count, err := f.WithContext(context.Background()).Where(f.FollowedID.Eq(uid)).FindByPage(int(pageNum*pageSize), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return followers, count, nil
}

func QueryFriendByUserIDPaged(uid, pageNum, pageSize int64) ([]int64, int64, error) {
	rows, err := dal.DB.Raw("SELECT followed_id FROM Follow WHERE follower_id = ? AND followed_id IN (SELECT follower_id FROM Follow WHERE followed_id = ?) LIMIT ?, ?", uid, uid, pageNum*pageSize, pageSize).Rows()
	if err != nil {
		return nil, 0, err
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
		return nil, 0, err
	}
	defer row.Close()
	var count int64
	row.Next()
	row.Scan(&count)
	return friends, count, nil
}

func InsertFollowRecord(follows ...*model.Follow) error {
	f := dal.Executor.Follow
	err := f.WithContext(context.Background()).Create(follows...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFollowRecord(followerID, followedID int64) error {
	f := dal.Executor.Follow
	_, err := f.WithContext(context.Background()).Where(f.FollowerID.Eq(followerID), f.FollowedID.Eq(followedID)).Delete()
	if err != nil {
		return err
	}
	return nil
}
