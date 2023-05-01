package rpc

import (
	"context"

	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc/commentservice"
	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var cClient commentservice.Client

func InitCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{config.Config.Address.Etcd})
	if err != nil {
		panic(err)
	}
	cClient = commentservice.MustNewClient(
		config.Config.Service.Comment,
		client.WithResolver(r),
	)
}

func CreateComment(ctx context.Context, req *comment_rpc.RPCCommentCreateReq) (*comment_rpc.RPCActionResp, error) {
	resp, err := cClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp, nil
}

func DelComment(ctx context.Context, req *comment_rpc.RPCCommentDelReq) (*comment_rpc.RPCActionResp, error) {
	resp, err := cClient.DelComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *comment_rpc.RPCCommentListReq) (*comment_rpc.RPCCommentListResp, error) {
	resp, err := cClient.CommentList(ctx, req)
	if err != nil {
		return &comment_rpc.RPCCommentListResp{}, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}
