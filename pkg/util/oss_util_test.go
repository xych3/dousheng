package util

import (
	"fmt"
	"gitee.com/derrickball/douyin/pkg/config"
	"os"
	"testing"
)

func TestGetOSSClient(t *testing.T) {
	config.LoadConfigWithPath("../../config-dev.yaml")
	client, err := GetOSSClient()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(client)
}

func TestOSSClient_UploadBytes(t *testing.T) {
	config.LoadConfigWithPath("../../config-dev.yaml")
	client, _ := GetOSSClient()
	bytes, err := os.ReadFile("../../res/bear.mp4")
	if err != nil {
		t.Fatal(err)
	}
	url, err := client.UploadBytes("douyin.mp4", bytes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(url)
}

func TestOSSClient_UploadFile(t *testing.T) {
	config.LoadConfigWithPath("../../config-dev.yaml")
	client, _ := GetOSSClient()
	url, err := client.UploadFile("aaaaa.mp4", config.Config.ResConfig.TmpDir+"bear.mp4")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(url)
}
