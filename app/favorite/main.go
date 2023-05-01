package main

import (
	"gitee.com/derrickball/douyin/app/favorite/dal"
	"gitee.com/derrickball/douyin/app/favorite/kitex_gen/favorite_rpc/favoriteservice"
	"gitee.com/derrickball/douyin/app/favorite/rpc"
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
	addr, err := net.ResolveTCPAddr("tcp", config.Config.Address.Favorite)
	if err != nil {
		panic(err)
	}

	favoriteSvr := favoriteservice.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Service.Favorite}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	favoriteErr := favoriteSvr.Run()
	if favoriteErr != nil {
		log.Println(favoriteErr.Error())
	}
}
