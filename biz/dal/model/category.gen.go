// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCategory = "Category"

// Category 分类静态库表
type Category struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:分区ID" json:"id"`  // 分区ID
	CategoryName string `gorm:"column:category_name;not null;comment:分区名称" json:"category_name"` // 分区名称
	CreatedAt    int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`       // 创建时间
	DeletedAt    int64  `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                // 删除时间
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}