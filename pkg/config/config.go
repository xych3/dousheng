package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type MySQL struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	DSN      string
}

type Address struct {
	Etcd     string
	User     string
	Video    string
	Favorite string
	Chat     string
	Api      string
	Comment  string
	Relation string
}

type ServiceName struct {
	User     string
	Video    string
	Favorite string
	Comment  string
	Relation string
	Chat     string
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type OSSConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"`
	PrefixPath      string `yaml:"prefix_path"`
}

type ResourceConfig struct {
	TmpDir string `yaml:"tmp_dir"`
}

type Configs struct {
	MySQL       *MySQL `json:"mysql"`
	Address     *Address
	Service     *ServiceName
	RedisConfig *RedisConfig    `yaml:"redis"`
	OSSConfig   *OSSConfig      `yaml:"oss"`
	ResConfig   *ResourceConfig `yaml:"resource"`
}

var Config Configs

func LoadConfig() {
	LoadConfigWithPath("config-prod.yaml")
}

func LoadConfigWithPath(configPath string) {
	pwd, _ := os.Getwd()
	log.Println(pwd)
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
	}
	Config = Configs{}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalln(err)
	}
	Config.MySQL.DSN = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		Config.MySQL.Username, Config.MySQL.Password, Config.MySQL.Host, Config.MySQL.Port, Config.MySQL.Database)
	log.Printf("redis connect result: config.redis: %#v\n", Config.RedisConfig)
	log.Printf("mysql connect result: config.mysql: %#v\n", Config.MySQL)

}
