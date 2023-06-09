// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"gitee.com/derrickball/douyin/app/api/biz/cron"
	"log"

	"gitee.com/derrickball/douyin/app/api/biz/mw"
	"gitee.com/derrickball/douyin/app/api/biz/rpc"
	"gitee.com/derrickball/douyin/pkg/config"
	"gitee.com/derrickball/douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Init() {
	config.LoadConfig()
	rpc.Init()
	mw.InitAuth()
}

func main() {
	Init()
	h := server.Default(
		server.WithHostPorts(config.Config.Address.Api),
		server.WithMaxRequestBodySize(20<<20),
	)

	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
		log.Printf("[Recovery] err=%v\nstack=%s", err, stack)
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"code":    errno.ServiceErrCode,
			"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
		})
	})))

	h.Use(mw.Logger())

	register(h)

	h.Spin()

	// 启动定时任务
	cron.InitCron()
}
