package main

import (
	"context"
	"fmt"
	"gitee.com/derrickball/douyin/app/favorite/dal/db"
	"gitee.com/derrickball/douyin/app/favorite/kitex_gen/favorite_rpc"
	"gitee.com/derrickball/douyin/app/favorite/model"
	"gitee.com/derrickball/douyin/app/favorite/rpc"
	"gitee.com/derrickball/douyin/pkg/errno"
	"time"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

func BaseResp(err error) *favorite_rpc.BaseResp {
	errNo := errno.ConvertErr(err)
	return &favorite_rpc.BaseResp{
		StatusCode:    errNo.ErrCode,
		StatusMessage: errNo.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}

// CreateFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CreateFavorite(ctx context.Context, req *favorite_rpc.RPCFavoriteCreateReq) (resp *favorite_rpc.RPCFavoriteCreateResp, err error) {
	// TODO: Your code here...
	resp = new(favorite_rpc.RPCFavoriteCreateResp)
	if len(fmt.Sprint(req.GetUserId())) == 0 {
		resp.BaseResp = BaseResp(errno.AuthorizationFailedErr)
		return resp, nil
	}
	if len(fmt.Sprint(req.GetVideoId())) == 0 {
		resp.BaseResp = BaseResp(errno.ParamErr)
		return resp, nil
	}
	// mysql
	err = db.CreateFavorite(ctx, &model.Favorite{
		UserID:  req.GetUserId(),
		VideoID: req.GetVideoId(),
	})
	if err != nil {
		resp.BaseResp = BaseResp(err)
		return resp, nil
	}
	// 异步增加视频点赞计数
	go rpc.IncreaseFavoriteCount(ctx, req.GetVideoId())
	resp.BaseResp = BaseResp(errno.Success)
	return resp, nil
}

// DelFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) DelFavorite(ctx context.Context, req *favorite_rpc.RPCFavoriteDelReq) (resp *favorite_rpc.RPCFavoriteDelResp, err error) {
	// TODO: Your code here...
	resp = new(favorite_rpc.RPCFavoriteDelResp)
	if len(fmt.Sprint(req.GetUserId())) == 0 {
		resp.BaseResp = BaseResp(errno.AuthorizationFailedErr)
		return resp, nil
	}
	if len(fmt.Sprint(req.GetVideoId())) == 0 {
		resp.BaseResp = BaseResp(errno.ParamErr)
		return resp, nil
	}
	// mysql
	err = db.DeleteFavorite(ctx, req.GetUserId(), req.GetVideoId())
	if err != nil {
		resp.BaseResp = BaseResp(err)
		return resp, nil
	}
	// 异步减少点赞计数
	go rpc.DecreaseFavoriteCount(ctx, req.GetVideoId())
	resp.BaseResp = BaseResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite_rpc.RPCFavoriteListReq) (resp *favorite_rpc.RPCFavoriteListResp, err error) {
	// TODO: Your code here...
	resp = new(favorite_rpc.RPCFavoriteListResp)
	if len(fmt.Sprint(req.GetUserId())) == 0 {
		resp.BaseResp = BaseResp(errno.AuthorizationFailedErr)
		return resp, nil
	}

	favorites, err := db.QueryFavorite(ctx, req.GetUserId())
	if err != nil {
		resp.BaseResp = BaseResp(err)
		return resp, nil
	}
	//TODO 投稿列表
	resp.VideoIdList = make([]int64, 0)
	for _, fav := range favorites {
		resp.VideoIdList = append(resp.VideoIdList, fav.VideoID)
	}
	resp.BaseResp = BaseResp(errno.Success)
	return resp, nil
}

// FilterFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FilterFavorite(ctx context.Context, req *favorite_rpc.RPCFilterFavoriteReq) (resp *favorite_rpc.RPCFilterFavoriteResp, err error) {
	resp = new(favorite_rpc.RPCFilterFavoriteResp)
	if req.UserId <= 0 || len(req.VideoIdList) == 0 {
		resp.BaseResp = BaseResp(errno.ParamErr)
		return resp, nil
	}

	favorites, err := db.FilterFavorite(ctx, req.UserId, req.VideoIdList)
	if err != nil {
		return nil, err
	}
	resp.VideoIdList = make([]int64, 0)
	for _, favorite := range favorites {
		resp.VideoIdList = append(resp.VideoIdList, favorite.VideoID)
	}
	resp.BaseResp = BaseResp(errno.Success)
	return resp, nil
}
