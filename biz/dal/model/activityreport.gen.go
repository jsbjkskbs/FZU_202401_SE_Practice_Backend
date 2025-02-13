// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameActivityReport = "ActivityReport"

// ActivityReport 动态举报表
type ActivityReport struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:举报ID" json:"id"` // 举报ID
	UserID     int64  `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`            // 用户ID
	ActivityID int64  `gorm:"column:activity_id;not null;comment:动态ID" json:"activity_id"`    // 动态ID
	Reason     string `gorm:"column:reason;not null;comment:举报原因" json:"reason"`              // 举报原因
	Label      string `gorm:"column:label;not null;comment:举报标签" json:"label"`                // 举报标签
	CreatedAt  int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`      // 创建时间
	Status     string `gorm:"column:status;not null;comment:举报状态" json:"status"`              // 举报状态
	ResolvedAt int64  `gorm:"column:resolved_at;comment:解决时间" json:"resolved_at"`             // 解决时间
	AdminID    int64  `gorm:"column:admin_id;not null;comment:管理员ID" json:"admin_id"`         // 管理员ID
}

// TableName ActivityReport's table name
func (*ActivityReport) TableName() string {
	return TableNameActivityReport
}
