package rpc

import (
	"context"

	"gitee.com/derrickball/douyin/app/chat/kitex_gen/chat_rpc"
	"gitee.com/derrickball/douyin/app/chat/kitex_gen/chat_rpc/chatservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var chatClient chatservice.Client

func InitChatRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	chatClient = chatservice.MustNewClient(
		config.Config.Service.Chat,
		client.WithResolver(r),
	)
}

func CreateMessage(ctx context.Context, req *chat_rpc.RPCMessageCreateReq) (*chat_rpc.RPCMessageCreateResp, error) {
	resp, err := chatClient.CreateMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp, nil
}

func ChatList(ctx context.Context, req *chat_rpc.RPCMessageChatReq) (*chat_rpc.RPCMessageChatResp, error) {
	resp, err := chatClient.MessageList(ctx, req)
	if err != nil {
		return &chat_rpc.RPCMessageChatResp{}, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp, nil
}

func FriendList(ctx context.Context, req *chat_rpc.RPCFriendListReq) (*chat_rpc.RPCFriendListResp, error) {
	resp, err := chatClient.FriendList(ctx, req)
	if err != nil {
		return &chat_rpc.RPCFriendListResp{}, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp, nil
}
