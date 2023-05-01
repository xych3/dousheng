package db

import (
	"context"
	"fmt"
	"time"

	chat_model "gitee.com/derrickball/douyin/app/chat/model"
	relation_model "gitee.com/derrickball/douyin/app/relation/model"
)

// 查询好友列表
//func FriendList(ctx context.Context, userId int64) ([]*user_model.User, error) {
//	users := make([]*user_model.User, 0)
//	if len(fmt.Sprint(userId)) != 0 {
//		// 查询与该用户互相关注的好友
//		subQuary := DB.WithContext(ctx).
//			Select("follower_id").
//			Where("to_user_id = ?", userId).Table("relations")
//		err := subQuary.Error
//		if err != nil {
//			return nil, err
//		}
//		friendsID := DB.WithContext(ctx).
//			Select("to_user_id").
//			Where("follower_id = ? AND to_user_id IN ?", userId, subQuary).
//			Table("relations")
//		err = friendsID.Error
//		if err != nil {
//			return nil, err
//		}
//		// 查询好友列表用户
//		result := DB.WithContext(ctx).
//			Where("id IN ?", friendsID).Find(&users)
//		err = result.Error
//		if err != nil {
//			return nil, err
//		}
//
//	}
//	return users, nil
//}

// FollowList 查询关注列表
func FollowList(ctx context.Context, userId int64) []*relation_model.Relation {
	data := make([]*relation_model.Relation, 0)
	DB.WithContext(ctx).Where("follower_id = ?", userId).Find(&data)
	return data

}

// FollowerList 查询粉丝列表
func FollowerList(ctx context.Context, userId int64) []*relation_model.Relation {
	data := make([]*relation_model.Relation, 0)
	DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&data)
	return data
}

// 查询是否关注
func QueryIsFollow(ctx context.Context, to_user int64, follower int64) (bool, error) {
	relations := make([]*relation_model.Relation, 0)
	if err := DB.WithContext(ctx).Where("to_user_id = ? AND follower_id = ?", to_user, follower).Find(&relations).Error; err != nil {
		return false, err
	}
	if len(relations) == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

// 查询聊天记录信息
func QueryMessage(ctx context.Context, toUserId int64, fromUserId int64) ([]*chat_model.Chat, error) {
	messages := make([]*chat_model.Chat, 0)
	if len(fmt.Sprint(toUserId)) != 0 || len(fmt.Sprint(fromUserId)) != 0 {
		if err := DB.WithContext(ctx).
			Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).
			Or("from_user_id = ? AND to_user_id = ?", toUserId, fromUserId).
			Order("id DESC").Find(&messages).Error; err != nil {
			return nil, err
		}
	}
	return messages, nil
}

// 创建聊天信息
func CreateMessage(ctx context.Context, msg *chat_model.Chat) (int64, error) {
	db := DB.WithContext(ctx)
	msg.CreatedAt = time.Now()
	if err := db.Create(msg).Error; err != nil {
		return -1, err
	}
	var id int64
	// 返回这条创建的消息的id
	if err := db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id).Error; err != nil {
		return -1, err
	}
	return id, nil
}
