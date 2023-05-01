// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVideo = "videos"

// Video mapped from table <videos>
type Video struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键ID
	Title         string         `gorm:"column:title" json:"title"`                         // 视频标题
	UserID        int64          `gorm:"column:user_id" json:"user_id"`                     // 视频作者ID
	PlayURL       string         `gorm:"column:play_url" json:"play_url"`                   // 视频播放地址
	CoverURL      string         `gorm:"column:cover_url" json:"cover_url"`                 // 视频封面地址
	FavoriteCount int64          `gorm:"column:favorite_count" json:"favorite_count"`       // 视频点赞总数
	CommentCount  int64          `gorm:"column:comment_count" json:"comment_count"`         // 视频评论总数
	CreatedAt     time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
