// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVideoCommentReport = "VideoCommentReport"

// VideoCommentReport 视频评论举报表
type VideoCommentReport struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:举报ID" json:"id"` // 举报ID
	UserID     int64  `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`            // 用户ID
	CommentID  int64  `gorm:"column:comment_id;not null;comment:视频评论ID" json:"comment_id"`    // 视频评论ID
	Reason     string `gorm:"column:reason;not null;comment:举报原因" json:"reason"`              // 举报原因
	Label      string `gorm:"column:label;not null;comment:举报标签" json:"label"`                // 举报标签
	CreatedAt  int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`      // 创建时间
	Status     string `gorm:"column:status;not null;comment:举报状态" json:"status"`              // 举报状态
	ResolvedAt int64  `gorm:"column:resolved_at;comment:解决时间" json:"resolved_at"`             // 解决时间
	AdminID    int64  `gorm:"column:admin_id;not null;comment:管理员ID" json:"admin_id"`         // 管理员ID
}

// TableName VideoCommentReport's table name
func (*VideoCommentReport) TableName() string {
	return TableNameVideoCommentReport
}
