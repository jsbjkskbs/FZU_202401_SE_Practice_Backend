package gorse

import (
	"context"
	"fmt"
	"time"

	"github.com/zhenghaoz/gorse/client"
)

func InsertVideo(videoId string, category string, labels []string) error {
	_, err := cli.InsertItem(context.Background(), client.Item{
		ItemId:     videoId,
		IsHidden:   true,
		Categories: []string{category, "*"},
		Labels:     labels,
		Timestamp:  fmt.Sprint(time.Unix(time.Now().Unix(), 0).UTC().Format(time.RFC3339)),
	})
	return err
}

func InsertUser(userId string) error {
	_, err := cli.InsertUser(context.Background(), client.User{
		UserId: userId,
	})
	return err
}

func PutVideoLabels(videoId string, labels []string) error {
	_, err := cli.UpdateItem(context.Background(), videoId, client.ItemPatch{
		Labels: labels,
	})
	return err
}

func PutVideoHiddenState(videoId string, hidden bool) error {
	_, err := cli.UpdateItem(context.Background(), videoId, client.ItemPatch{
		IsHidden: &hidden,
	})
	return err
}

func PutVideoCategories(videoId string, categories []string) error {
	_, err := cli.UpdateItem(context.Background(), videoId, client.ItemPatch{
		Categories: categories,
	})
	return err
}

func DelVideo(videoId string) error {
	_, err := cli.DeleteItem(context.Background(), videoId)
	return err
}

func DelUser(userId string) error {
	_, err := cli.DeleteUser(context.Background(), userId)
	return err
}

func PutFeedback(userId, videoId, feedback string) error {
	_, err := cli.InsertFeedback(context.Background(), []client.Feedback{
		{
			FeedbackType: feedback,
			UserId:       userId,
			ItemId:       videoId,
			Timestamp:    fmt.Sprint(time.Unix(time.Now().Unix(), 0).UTC().Format(time.RFC3339)),
		},
	})
	return err
}

func PutFeedbacks(feedbacks []client.Feedback) error {
	_, err := cli.InsertFeedback(context.Background(), feedbacks)
	return err
}

func GetRecommend(userId string, n int) ([]string, error) {
	resp, err := cli.GetRecommend(context.Background(), userId, "*", n)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetRecommendWithCategory(userId, category string, n int) ([]string, error) {
	list, err := cli.GetItemLatestWithCategory(context.Background(), userId, category, n, 0)
	if err != nil {
		return nil, err
	}
	resp := []string{}
	for _, item := range list {
		resp = append(resp, item.Id)
	}
	return resp, nil
}
