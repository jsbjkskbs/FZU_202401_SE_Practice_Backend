// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVideoLabel = "VideoLabel"

// VideoLabel 视频标签关联表
type VideoLabel struct {
	VideoID   int64  `gorm:"column:video_id;primaryKey;comment:视频ID" json:"video_id"`     // 视频ID
	LabelName string `gorm:"column:label_name;primaryKey;comment:标签ID" json:"label_name"` // 标签ID
	CreatedAt int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`   // 创建时间
	DeletedAt int64  `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`            // 删除时间
}

// TableName VideoLabel's table name
func (*VideoLabel) TableName() string {
	return TableNameVideoLabel
}
