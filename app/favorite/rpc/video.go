package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc/videoservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	videoClient = videoservice.MustNewClient(
		config.Config.Service.Video,
		client.WithResolver(r),
	)
}

func IncreaseFavoriteCount(ctx context.Context, videoID int64) {
	_, err := videoClient.IncreaseFavoriteCount(ctx, &video_rpc.IncreaseFavoriteCountReq{
		VideoID: videoID,
	})
	if err != nil {
		return
	}
}

func DecreaseFavoriteCount(ctx context.Context, videoID int64) {
	_, err := videoClient.DecreaseFavoriteCount(ctx, &video_rpc.DecreaseFavoriteCountReq{
		VideoID: videoID,
	})
	if err != nil {
		return
	}
}
