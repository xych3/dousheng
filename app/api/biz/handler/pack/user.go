package pack

import (
	"gitee.com/derrickball/douyin/app/api/biz/model/common"
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
)

func UserVideo(u *video_rpc.User) *common.User {
	return &common.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}

func UserComment(u *comment_rpc.User) *common.User {
	return &common.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
