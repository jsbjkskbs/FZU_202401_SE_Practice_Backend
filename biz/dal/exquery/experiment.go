package exquery

import (
	"sfw/biz/dal"
	"sfw/biz/dal/model"
)

/*

	might be changed if you want to restructuring the project as a distributed system
	because the query is not only related to only one table,
	and these tables is desperately stored in different databases

*/

func QueryUserLikeCount(id int64) (int64, error) {
	var sum int64
	if err := dal.DB.
		Raw("select count(*) from Video, VideoLike where Video.id = VideoLike.video_id and Video.user_id = ?", id).
		Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}

func QueryVideoLikedByUserIdPaged(userId int64, pageNum, pageSize int) ([]*model.Video, int64, error) {
	rows, err := dal.DB.Raw(
		`SELECT v.*  
		FROM Video v  
		JOIN (  
			SELECT video_id, created_at  
			FROM VideoLike  
			WHERE user_id = ?  
			ORDER BY created_at DESC  
			LIMIT ?, ?  
		) vl ON v.id = vl.video_id;`,
		userId, pageNum*pageSize, pageSize,
	).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	row, err := dal.DB.Raw(
		`SELECT COUNT(*) FROM Video WHERE id IN (SELECT video_id FROM VideoLike WHERE user_id = ?)`,
		userId,
	).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer row.Close()

	videos := make([]*model.Video, 0, pageSize)
	for rows.Next() {
		var video model.Video
		dal.DB.ScanRows(rows, &video)
		videos = append(videos, &video)
	}

	var count int64
	row.Next()
	row.Scan(&count)

	return videos, count, nil
}

func QueryImageUrlsByActivityId(activityId int64) ([]string, error) {
	images := []string{}
	result := dal.DB.Raw(
		`SELECT i.image_url  	
		FROM Image i  
		JOIN ActivityImages ai ON i.id = ai.image_id  
		WHERE ai.activity_id = ?;`, activityId).Scan(&images)
	if result.Error != nil {
		return nil, result.Error
	}
	return images, nil
}
