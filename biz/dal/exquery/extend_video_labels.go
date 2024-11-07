package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryVideoLabels(videoId int64) ([]string, error) {
	labelItems := []model.VideoLabel{}
	l := dal.Executor.VideoLabel
	err := l.WithContext(context.Background()).Where(l.VideoID.Eq(videoId)).Select(l.LabelName).Scan(&labelItems)
	if err != nil {
		return nil, err
	}
	labels := []string{}
	for _, item := range labelItems {
		labels = append(labels, item.LabelName)
	}
	return labels, nil
}

func InsertVideoLabel(videoLabels ...*model.VideoLabel) error {
	vl := dal.Executor.VideoLabel
	err := vl.WithContext(context.Background()).Create(videoLabels...)
	if err != nil {
		return err
	}
	return nil
}
