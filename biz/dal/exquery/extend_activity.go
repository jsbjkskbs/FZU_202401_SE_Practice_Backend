package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryActivityExistByIdAndUserId(id, userId int64) (bool, error) {
	a := dal.Executor.Activity
	count, err := a.WithContext(context.Background()).Where(a.ID.Eq(id), a.UserID.Eq(userId)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryActivityByFollowedIdPaged(followerId int64, pageNum, pageSize int) ([]*model.Activity, int64, error) {
	rows, err := dal.DB.Raw("SELECT * FROM Activity WHERE user_id IN (SELECT followed_id FROM Follow WHERE follower_id = ?) OR user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", followerId, followerId, pageSize, pageNum*pageSize).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	activities := []*model.Activity{}
	for rows.Next() {
		var activity model.Activity
		err = dal.DB.ScanRows(rows, &activity)
		if err != nil {
			return nil, 0, err
		}
		activities = append(activities, &activity)
	}
	row, err := dal.DB.Raw("SELECT COUNT(*) FROM Activity WHERE user_id IN (SELECT followed_id FROM Follow WHERE follower_id = ?)", followerId).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer row.Close()
	var count int64
	row.Next()
	row.Scan(&count)

	return activities, count, nil
}

func QueryActivityByUserIdPaged(userId int64, pageNum, pageSize int) ([]*model.Activity, int64, error) {
	a := dal.Executor.Activity
	activities, count, err := a.WithContext(context.Background()).
		Where(a.UserID.Eq(userId)).
		FindByPage(int(pageNum*pageSize), int(pageSize))
	if err != nil {
		return nil, 0, err
	}
	return activities, count, nil
}

func QueryActivityExistById(id int64) (bool, error) {
	a := dal.Executor.Activity
	count, err := a.WithContext(context.Background()).Where(a.ID.Eq(id)).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertActivity(activity *model.Activity) error {
	a := dal.Executor.Activity
	err := a.WithContext(context.Background()).Create(activity)
	if err != nil {
		return err
	}
	return nil
}

func DeleteActivityById(id int64) error {
	a := dal.Executor.Activity
	_, err := a.WithContext(context.Background()).Where(a.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}
