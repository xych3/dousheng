package db

import (
	"context"
	"gitee.com/derrickball/douyin/app/user/model"
	"gorm.io/gorm"
)

func IsFollower(ctx context.Context, toUserID int64, followerID int64) (bool, error) {
	relation := &model.Relation{}
	err := DB.WithContext(ctx).Where("to_user_id = ? AND follower_id = ?", toUserID, followerID).First(relation).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
