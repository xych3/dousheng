package pack

import (
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
)

func User(u *user_rpc.RPCQueryUserResp) *comment_rpc.User {
	return &comment_rpc.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}

func UserVideo(user *video_rpc.User) *comment_rpc.User {
	return &comment_rpc.User{
		Id:            user.Id,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
		Name:          user.Name,
	}
}
