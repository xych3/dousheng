package main

import (
	"gitee.com/derrickball/douyin/app/comment/dal"
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc/commentservice"
	"gitee.com/derrickball/douyin/app/comment/rpc"
	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func Init() {
	config.LoadConfig()
	dal.Init()
	rpc.Init()
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", config.Config.Address.Comment)
	if err != nil {
		panic(err)
	}
	commentSvr := commentservice.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Service.Comment}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	commentErr := commentSvr.Run()
	if commentErr != nil {
		log.Println(commentErr.Error())
	}
}
