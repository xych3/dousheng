package pack

import (
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *user_rpc.BaseResp {
	errNo := errno.ConvertErr(err)
	return baseResp(errNo)
}

func baseResp(err errno.ErrNo) *user_rpc.BaseResp {
	return &user_rpc.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
