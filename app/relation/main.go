package main

import (
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc/relationservice"
	"log"
	"net"

	"gitee.com/derrickball/douyin/app/relation/dal"

	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	config.LoadConfig()
	dal.Init()
}
func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", config.Config.Address.Relation)
	if err != nil {
		panic(err)
	}
	svr := relationservice.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Service.Relation}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}

	relationSvr := relationservice.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Service.User}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	err = relationSvr.Run()
	if err != nil {
		return
	}

}
