package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func InsertActivityImage(activityImage ...*model.ActivityImage) error {
	a := dal.Executor.ActivityImage
	err := a.WithContext(context.Background()).Create(activityImage...)
	if err != nil {
		return err
	}
	return nil
}
