// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameComment = "comments"

// Comment mapped from table <comments>
type Comment struct {
	ID          int64          `gorm:"column:id;primaryKey" json:"id"`          // 主键ID
	UserID      int64          `gorm:"column:user_id" json:"user_id"`           // 用户ID
	VideoID     int64          `gorm:"column:video_id" json:"video_id"`         // 视频ID
	CommentText string         `gorm:"column:comment_text" json:"comment_text"` // 评论内容
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
