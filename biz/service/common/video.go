package common

import (
	"context"
	"sfw/biz/dal"
	"sfw/biz/model/base"
	"sfw/biz/service/model_converter"
	"sfw/pkg/errno"
	"strconv"

	"gorm.io/gen"
)

/*

	用于节省代码块，将相同的代码块提取到common.go中

*/

func QueryVideoSubmit(userId, status string, pageNum, pageSize int64) (*[]*base.Video, int64, error) {
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, 0, errno.InternalServerError
	}

	v := dal.Executor.Video

	conditions := []gen.Condition{v.UserID.Eq(uid)}
	if status != "" {
		conditions = append(conditions, v.Status.Eq(status))
	}

	result, count, err := v.WithContext(context.Background()).
		Where(conditions...).
		FindByPage(int(pageNum), int(pageSize))

	videos, err := model_converter.VideoListDal2Resp(&result)
	if err != nil {
		return nil, 0, err
	}

	return &videos, count, nil
}
