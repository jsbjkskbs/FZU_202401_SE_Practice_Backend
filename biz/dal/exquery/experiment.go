package exquery

import "sfw/biz/dal"

/*

	might be changed if you want to restructuring the project as a distributed system
	because the query is not only related to only one table, 
	and these tables is desperately stored in different databases

*/

func QueryUserLikeCount(id int64) (int64, error) {
	var sum int64
	if err :=
		dal.DB.
			Raw("select count(*) from Video, VideoLike where Video.id = VideoLike.video_id and Video.user_id = ?", id).
			Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}
