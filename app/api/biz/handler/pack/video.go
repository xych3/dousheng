package pack

import (
	"gitee.com/derrickball/douyin/app/api/biz/model/common"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
)

func Video(v *video_rpc.Video) *common.Video {
	return &common.Video{
		Id:            v.Id,
		Author:        UserVideo(v.Author),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

func Videos(vs []*video_rpc.Video) []*common.Video {
	videos := make([]*common.Video, 0)
	for _, v := range vs {
		if video := Video(v); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}
