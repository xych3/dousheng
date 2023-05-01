package db

import (
	"context"
	"strconv"

	"gitee.com/derrickball/douyin/app/user/model"
)

func QueryRelation(ctx context.Context, to_user_id int64, follower_id int64) ([]*model.Relation, error) {
	relations := make([]*model.Relation, 0)
	if err := DB.WithContext(ctx).Where(map[string]interface{}{"to_user_id": to_user_id, "follower_id": follower_id}).Find(&relations).Error; err != nil {
		return nil, err
	}
	return relations, nil
}
func CreateRelation(ctx context.Context, relation *model.Relation) error {
	if err := DB.WithContext(ctx).Create(relation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRelation(ctx context.Context, to_user_id int64, follower_id int64) error {
	return DB.WithContext(ctx).Where("to_user_id = ? AND follower_id = ?", to_user_id, follower_id).Delete(&model.Relation{}).Error
}

// query the 粉丝
func QueryRelationFollower(ctx context.Context, to_user_id int64) ([]*model.Relation, error) {
	relations := make([]*model.Relation, 0)
	if err := DB.WithContext(ctx).Where("to_user_id = ?", strconv.FormatInt(to_user_id, 10)).Find(&relations).Error; err != nil {
		return nil, err
	}
	return relations, nil
}

// query the 关注
func QueryRelationFollow(ctx context.Context, follower_id int64) ([]*model.Relation, error) {
	relations := make([]*model.Relation, 0)
	if err := DB.WithContext(ctx).Where("follower_id = ?", strconv.FormatInt(follower_id, 10)).Find(&relations).Error; err != nil {
		return nil, err
	}
	return relations, nil
}

func QueryIsFollow(ctx context.Context, user int64, user_id int64) (bool, error) {
	relations := make([]*model.Relation, 0)
	if err := DB.WithContext(ctx).Where("to_user_id = ? AND follower_id = ?", user, user_id).Find(&relations).Error; err != nil {
		return false, err
	}
	if len(relations) == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
