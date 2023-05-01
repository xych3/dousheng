package cron

import (
	"github.com/prometheus/common/log"
	"github.com/robfig/cron"
)

func InitCron() {
	c := cron.New()
	spec := "0 0 */1 * * ?"
	err := c.AddFunc(spec, SyncToDB)
	if err != nil {
		return
	}
	go InitVideoCount()
	go InitUserCount()
}

func SyncToDB() {
	log.Info("----定时任务同步开始----")
	go SyncVideoCount()
	go SyncUserCount()
}

// TODO 初始化视频数据
func InitVideoCount() {
	log.Info("----初始化视频数据----")

}

// TODO 初始化用户表数据
func InitUserCount() {
	log.Info("----初始化用户数据----")
}

// TODO 同步视频表的数据
func SyncVideoCount() {
	log.Info("----同步视频数据----")
}

// TODO 同步用户表的数据
func SyncUserCount() {
	log.Info("----同步用户数据----")

}
