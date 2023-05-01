package pack

import (
	"gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	"gitee.com/derrickball/douyin/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *video_rpc.BaseResp {
	errNo := errno.ConvertErr(err)
	return baseResp(errNo)
}

func baseResp(err errno.ErrNo) *video_rpc.BaseResp {
	return &video_rpc.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
