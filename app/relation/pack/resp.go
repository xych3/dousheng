package pack

import (
	"gitee.com/derrickball/douyin/app/relation/kitex_gen/relation_rpc"
	"time"

	"gitee.com/derrickball/douyin/pkg/errno"
)

func BuildRelaBaseResp(err error) *relation_rpc.BaseResp {
	errNo := errno.ConvertErr(err)
	return baseRelaResp(errNo)
}

func baseRelaResp(err errno.ErrNo) *relation_rpc.BaseResp {
	return &relation_rpc.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
