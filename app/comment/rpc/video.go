package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"
	"gitee.com/derrickball/douyin/app/comment/pack"
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

func IncreaseCommentCount(ctx context.Context, videoID int64) {
	_, err := videoClient.IncreaseCommentCount(ctx, &video_rpc.IncreaseCommentCountReq{
		VideoID: videoID,
	})
	if err != nil {
		return
	}
}

func DecreaseCommentCount(ctx context.Context, videoID int64) {
	_, err := videoClient.DecreaseCommentCount(ctx, &video_rpc.DecreaseCommentCountReq{
		VideoID: videoID,
	})
	if err != nil {
		return
	}
}

func GetAuthor(ctx context.Context, videoID int64, MyUserID int64) *comment_rpc.User {
	author, err := videoClient.GetAuthor(ctx, &video_rpc.GetAuthorReq{
		VideoID:  videoID,
		MyUserID: MyUserID})
	if err != nil {
		return nil
	}
	return pack.UserVideo(author.Author)
}
