package rpc

import (
	"context"
	"gitee.com/derrickball/douyin/app/favorite/kitex_gen/favorite_rpc"
	"gitee.com/derrickball/douyin/app/favorite/kitex_gen/favorite_rpc/favoriteservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	favoriteClient = favoriteservice.MustNewClient(
		config.Config.Service.Favorite,
		client.WithResolver(r),
	)
}

func FilterFavorite(ctx context.Context, req *favorite_rpc.RPCFilterFavoriteReq) (map[int64]bool, error) {
	resp, err := favoriteClient.FilterFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]bool)
	for _, videoId := range resp.VideoIdList {
		res[videoId] = true
	}
	return res, nil
}
