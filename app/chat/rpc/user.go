package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc/userservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
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

func QueryUser(ctx context.Context, req *user_rpc.RPCQueryUserReq) (resp *user_rpc.RPCQueryUserResp) {
	resp, err := userClient.QueryUser(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
