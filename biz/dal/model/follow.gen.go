// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFollow = "Follow"

// Follow 关注关系表
type Follow struct {
	FollowedID int64 `gorm:"column:followed_id;primaryKey;comment:被关注者ID" json:"followed_id"` // 被关注者ID
	FollowerID int64 `gorm:"column:follower_id;primaryKey;comment:关注者ID" json:"follower_id"`  // 关注者ID
	CreatedAt  int64 `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`       // 创建时间
	DeletedAt  int64 `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                // 删除时间
}

// TableName Follow's table name
func (*Follow) TableName() string {
	return TableNameFollow
}
