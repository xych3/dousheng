package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc/userservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
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

func CreateUser(ctx context.Context, req *user_rpc.RPCCreateUserReq) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func CheckUser(ctx context.Context, req *user_rpc.RPCCheckUserReq) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func QueryUser(ctx context.Context, req *user_rpc.RPCQueryUserReq) (resp *user_rpc.RPCQueryUserResp) {
	resp, err := userClient.QueryUser(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
