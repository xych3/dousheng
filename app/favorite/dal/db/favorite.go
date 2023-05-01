package db

import (
	"context"
	"gitee.com/derrickball/douyin/app/favorite/model"
)

func CreateFavorite(ctx context.Context, favorite *model.Favorite) error {
	// 增加判断点赞视频是否存在
	//var video *model.Video
	//if err := DB.WithContext(ctx).Where("id = ?", favorite.VideoID).Find(&video).Error; err != nil {
	//	return err
	//}
	return DB.WithContext(ctx).Create(favorite).Error
}

func DeleteFavorite(ctx context.Context, userID int64, videoID int64) error {
	return DB.
		WithContext(ctx).
		Where("user_id = ? and video_id = ? ", userID, videoID).
		Delete(&model.Favorite{}).Error
}

func QueryFavorite(ctx context.Context, userID int64) ([]*model.Favorite, error) {
	var res []*model.Favorite
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func FilterFavorite(ctx context.Context, userID int64, videoID []int64) ([]*model.Favorite, error) {
	favorites := make([]*model.Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id IN ?", userID, videoID).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}
