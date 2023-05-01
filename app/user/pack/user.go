package pack

import (
	"gitee.com/derrickball/douyin/app/user/model"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
)

func User(u *model.User, isFollower bool) *video_rpc.User {
	return &video_rpc.User{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      isFollower,
	}
}
