package db

import (
	"context"
	"gitee.com/derrickball/douyin/app/user/model"
)

func QueryUser(ctx context.Context, userName string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//func QueryUserID(ctx context.Context, uid int64) ([]*model.User, error) {
//	users := make([]*model.User, 0)
//	if err := DB.WithContext(ctx).Where("id = ?", uid).Find(&users).Error; err != nil {
//		return nil, err
//	}
//	return users, nil
//}

func CreateUser(ctx context.Context, user *model.User) error {
	return DB.WithContext(ctx).Create(user).Error
}

func QueryUserByID(ctx context.Context, userID int64) (*model.User, error) {
	var user *model.User
	if err := DB.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
