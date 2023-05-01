package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc/userservice"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"

	//"gitee.com/derrickball/douyin/kitex_gen/commonrpc"
	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	userClient = userservice.MustNewClient(
		config.Config.Service.User,
		client.WithResolver(r),
	)
}

func GetUser(ctx context.Context, req *user_rpc.RPCQueryUserReq) (*video_rpc.User, error) {
	resp, err := userClient.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	user := &video_rpc.User{
		Id:            resp.Id,
		Name:          resp.Name,
		FollowerCount: resp.FollowerCount,
		FollowCount:   resp.FollowCount,
		IsFollow:      resp.IsFollow,
	}
	return user, nil
}
