package db

import (
	"context"
	"fmt"
	"time"

	"gitee.com/derrickball/douyin/app/comment/model"
)

func DelComments(ctx context.Context, commentId int64) (string, string, error) {
	db := DB.WithContext(ctx)
	comm := model.Comment{}
	if err := db.Find(&comm, commentId).Error; err != nil {
		return "", "", err
	}
	createTime := comm.CreatedAt.Format("2006-01-02 15:04:05")
	content := comm.CommentText
	if err := db.Delete(&comm).Error; err != nil {
		return "", "", err
	}
	return content, createTime, nil
}

func QueryComments(ctx context.Context, videoId int64) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0)
	if len(fmt.Sprint(videoId)) != 0 {
		if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&comments).Error; err != nil {
			return nil, err
		}
	}
	return comments, nil
}

func CreateComment(ctx context.Context, com *model.Comment) error {
	db := DB.WithContext(ctx)
	// 保证在外面能获取到创建时间
	com.CreatedAt = time.Now()
	if err := db.Create(com).Error; err != nil {
		return err
	}
	return nil
}
