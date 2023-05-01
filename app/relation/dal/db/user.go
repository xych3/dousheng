package db

import (
	"context"

	"gitee.com/derrickball/douyin/app/user/model"
	"gorm.io/gorm"
)

func QueryUser(ctx context.Context, userName string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(ctx context.Context, user *model.User) error {
	return DB.WithContext(ctx).Create(user).Error
}

func QueryUserById(ctx context.Context, users_id []int64) ([]*model.User, error) {
	users := make([]*model.User, 0)
	result := DB.WithContext(ctx).Where(users_id).Find(&users)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateFollowCount(ctx context.Context, to_user_id int64, follower_id int64, flag bool) error {
	if flag {
		if err := DB.Model(&model.User{}).Where("id = ?", to_user_id).Update("follower_Count", gorm.Expr("follower_count+?", 1)).Error; err != nil {
			return err
		}
		if err := DB.Model(&model.User{}).Where("id = ?", follower_id).Update("follow_count", gorm.Expr("follow_count+?", 1)).Error; err != nil {
			return err
		}
	} else {
		if err := DB.Model(&model.User{}).Where("id = ?", to_user_id).Update("follower_Count", gorm.Expr("follower_count-?", 1)).Error; err != nil {
			return err
		}
		if err := DB.Model(&model.User{}).Where("id = ?", follower_id).Update("follow_count", gorm.Expr("follow_count-?", 1)).Error; err != nil {
			return err
		}
	}
	return nil
}
