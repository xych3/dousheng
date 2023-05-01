package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"gitee.com/derrickball/douyin/app/user/dal/db"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/app/user/model"
	"gitee.com/derrickball/douyin/pkg/errno"
	"io"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) CheckUser(req *user_rpc.RPCCheckUserReq) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 1 {
		return 0, errno.AuthorizationFailedErr
	}
	h := md5.New()
	_, err = io.WriteString(h, req.Password)
	if err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	user := users[0]
	if user.Password != password {
		return 0, errno.AuthorizationFailedErr
	}
	return user.ID, nil
}

func (s *UserService) CreateUser(req *user_rpc.RPCCreateUserReq) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return errno.UserAlreadyExistErr
	}
	h := md5.New()
	_, err = io.WriteString(h, req.Password)
	if err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	err = db.CreateUser(s.ctx, &model.User{
		Username:      req.Username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	})
	if err != nil {
		return err
	}
	return nil
}
