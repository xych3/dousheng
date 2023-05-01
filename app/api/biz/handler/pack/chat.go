package pack

import (
	"gitee.com/derrickball/douyin/app/api/biz/model/chat"
	"gitee.com/derrickball/douyin/app/chat/kitex_gen/chat_rpc"
)

func Friend(u *chat_rpc.RPCFriendUser) *chat.FriendUser {
	return &chat.FriendUser{
		Id:            &u.Id,
		Name:          &u.Name,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      &u.IsFollow,
		Avatar:        &u.Avatar,
		Message:       &u.Message,
		MsgType:       &u.MsgType,
	}
}

func UserList(usrl []*chat_rpc.RPCFriendUser) []*chat.FriendUser {
	userList := make([]*chat.FriendUser, 0)
	for _, f := range usrl {
		if msg := Friend(f); msg != nil {
			userList = append(userList, msg)
		}
	}
	return userList
}

func Message(m *chat_rpc.RPCMessage) *chat.Message {
	return &chat.Message{
		Id:         &m.Id,
		ToUserId:   &m.ToUserId,
		FromUserId: &m.FromUserId,
		Content:    &m.Content,
		CreateTime: &m.CreateTime,
	}
}

func MessageList(msgl []*chat_rpc.RPCMessage) []*chat.Message {
	msgList := make([]*chat.Message, 0)
	for _, f := range msgl {
		if msg := Message(f); msg != nil {
			msgList = append(msgList, msg)
		}
	}
	return msgList
}
