package main

import (
	"gitee.com/derrickball/douyin/app/video/dal"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc/videoservice"
	"gitee.com/derrickball/douyin/app/video/rpc"
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
	addr, err := net.ResolveTCPAddr("tcp", config.Config.Address.Video)
	if err != nil {
		panic(err)
	}
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Service.Video}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}
}
