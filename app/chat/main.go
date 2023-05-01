package main

import (
	"gitee.com/derrickball/douyin/app/chat/dal"
	"gitee.com/derrickball/douyin/app/chat/rpc"
	"log"
	"net"

	"gitee.com/derrickball/douyin/app/chat/kitex_gen/chat_rpc/chatservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
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
	addr, err := net.ResolveTCPAddr("tcp", config.Config.Address.Chat)
	if err != nil {
		panic(err)
	}
	chatSvr := chatservice.NewServer(new(ChatServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Service.Chat}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	chatErr := chatSvr.Run()
	if chatErr != nil {
		log.Println(chatErr.Error())
	}
}
