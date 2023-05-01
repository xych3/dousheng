package service

import (
	"context"
	"gitee.com/derrickball/douyin/app/video/dal/db"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	"gitee.com/derrickball/douyin/app/video/model"
	"gitee.com/derrickball/douyin/app/video/pack"
	"gitee.com/derrickball/douyin/pkg/constants"
	"time"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (s *VideoService) CreateVideo(req *video_rpc.CreateVideoReq) error {
	video := &model.Video{
		Title:         req.Title,
		UserID:        req.UserId,
		PlayURL:       req.PlayUrl,
		CoverURL:      req.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err := db.CreateVideo(s.ctx, video)
	if err != nil {
		return err
	}
	return nil
}

func (s *VideoService) QueryVideo(req *video_rpc.QueryVideoReq) ([]*video_rpc.Video, int64, error) {
	var latestTime int64
	if req.LatestTime == 0 {
		latestTime = time.Now().Unix()
	} else {
		latestTime = req.LatestTime
	}
	videos, err := db.QueryVideo(s.ctx, latestTime, constants.DefaultLimit)
	if err != nil {
		return nil, 0, err
	}
	var nextTime int64
	if len(videos) == 0 {
		nextTime = latestTime
	} else {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	return pack.Videos(s.ctx, videos, req.UserId), nextTime, nil
}

func (s *VideoService) QueryVideoByUser(req *video_rpc.QueryVideoByUserReq) ([]*video_rpc.Video, error) {
	videos, err := db.QueryVideoByUserID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.Videos(s.ctx, videos, req.MyUserId), nil
}
