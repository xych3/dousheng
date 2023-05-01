package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/favorite/kitex_gen/favorite_rpc"
	"gitee.com/derrickball/douyin/app/favorite/kitex_gen/favorite_rpc/favoriteservice"

	//"gitee.com/derrickball/douyin/app/video/kitex_gen/favoriterpc/favoriteservice"

	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

//var favoriteClient favoritehandler.Client
//var favoriteClient favoriterpc.FavoriteService
var fClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	fClient = favoriteservice.MustNewClient(
		config.Config.Service.Favorite,
		client.WithResolver(r),
	)
}

func CreateFavorite(ctx context.Context, req *favorite_rpc.RPCFavoriteCreateReq) error {

	resp, err := fClient.CreateFavorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DelFavorite(ctx context.Context, req *favorite_rpc.RPCFavoriteDelReq) error {
	resp, err := fClient.DelFavorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func FavoriteList(ctx context.Context, req *favorite_rpc.RPCFavoriteListReq) ([]int64, error) {
	resp, err := fClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	list := resp.VideoIdList
	return list, nil
}
