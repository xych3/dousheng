package pack

import (
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"
	"gitee.com/derrickball/douyin/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *comment_rpc.BaseResp {
	errNo := errno.ConvertErr(err)
	return baseResp(errNo)
}

func baseResp(err errno.ErrNo) *comment_rpc.BaseResp {
	return &comment_rpc.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
