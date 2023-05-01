package main

import (
	"context"
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc"
	"gitee.com/derrickball/douyin/app/relation/pack"
	"gitee.com/derrickball/douyin/app/relation/service"
	"gitee.com/derrickball/douyin/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Action implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Action(ctx context.Context, req *relation_rpc.RPCRelationActionReq) (resp *relation_rpc.RPCRelationActionResp, err error) {
	resp = new(relation_rpc.RPCRelationActionResp)

	if req.ActionType == 1 {
		err = service.NewRelationService(ctx).Action(req)
		if err != nil {
			resp.BaseResp = pack.BuildRelaBaseResp(err)
			return resp, nil
		}
	}
	if req.ActionType == 2 {
		err = service.NewRelationService(ctx).Cancel(req)
		if err != nil {
			resp.BaseResp = pack.BuildRelaBaseResp(err)
			return resp, nil
		}
	}
	resp.BaseResp = pack.BuildRelaBaseResp(errno.Success)
	return resp, nil
}

// FollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowList(ctx context.Context, req *relation_rpc.RPCRelationFollowListReq) (resp *relation_rpc.RPCRelationFollowListResp, err error) {
	resp = new(relation_rpc.RPCRelationFollowListResp)
	users, flags, err := service.NewRelationService(ctx).FollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildRelaBaseResp(err)
		return resp, nil
	}
	followList := make([]*relation_rpc.User, 0)
	for k, v := range users {
		followList = append(followList, &relation_rpc.User{
			Id:            v.ID,
			Name:          v.Username,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      flags[k],
		})
	}
	resp.BaseResp = pack.BuildRelaBaseResp(errno.Success)
	resp.FollowList = followList
	return resp, nil
}

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *relation_rpc.RPCRelationFollowerListReq) (resp *relation_rpc.RPCRelationFollowerListResp, err error) {
	resp = new(relation_rpc.RPCRelationFollowerListResp)
	users, flags, err := service.NewRelationService(ctx).FollowerList(req)
	if err != nil {
		resp.BaseResp = pack.BuildRelaBaseResp(err)
		return resp, nil
	}
	followerList := make([]*relation_rpc.User, 0)
	for k, v := range users {
		followerList = append(followerList, &relation_rpc.User{
			Id:            v.ID,
			Name:          v.Username,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      flags[k],
		})
	}
	resp.BaseResp = pack.BuildRelaBaseResp(errno.Success)
	resp.FollowerList = followerList
	return resp, nil
}

// CheckFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CheckFollow(ctx context.Context, req *relation_rpc.RPCCheckFollowReq) (resp *relation_rpc.RPCCheckFollowResp, err error) {
	resp = new(relation_rpc.RPCCheckFollowResp)
	isFollow, err := service.NewRelationService(ctx).CheckIsFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildRelaBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildRelaBaseResp(errno.Success)
	resp.IsFollow = isFollow
	return resp, nil
}
