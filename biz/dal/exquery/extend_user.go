package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

/*

	主要用于扩展数据访问操作

*/

func QueryUserByID(id int64) (*model.User, error) {
	u := dal.Executor.User
	uc := u.WithContext(context.Background())
	user, err := uc.Where(uc.Where(u.ID.Eq(id))).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func QueryUserByUsername(username string) (*model.User, error) {
	u := dal.Executor.User
	uc := u.WithContext(context.Background())
	user, err := uc.Where(uc.Where(u.Username.Eq(username))).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func QueryUserByEmail(email string) (*model.User, error) {
	u := dal.Executor.User
	uc := u.WithContext(context.Background())
	user, err := uc.Where(uc.Where(u.Email.Eq(email))).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func QueryUserExistByID(id int64) (bool, error) {
	u := dal.Executor.User
	uc := u.WithContext(context.Background())
	count, err := uc.Where(uc.Where(u.ID.Eq(id))).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func QueryUserExistByUsername(username string) (bool, error) {
	u := dal.Executor.User
	uc := u.WithContext(context.Background())
	count, err := uc.Where(uc.Where(u.Username.Eq(username))).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func QueryUserExistByEmail(email string) (bool, error) {
	u := dal.Executor.User
	uc := u.WithContext(context.Background())
	count, err := uc.Where(uc.Where(u.Email.Eq(email))).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func QueryUserExistByUsernameOrEmail(username, email string) (bool, error) {
	u := dal.Executor.User
	uc := u.WithContext(context.Background())
	count, err := uc.
		Where(
			uc.Where(u.Username.Eq(username)).
				Or(uc.Where(u.Email.Eq(email))),
		).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func InsertUser(user *model.User) error {
	u := dal.Executor.User
	err := u.WithContext(context.Background()).Create(user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserWithId(user *model.User) error {
	u := dal.Executor.User
	_, err := u.WithContext(context.Background()).Where(u.ID.Eq(user.ID)).Updates(user)
	if err != nil {
		return err
	}
	return nil
}
