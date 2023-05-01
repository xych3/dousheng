package main

import (
	"context"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/video/dal/db"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	"gitee.com/derrickball/douyin/app/video/pack"
	"gitee.com/derrickball/douyin/app/video/rpc"
	"gitee.com/derrickball/douyin/app/video/service"
	"gitee.com/derrickball/douyin/pkg/errno"
	"log"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// QueryVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryVideo(ctx context.Context, req *video_rpc.QueryVideoReq) (resp *video_rpc.QueryVideoResp, err error) {
	resp = new(video_rpc.QueryVideoResp)
	if (req.LatestTime <= 0) || req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	videos, nextTime, err := service.NewVideoService(ctx).QueryVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = nextTime
	return resp, nil
}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *video_rpc.CreateVideoReq) (resp *video_rpc.CreateVideoResp, err error) {
	log.Println("video create video")
	resp = new(video_rpc.CreateVideoResp)
	if len(req.Title) == 0 || len(req.CoverUrl) == 0 || len(req.PlayUrl) == 0 || req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// QueryVideoByUser implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryVideoByUser(ctx context.Context, req *video_rpc.QueryVideoByUserReq) (resp *video_rpc.QueryVideoByUserResp, err error) {
	// TODO: Your code here...
	resp = new(video_rpc.QueryVideoByUserResp)
	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	videos, err := service.NewVideoService(ctx).QueryVideoByUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

// QueryVideoByIdBatch implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryVideoByIdBatch(ctx context.Context, req *video_rpc.QueryVideoByVideoIdBatchReq) (resp *video_rpc.QueryVideoByVideoIdBatchResp, err error) {
	resp = new(video_rpc.QueryVideoByVideoIdBatchResp)
	if len(req.VideoID) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	videos, err := db.QueryVideoByIdBatch(ctx, req.GetVideoID())
	if len(videos) != 0 {
		resp.VideoList = pack.Videos(ctx, videos, 0)
	} else {
		resp.VideoList = make([]*video_rpc.Video, 0)
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

// IncreaseFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) IncreaseFavoriteCount(ctx context.Context, req *video_rpc.IncreaseFavoriteCountReq) (resp *video_rpc.IncreaseFavoriteCountResp, err error) {
	db.IncreaseFavoriteCountById(ctx, req.GetVideoID())
	return
}

// DecreaseFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DecreaseFavoriteCount(ctx context.Context, req *video_rpc.DecreaseFavoriteCountReq) (resp *video_rpc.DecreaseFavoriteCountResp, err error) {
	db.DecreaseFavoriteCountById(ctx, req.GetVideoID())
	return
}

// IncreaseCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) IncreaseCommentCount(ctx context.Context, req *video_rpc.IncreaseCommentCountReq) (resp *video_rpc.IncreaseCommentCountResp, err error) {
	db.IncreaseCommentCountById(ctx, req.GetVideoID())
	return
}

// DecreaseCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DecreaseCommentCount(ctx context.Context, req *video_rpc.DecreaseCommentCountReq) (resp *video_rpc.DecreaseCommentCountResp, err error) {
	db.DecreaseCommentCountById(ctx, req.GetVideoID())
	return
}

// GetAuthor implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetAuthor(ctx context.Context, req *video_rpc.GetAuthorReq) (resp *video_rpc.GetAuthorResp, err error) {
	AuthorId := db.GetAuthorId(ctx, req.GetVideoID())
	user, err := rpc.GetUser(ctx, &user_rpc.RPCQueryUserReq{
		TargetUserId: AuthorId,
		MyUserId:     req.MyUserID,
	})
	if err != nil {
		return nil, err
	}
	return &video_rpc.GetAuthorResp{
		Author:   user,
		BaseResp: pack.BuildBaseResp(errno.Success),
	}, nil
}
