package main

import (
	"context"
	"fmt"
	"time"

	"gitee.com/derrickball/douyin/app/chat/rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"

	"gitee.com/derrickball/douyin/app/chat/dal/db"
	"gitee.com/derrickball/douyin/app/chat/kitex_gen/chat_rpc"
	"gitee.com/derrickball/douyin/app/chat/model"
	"gitee.com/derrickball/douyin/pkg/errno"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

func BaseResp(err error) *chat_rpc.BaseResp {
	errNo := errno.ConvertErr(err)
	return &chat_rpc.BaseResp{
		StatusCode:    errNo.ErrCode,
		StatusMessage: errNo.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}

// FriendList implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) FriendList(ctx context.Context, req *chat_rpc.RPCFriendListReq) (resp *chat_rpc.RPCFriendListResp, err error) {
	resp = new(chat_rpc.RPCFriendListResp)
	if len(fmt.Sprint(req.GetUserId())) == 0 {
		resp.BaseResp = BaseResp(errno.ParamErr)
		return resp, nil
	}

	followList := db.FollowList(ctx, req.UserId)
	followerList := db.FollowerList(ctx, req.UserId)
	var friendUserIdList []int64
	// 数组求交集
	mp := map[int64]struct{}{}
	for _, v := range followList {
		mp[v.ToUserID] = struct{}{}
	}
	for _, v := range followerList {
		if _, ok := mp[v.FollowerID]; ok {
			friendUserIdList = append(friendUserIdList, v.FollowerID)
			delete(mp, v.FollowerID)
		}
	}
	FriendUserList := make([]*chat_rpc.RPCFriendUser, 0)
	for _, friendUserId := range friendUserIdList {
		queryUser := rpc.QueryUser(ctx, &user_rpc.RPCQueryUserReq{
			MyUserId:     req.UserId,
			TargetUserId: friendUserId,
		})
		// 查询最新一条聊天记录
		messages, err := db.QueryMessage(ctx, friendUserId, req.UserId)
		if err != nil {
			resp.BaseResp = BaseResp(err)
			return resp, nil
		}
		// 若当前最新消息的发送者为当前用户，则返回1；否则返回0
		var msgtype int64 = 0
		if messages[0].FromUserID == req.UserId {
			msgtype = 1
		}
		// 查询是否关注
		isfollow, err := db.QueryIsFollow(ctx, friendUserId, req.UserId)
		if err != nil {
			resp.BaseResp = BaseResp(err)
			return resp, nil
		}
		// 增加好友列表信息
		FriendUserList = append(FriendUserList, &chat_rpc.RPCFriendUser{
			Id:            queryUser.Id,
			Name:          queryUser.Name,
			FollowCount:   queryUser.FollowCount,
			FollowerCount: queryUser.FollowerCount,
			IsFollow:      isfollow,
			Avatar:        "https://img.zcool.cn/community/01787d5afe768ea801218cf41cf7eb.jpg",
			Message:       messages[0].Content,
			MsgType:       msgtype,
		})
	}

	//FriendUserList := make([]*chat_rpc.RPCFriendUser, 0)
	//for _, usr := range frienduser {
	//	FriendUserList = append(FriendUserList, &chat_rpc.RPCFriendUser{
	//		Id:            usr.ID,
	//		Name:          usr.Username,
	//		FollowCount:   usr.FollowCount,
	//		FollowerCount: usr.FollowerCount,
	//		IsFollow:      true,
	//		Avatar:        "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
	//		Message:       "你好",
	//		MsgType:       1,
	//	})
	//}
	resp.BaseResp = BaseResp(errno.Success)
	resp.UserList = FriendUserList
	return resp, nil
}

// CreateMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) CreateMessage(ctx context.Context, req *chat_rpc.RPCMessageCreateReq) (resp *chat_rpc.RPCMessageCreateResp, err error) {
	resp = new(chat_rpc.RPCMessageCreateResp)
	if len(fmt.Sprint(req.GetToUserId())) == 0 || len(fmt.Sprint(req.GetFromUserId())) == 0 {
		resp.BaseResp = BaseResp(errno.ParamErr)
		return resp, nil
	}
	//mysql
	message := &model.Chat{
		FromUserID: req.GetFromUserId(),
		ToUserID:   req.GetToUserId(),
		Content:    req.GetContent(),
	}

	messageId, err := db.CreateMessage(ctx, message)
	if err != nil {
		resp.BaseResp = BaseResp(err)
		return resp, nil
	}

	resp.Message = &chat_rpc.RPCMessage{
		Id:         messageId,
		ToUserId:   req.ToUserId,
		FromUserId: req.FromUserId,
		Content:    req.Content,
		// CreateTime: message.CreatedAt.String(),
		CreateTime: message.CreatedAt.Unix(),
	}
	resp.BaseResp = BaseResp(errno.Success)
	return resp, nil
}

// MessageList implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) MessageList(ctx context.Context, req *chat_rpc.RPCMessageChatReq) (resp *chat_rpc.RPCMessageChatResp, err error) {
	// TODO: Your code here...
	resp = new(chat_rpc.RPCMessageChatResp)
	if len(fmt.Sprint(req.GetToUserId())) == 0 {
		resp.BaseResp = BaseResp(errno.ParamErr)
		return resp, nil
	}

	messages, err := db.QueryMessage(ctx, req.GetToUserId(), req.GetUserId())
	if err != nil {
		resp.BaseResp = BaseResp(err)
		return resp, nil
	}

	MessageList := make([]*chat_rpc.RPCMessage, 0)
	for _, msg := range messages {
		MessageList = append(MessageList, &chat_rpc.RPCMessage{
			Id:         msg.ID,
			ToUserId:   msg.ToUserID,
			FromUserId: msg.FromUserID,
			Content:    msg.Content,
			// CreateTime: msg.CreatedAt.Format("2006-01-02 15:04:05"),
			CreateTime: msg.CreatedAt.Unix(),
		})
	}
	resp.BaseResp = BaseResp(errno.Success)
	resp.MessageList = MessageList
	return resp, nil
}
