package pack

import (
	"context"
	"gitee.com/derrickball/douyin/app/favorite/kitex_gen/favorite_rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	"gitee.com/derrickball/douyin/app/video/model"
	"gitee.com/derrickball/douyin/app/video/rpc"
)

func Video(ctx context.Context, v *model.Video, myUserID int64) *video_rpc.Video {
	if v == nil {
		return nil
	}
	user, err := rpc.GetUser(ctx, &user_rpc.RPCQueryUserReq{
		TargetUserId: v.UserID,
		MyUserId:     myUserID,
	})
	if err != nil {
		return nil
	}

	return &video_rpc.Video{
		Id:            v.ID,
		Author:        user,
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    false,
		Title:         v.Title,
	}
}

// Videos param userID为0的话会置所有视频为喜欢 使用场景：喜欢列表
func Videos(ctx context.Context, vs []*model.Video, userID int64) []*video_rpc.Video {
	// todo batch rpc GerUser
	videos := make([]*video_rpc.Video, 0)
	ids := make([]int64, 0)
	for _, v := range vs {
		if video := Video(ctx, v, userID); video != nil {
			videos = append(videos, video)
			ids = append(ids, v.ID)
		}
	}
	filterMap := make(map[int64]bool)
	if userID > 0 {
		var err error
		filterMap, err = rpc.FilterFavorite(ctx, &favorite_rpc.RPCFilterFavoriteReq{
			UserId:      userID,
			VideoIdList: ids,
		})
		if err != nil {
			return nil
		}
		for _, video := range videos {
			if _, exist := filterMap[video.Id]; exist {
				video.IsFavorite = true
			}
		}
	} else if userID == 0 {
		for _, video := range videos {
			video.IsFavorite = false
		}
	}

	return videos
}
