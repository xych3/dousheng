package rpc

import (
	"context"
	"fmt"
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc"
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc/relationservice"

	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		fmt.Println("abc123")
		panic(err)
	}
	fmt.Println("relationservice name :", config.Config.Service.Relation)
	relationClient = relationservice.MustNewClient(
		config.Config.Service.Relation,
		client.WithResolver(r),
	)
}

func Action(ctx context.Context, req *relation_rpc.RPCRelationActionReq) error {
	resp, err := relationClient.Action(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func FollowList(ctx context.Context, req *relation_rpc.RPCRelationFollowListReq) (*relation_rpc.RPCRelationFollowListResp, error) {
	resp, err := relationClient.FollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, err
}

func FollowerList(ctx context.Context, req *relation_rpc.RPCRelationFollowerListReq) (*relation_rpc.RPCRelationFollowerListResp, error) {
	resp, err := relationClient.FollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, err
}
