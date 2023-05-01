package db

import (
	"fmt"
	"gitee.com/derrickball/douyin/pkg/config"
	"testing"
)

func TestQueryUserByID(t *testing.T) {
	config.LoadConfig()
	Init()
	user, err := QueryUserByID(nil, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}
