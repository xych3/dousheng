package rpc

import (
	"context"
	"log"

	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc"
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc/relationservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	relationClient = relationservice.MustNewClient(
		config.Config.Service.Relation,
		client.WithResolver(r),
	)
}

func CheckFollow(ctx context.Context, req *relation_rpc.RPCCheckFollowReq) bool {
	resp, err := relationClient.CheckFollow(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	return resp.IsFollow
}
