package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc/videoservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
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

func CreateVideo(ctx context.Context, req *video_rpc.CreateVideoReq) error {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func QueryVideo(ctx context.Context, req *video_rpc.QueryVideoReq) ([]*video_rpc.Video, int64, error) {
	resp, err := videoClient.QueryVideo(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, resp.NextTime, nil
}

func QueryVideoByUser(ctx context.Context, req *video_rpc.QueryVideoByUserReq) ([]*video_rpc.Video, error) {
	resp, err := videoClient.QueryVideoByUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}

func QueryVideoByIDBatch(ctx context.Context, req *video_rpc.QueryVideoByVideoIdBatchReq) ([]*video_rpc.Video, error) {
	resp, err := videoClient.QueryVideoByIdBatch(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil

}
