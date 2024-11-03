package model_converter

import (
	"fmt"
	"sfw/biz/dal/model"
	"sfw/biz/model/base"
)

func VideoReportDal2Resp(list *[]*model.VideoReport) (resp []*base.VideoReport) {
	for _, item := range *list {
		resp = append(resp, &base.VideoReport{
			ID:         fmt.Sprint(item.ID),
			UserID:     fmt.Sprint(item.UserID),
			VideoID:    fmt.Sprint(item.VideoID),
			Reason:     item.Reason,
			Label:      item.Label,
			Status:     item.Status,
			CreatedAt:  item.CreatedAt,
			ResolvedAt: item.ResolvedAt,
			AdminID:    fmt.Sprint(item.AdminID),
		})
	}
	return
}

func ActivityReportDal2Resp(list *[]*model.ActivityReport) (resp []*base.ActivityReport) {
	for _, item := range *list {
		resp = append(resp, &base.ActivityReport{
			ID:         fmt.Sprint(item.ID),
			UserID:     fmt.Sprint(item.UserID),
			ActivityID: fmt.Sprint(item.ActivityID),
			Reason:     item.Reason,
			Label:      item.Label,
			Status:     item.Status,
			CreatedAt:  item.CreatedAt,
			ResolvedAt: item.ResolvedAt,
			AdminID:    fmt.Sprint(item.AdminID),
		})
	}
	return
}

func VideoCommentReportDal2Resp(list *[]*model.VideoCommentReport) (resp []*base.CommentReport) {
	for _, item := range *list {
		resp = append(resp, &base.CommentReport{
			ID:          fmt.Sprint(item.ID),
			UserID:      fmt.Sprint(item.UserID),
			CommentID:   fmt.Sprint(item.CommentID),
			CommentType: "video",
			Reason:      item.Reason,
			Label:       item.Label,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			ResolvedAt:  item.ResolvedAt,
			AdminID:     fmt.Sprint(item.AdminID),
		})
	}
	return
}

func ActivityCommentReportDal2Resp(list *[]*model.ActivityCommentReport) (resp []*base.CommentReport) {
	for _, item := range *list {
		resp = append(resp, &base.CommentReport{
			ID:          fmt.Sprint(item.ID),
			UserID:      fmt.Sprint(item.UserID),
			CommentID:   fmt.Sprint(item.CommentID),
			CommentType: "activity",
			Reason:      item.Reason,
			Label:       item.Label,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			ResolvedAt:  item.ResolvedAt,
			AdminID:     fmt.Sprint(item.AdminID),
		})
	}
	return
}
