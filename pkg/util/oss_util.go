package util

import (
	"bytes"
	"fmt"
	"gitee.com/derrickball/douyin/pkg/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
	"sync"
)

type OSSClient struct {
	bucket *oss.Bucket
}

func (c *OSSClient) UploadBytes(name string, data []byte) (string, error) {
	err := c.bucket.PutObject(config.Config.OSSConfig.PrefixPath+name, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return getURL(name), nil
}

func (c *OSSClient) UploadFile(name string, path string) (string, error) {
	fd, err := os.Open(path)
	if err != nil {
		return "", err
	}
	err = c.bucket.PutObject(config.Config.OSSConfig.PrefixPath+name, fd)
	if err != nil {
		return "", err
	}
	return getURL(name), nil
}

func getURL(name string) string {
	ossConfig := config.Config.OSSConfig
	return fmt.Sprintf("https://%s.%s/%s", ossConfig.BucketName, ossConfig.Endpoint, ossConfig.PrefixPath+name)
}

var (
	ossClient *OSSClient
	once      sync.Once
)

func GetOSSClient() (*OSSClient, error) {
	var err error
	once.Do(func() {
		var client *oss.Client
		ossConfig := config.Config.OSSConfig
		client, err = oss.New(ossConfig.Endpoint, ossConfig.AccessKeyId, ossConfig.AccessKeySecret)
		if err != nil {
			log.Println(err)
			return
		}
		bucket, err := client.Bucket(ossConfig.BucketName)
		if err != nil {
			log.Println(err)
			return
		}
		ossClient = &OSSClient{bucket: bucket}
	})
	return ossClient, err
}
