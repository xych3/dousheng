package main

import (
	"log"

	"gitee.com/derrickball/douyin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func newGenerator(modelPkgPath string, dsn string) *gen.Generator {
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: modelPkgPath,
		Mode:         gen.WithoutContext,
	})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("fail to open db")
	}
	g.UseDB(db)
	return g
}

func generateUser() {
	g := newGenerator("app/user/model", config.Config.MySQL.DSN)
	g.GenerateModel("users")
	g.GenerateModel("relations")
	g.Execute()
}

func generateVideo() {
	g := newGenerator("app/video/model", config.Config.MySQL.DSN)
	g.GenerateModel("videos")
	g.GenerateModel("favorites")
	g.GenerateModel("comments")
	g.Execute()
}

func generateChat() {
	g := newGenerator("app/chat/model", config.Config.MySQL.DSN)
	g.GenerateModel("chat")
	g.Execute()
}

func main() {
	config.LoadConfig()
	generateUser()
	generateVideo()
	generateChat()
}
