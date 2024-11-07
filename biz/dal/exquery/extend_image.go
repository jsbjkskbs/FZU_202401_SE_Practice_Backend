package exquery

import (
	"context"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

func QueryImageExistById(imageId int64) (bool, error) {
	img := dal.Executor.Image
	exist, err := img.WithContext(context.Background()).Where(img.ID.Eq(imageId)).Count()
	if err != nil {
		return false, err
	}
	return exist > 0, nil
}

func QueryImageById(imageId int64) (*model.Image, error) {
	img := dal.Executor.Image
	images, _, err := img.WithContext(context.Background()).Where(img.ID.Eq(imageId)).FindByPage(0, 1)
	if err != nil {
		return nil, err
	}
	if len(images) == 0 {
		return nil, nil
	}
	return images[0], nil
}

func QueryImageUrlById(id int64) (string, error) {
	img := dal.Executor.Image
	images, _, err := img.WithContext(context.Background()).Where(img.ID.Eq(id)).Select(img.ImageURL).FindByPage(0, 1)
	if err != nil {
		return "", err
	}
	if len(images) == 0 {
		return "", nil
	}
	return images[0].ImageURL, nil
}

func InsertImage(images ...*model.Image) error {
	img := dal.Executor.Image
	err := img.WithContext(context.Background()).Create(images...)
	if err != nil {
		return err
	}
	return nil
}

func UpdateImageWithId(image *model.Image) error {
	img := dal.Executor.Image
	_, err := img.WithContext(context.Background()).Where(img.ID.Eq(image.ID)).Updates(image)
	if err != nil {
		return err
	}
	return nil
}
