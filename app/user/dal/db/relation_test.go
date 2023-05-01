package db

import (
	"fmt"
	"gitee.com/derrickball/douyin/pkg/config"
	"log"
	"testing"
)

func TestCheckFollow(t *testing.T) {
	config.LoadConfig()
	Init()
	res, err := CheckFollow(nil, 2, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
