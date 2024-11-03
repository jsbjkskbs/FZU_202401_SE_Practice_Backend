package model_converter

import (
	"context"
	"fmt"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
)

func ActivityListDal2Resp(list *[]*model.Activity) *[]*base.Activity {
	resp := &[]*base.Activity{}
	for _, v := range *list {
		ai := dal.Executor.ActivityImage
		activityImgData, err := ai.WithContext(context.Background()).Where(ai.ActivityID.Eq(v.ID)).Find()
		if err != nil {
			continue
		}
		images := []string{}
		for _, img := range activityImgData {
			images = append(images, fmt.Sprint(img.ImageID))
		}
		refVideoID := ""
		refActivityID := ""
		if v.RefVideoID != 0 {
			refVideoID = fmt.Sprint(v.RefVideoID)
		}
		if v.RefActivityID != 0 {
			refActivityID = fmt.Sprint(v.RefActivityID)
		}

		*resp = append(*resp, &base.Activity{
			ID:          fmt.Sprint(v.ID),
			UserID:      fmt.Sprint(v.UserID),
			Content:     v.Content,
			Image:       images,
			RefVideo:    refVideoID,
			RefActivity: refActivityID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
		})
	}
	return resp
}
