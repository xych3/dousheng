package service

import (
	"context"
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc"
	"log"

	"gitee.com/derrickball/douyin/app/relation/dal/db"
	"gitee.com/derrickball/douyin/app/user/model"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{ctx: ctx}
}

func (s *RelationService) Action(req *relation_rpc.RPCRelationActionReq) error {
	relations, err := db.QueryRelation(s.ctx, req.ToUserId, req.FollowerId)
	if err != nil {
		return err
	}
	if len(relations) != 0 {
		return nil
	}
	err = db.CreateRelation(s.ctx, &model.Relation{
		ToUserID:   req.ToUserId,
		FollowerID: req.FollowerId,
	})
	if err != nil {
		return err
	}
	go func() {
		err := db.UpdateFollowCount(s.ctx, req.ToUserId, req.FollowerId, true)
		if err != nil {
			log.Println(err)
		}
	}()
	return nil
}

func (s *RelationService) Cancel(req *relation_rpc.RPCRelationActionReq) error {
	relations, err := db.QueryRelation(s.ctx, req.ToUserId, req.FollowerId)
	if err != nil {
		return err
	}
	if len(relations) == 0 {
		return nil
	}
	err = db.DeleteRelation(s.ctx, req.ToUserId, req.FollowerId)
	if err != nil {
		return err
	}
	go func() {
		err := db.UpdateFollowCount(s.ctx, req.ToUserId, req.FollowerId, false)
		if err != nil {
			log.Println(err)
		}
	}()
	return nil
}

func (s *RelationService) FollowList(req *relation_rpc.RPCRelationFollowListReq) ([]*model.User, []bool, error) {
	relations, err := db.QueryRelationFollow(s.ctx, req.FollowerId)
	if err != nil {
		return nil, nil, err
	}
	queryRows := len(relations)

	if queryRows == 0 {
		return nil, nil, nil
	}
	follows := make([]int64, queryRows)
	flags := make([]bool, queryRows)
	for k, v := range relations {
		follows[k] = v.ToUserID
		flags[k], _ = db.QueryIsFollow(s.ctx, v.ToUserID, req.UserId)
	}
	users, err := db.QueryUserById(s.ctx, follows)
	if err != nil {
		return nil, flags, err
	}
	return users, flags, nil
}

func (s *RelationService) FollowerList(req *relation_rpc.RPCRelationFollowerListReq) ([]*model.User, []bool, error) {
	relations, err := db.QueryRelationFollower(s.ctx, req.ToUserId)
	if err != nil {
		return nil, nil, err
	}
	queryRows := len(relations)
	if queryRows == 0 {
		return nil, nil, nil
	}
	followers := make([]int64, queryRows)
	flags := make([]bool, queryRows)
	for k, v := range relations {
		followers[k] = v.FollowerID
		flags[k], _ = db.QueryIsFollow(s.ctx, v.FollowerID, req.UserId)
	}
	users, err := db.QueryUserById(s.ctx, followers)
	if err != nil {
		return nil, flags, err
	}
	return users, flags, nil
}

func (s *RelationService) CheckIsFollow(req *relation_rpc.RPCCheckFollowReq) (bool, error) {
	isFollow, err := db.QueryIsFollow(s.ctx, req.ToUserId, req.FollowerId)
	if err != nil {
		return false, err
	}
	return isFollow, err
}
