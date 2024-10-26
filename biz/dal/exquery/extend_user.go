package exquery

import "sfw/biz/dal"

/*

	主要用于扩展数据访问操作

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
