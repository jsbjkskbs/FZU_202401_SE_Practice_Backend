// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTag = "Tag"

// Tag 标签静态库表
type Tag struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:标签ID" json:"id"` // 标签ID
	TagName   string `gorm:"column:tag_name;not null;comment:标签名称" json:"tag_name"`          // 标签名称
	CreatedAt int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`      // 创建时间
	DeletedAt int64  `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`               // 删除时间
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
