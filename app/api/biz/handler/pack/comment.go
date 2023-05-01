package pack

import (
	"gitee.com/derrickball/douyin/app/api/biz/model/common"
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"
)

func Comment(u *comment_rpc.Comment) *common.Comment {
	return &common.Comment{
		Id:         u.Id,
		Content:    u.Content,
		CreateDate: u.CreateDate,
		User:       UserComment(u.User),
	}
}

func Comments(cs []*comment_rpc.Comment) []*common.Comment {
	comments := make([]*common.Comment, 0)
	for _, c := range cs {
		if comment := Comment(c); comment != nil {
			comments = append(comments, comment)
		}
	}
	return comments
}
