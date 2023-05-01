package main

import (
	"gitee.com/derrickball/douyin/app/user/dal"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc/userservice"
	"gitee.com/derrickball/douyin/app/user/rpc"
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
	addr, err := net.ResolveTCPAddr("tcp", config.Config.Address.User)
	if err != nil {
		panic(err)
	}
	userSvr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Service.User}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	userErr := userSvr.Run()
	if userErr != nil {
		log.Fatal(userErr)
	}
}
