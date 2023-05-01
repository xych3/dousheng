// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	video_rpc "gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	QueryVideo(ctx context.Context, Req *video_rpc.QueryVideoReq, callOptions ...callopt.Option) (r *video_rpc.QueryVideoResp, err error)
	CreateVideo(ctx context.Context, Req *video_rpc.CreateVideoReq, callOptions ...callopt.Option) (r *video_rpc.CreateVideoResp, err error)
	QueryVideoByUser(ctx context.Context, Req *video_rpc.QueryVideoByUserReq, callOptions ...callopt.Option) (r *video_rpc.QueryVideoByUserResp, err error)
	QueryVideoByIdBatch(ctx context.Context, Req *video_rpc.QueryVideoByVideoIdBatchReq, callOptions ...callopt.Option) (r *video_rpc.QueryVideoByVideoIdBatchResp, err error)
	IncreaseFavoriteCount(ctx context.Context, Req *video_rpc.IncreaseFavoriteCountReq, callOptions ...callopt.Option) (r *video_rpc.IncreaseFavoriteCountResp, err error)
	DecreaseFavoriteCount(ctx context.Context, Req *video_rpc.DecreaseFavoriteCountReq, callOptions ...callopt.Option) (r *video_rpc.DecreaseFavoriteCountResp, err error)
	IncreaseCommentCount(ctx context.Context, Req *video_rpc.IncreaseCommentCountReq, callOptions ...callopt.Option) (r *video_rpc.IncreaseCommentCountResp, err error)
	DecreaseCommentCount(ctx context.Context, Req *video_rpc.DecreaseCommentCountReq, callOptions ...callopt.Option) (r *video_rpc.DecreaseCommentCountResp, err error)
	GetAuthor(ctx context.Context, Req *video_rpc.GetAuthorReq, callOptions ...callopt.Option) (r *video_rpc.GetAuthorResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) QueryVideo(ctx context.Context, Req *video_rpc.QueryVideoReq, callOptions ...callopt.Option) (r *video_rpc.QueryVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideo(ctx, Req)
}

func (p *kVideoServiceClient) CreateVideo(ctx context.Context, Req *video_rpc.CreateVideoReq, callOptions ...callopt.Option) (r *video_rpc.CreateVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVideo(ctx, Req)
}

func (p *kVideoServiceClient) QueryVideoByUser(ctx context.Context, Req *video_rpc.QueryVideoByUserReq, callOptions ...callopt.Option) (r *video_rpc.QueryVideoByUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoByUser(ctx, Req)
}

func (p *kVideoServiceClient) QueryVideoByIdBatch(ctx context.Context, Req *video_rpc.QueryVideoByVideoIdBatchReq, callOptions ...callopt.Option) (r *video_rpc.QueryVideoByVideoIdBatchResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoByIdBatch(ctx, Req)
}

func (p *kVideoServiceClient) IncreaseFavoriteCount(ctx context.Context, Req *video_rpc.IncreaseFavoriteCountReq, callOptions ...callopt.Option) (r *video_rpc.IncreaseFavoriteCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IncreaseFavoriteCount(ctx, Req)
}

func (p *kVideoServiceClient) DecreaseFavoriteCount(ctx context.Context, Req *video_rpc.DecreaseFavoriteCountReq, callOptions ...callopt.Option) (r *video_rpc.DecreaseFavoriteCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DecreaseFavoriteCount(ctx, Req)
}

func (p *kVideoServiceClient) IncreaseCommentCount(ctx context.Context, Req *video_rpc.IncreaseCommentCountReq, callOptions ...callopt.Option) (r *video_rpc.IncreaseCommentCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IncreaseCommentCount(ctx, Req)
}

func (p *kVideoServiceClient) DecreaseCommentCount(ctx context.Context, Req *video_rpc.DecreaseCommentCountReq, callOptions ...callopt.Option) (r *video_rpc.DecreaseCommentCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DecreaseCommentCount(ctx, Req)
}

func (p *kVideoServiceClient) GetAuthor(ctx context.Context, Req *video_rpc.GetAuthorReq, callOptions ...callopt.Option) (r *video_rpc.GetAuthorResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAuthor(ctx, Req)
}
