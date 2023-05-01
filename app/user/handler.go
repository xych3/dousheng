package main

import (
	"context"
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc"
	"gitee.com/derrickball/douyin/app/user/dal/db"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/user/pack"
	"gitee.com/derrickball/douyin/app/user/rpc"
	"gitee.com/derrickball/douyin/app/user/service"
	"gitee.com/derrickball/douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user_rpc.RPCCheckUserReq) (resp *user_rpc.RPCCheckUserResp, err error) {
	resp = new(user_rpc.RPCCheckUserResp)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	userID, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = userID
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user_rpc.RPCCreateUserReq) (resp *user_rpc.RPCCreateUserResp, err error) {
	resp = new(user_rpc.RPCCreateUserResp)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user_rpc.RPCQueryUserReq) (resp *user_rpc.RPCQueryUserResp, err error) {
	resp = new(user_rpc.RPCQueryUserResp)
	if req.TargetUserId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	user, err := db.QueryUserByID(ctx, req.TargetUserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Id = user.ID
	resp.Name = user.Username
	resp.FollowerCount = user.FollowerCount
	resp.FollowCount = user.FollowCount
	resp.IsFollow = rpc.CheckFollow(ctx, &relation_rpc.RPCCheckFollowReq{
		ToUserId:   req.TargetUserId,
		FollowerId: req.MyUserId,
	})
	return resp, nil
}

//// GetUser implements the UserServiceImpl interface.
//func (s *UserServiceImpl) GetUser(ctx context.Context, req *userrpc.GetUserReq) (resp *userrpc.GetUserResp, err error) {
//	// TODO: Your code here...
//	resp = new(userrpc.GetUserResp)
//	if req.UserId <= 0 || req.MyUserId <= 0 {
//		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
//		return resp, nil
//	}
//	user, err := service.NewUserService(ctx).GetUser(req)
//	if err != nil {
//		resp.BaseResp = pack.BuildBaseResp(err)
//		return resp, nil
//	}
//	resp.User = user
//	resp.BaseResp = pack.BuildBaseResp(errno.Success)
//	return resp, nil
//}
