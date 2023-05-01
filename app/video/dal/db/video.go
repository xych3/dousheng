package db

import (
	"context"
	"gitee.com/derrickball/douyin/app/video/model"
	"gorm.io/gorm"
	"time"
)

func CreateVideo(ctx context.Context, video *model.Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func QueryVideo(ctx context.Context, latestTime int64, limit int) ([]*model.Video, error) {
	videos := make([]*model.Video, 0)
	err := DB.WithContext(ctx).Where("created_at < ?", time.Unix(latestTime, 0)).Limit(limit).Order("created_at desc").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func QueryVideoByUserID(ctx context.Context, userID int64) ([]*model.Video, error) {
	videos := make([]*model.Video, 0)
	err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&videos).Order("created_at desc").Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func QueryVideoByIdBatch(ctx context.Context, idList []int64) ([]*model.Video, error) {
	videos := make([]*model.Video, 0)
	DB.WithContext(ctx).Where("id IN ?", idList).Find(&videos)
	return videos, nil
}

func IncreaseFavoriteCountById(ctx context.Context, videoId int64) {
	video := &model.Video{ID: videoId}
	DB.WithContext(ctx).Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
}

func DecreaseFavoriteCountById(ctx context.Context, videoId int64) {
	video := &model.Video{ID: videoId}
	DB.WithContext(ctx).Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
}

func IncreaseCommentCountById(ctx context.Context, videoId int64) {
	video := &model.Video{ID: videoId}
	DB.WithContext(ctx).Model(&video).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
}

func DecreaseCommentCountById(ctx context.Context, videoId int64) {
	video := &model.Video{ID: videoId}
	DB.WithContext(ctx).Model(&video).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
}

func GetAuthorId(ctx context.Context, videoId int64) int64 {
	video := &model.Video{ID: videoId}
	DB.WithContext(ctx).First(video)
	return video.UserID
}
