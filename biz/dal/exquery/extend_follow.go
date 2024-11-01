package exquery

import (
	"context"
	"sfw/biz/dal"
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
