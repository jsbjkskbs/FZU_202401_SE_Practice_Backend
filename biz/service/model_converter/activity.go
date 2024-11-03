package model_converter

import (
	"fmt"
	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
	"sfw/biz/mw/redis"
	"sfw/pkg/oss"
)

func ActivityListDal2Resp(list *[]*model.Activity) (*[]*base.Activity, error) {
	resp := &[]*base.Activity{}
	for _, v := range *list {
		images := []string{}
		dal.DB.Raw(
			`SELECT i.image_url  	
			FROM Image i  
			JOIN ActivityImages ai ON i.id = ai.image_id  
			WHERE ai.activity_id = ?;`, v.ID).Scan(&images)
		for i, image := range images {
			images[i] = oss.Key2Url(image)
		}

		refVideoID := ""
		refActivityID := ""
		if v.RefVideoID != 0 {
			refVideoID = fmt.Sprint(v.RefVideoID)
		}
		if v.RefActivityID != 0 {
			refActivityID = fmt.Sprint(v.RefActivityID)
		}

		likeCount, err := redis.GetActivityLikeCount(fmt.Sprint(v.ID))
		if err != nil {
			return nil, err
		}

		*resp = append(*resp, &base.Activity{
			ID:          fmt.Sprint(v.ID),
			UserID:      fmt.Sprint(v.UserID),
			Content:     v.Content,
			Image:       images,
			RefVideo:    refVideoID,
			RefActivity: refActivityID,
			LikeCount:   likeCount,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
		})
	}
	return resp, nil
}
