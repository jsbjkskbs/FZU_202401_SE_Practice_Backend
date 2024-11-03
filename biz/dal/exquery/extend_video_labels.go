package exquery

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func InsertVideoLabel(videoLabels ...*model.VideoLabel) error {
	vl := dal.Executor.VideoLabel
	err := vl.WithContext(context.Background()).Create(videoLabels...)
	if err != nil {
		return err
	}
	return nil
}
