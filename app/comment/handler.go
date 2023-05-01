package main

import (
	"context"
	"fmt"

	"gitee.com/derrickball/douyin/app/comment/dal/db"
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"
	"gitee.com/derrickball/douyin/app/comment/model"
	"gitee.com/derrickball/douyin/app/comment/pack"
	"gitee.com/derrickball/douyin/app/comment/rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *comment_rpc.RPCCommentCreateReq) (resp *comment_rpc.RPCActionResp, err error) {
	resp = new(comment_rpc.RPCActionResp)
	if len(fmt.Sprint(req.GetUserId())) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.AuthorizationFailedErr)
		return resp, nil
	}
	if len(fmt.Sprint(req.GetVideoId())) == 0 || len(req.GetCommentText()) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	// mysql
	comment := &model.Comment{
		UserID:      req.GetUserId(),
		VideoID:     req.GetVideoId(),
		CommentText: req.GetCommentText(),
	}

	err = db.CreateComment(ctx, comment)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	//user := rpc.QueryUser(ctx, &user_rpc.RPCQueryUserReq{UserId: req.UserId})
	//u := pack.User(user, true)
	u := rpc.GetAuthor(ctx, req.VideoId, req.UserId)
	resp.Comment = &comment_rpc.Comment{
		Id:         comment.ID,
		User:       u,
		Content:    req.CommentText,
		CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	go rpc.IncreaseCommentCount(ctx, req.GetVideoId())
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil

}

// DelComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DelComment(ctx context.Context, req *comment_rpc.RPCCommentDelReq) (resp *comment_rpc.RPCActionResp, err error) {
	resp = new(comment_rpc.RPCActionResp)
	content, createTime, err := db.DelComments(ctx, req.GetCommentId())
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	//user := rpc.QueryUser(ctx, &user_rpc.RPCQueryUserReq{UserId: req.UserId})
	user := rpc.GetAuthor(ctx, req.VideoId, req.UserId)
	resp.Comment = &comment_rpc.Comment{
		Id:         req.CommentId,
		User:       user,
		Content:    content,
		CreateDate: createTime,
	}
	go rpc.DecreaseCommentCount(ctx, req.GetVideoId())
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment_rpc.RPCCommentListReq) (resp *comment_rpc.RPCCommentListResp, err error) {
	resp = new(comment_rpc.RPCCommentListResp)
	if len(fmt.Sprint(req.GetVideoId())) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	comments, err := db.QueryComments(ctx, req.GetVideoId())
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	//isFollow := false

	resp.CommentList = make([]*comment_rpc.Comment, 0)
	for _, comm := range comments {

		u := rpc.QueryUser(ctx, &user_rpc.RPCQueryUserReq{
			MyUserId:     req.UserId,
			TargetUserId: comm.UserID,
		})
		resp.CommentList = append(resp.CommentList, &comment_rpc.Comment{
			Id:         comm.ID,
			User:       pack.User(u),
			Content:    comm.CommentText,
			CreateDate: comm.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
