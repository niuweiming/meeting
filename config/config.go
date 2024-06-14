package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Config *config

// 总配文件
type config struct {
	Db db `yaml:"db"`
}

// 数据库的配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// 配置初始化
func InitConfig() {
	yamlFile, err := os.ReadFile("./config/config.yaml")
	//有错就down机
	if err != nil {
		panic(err)
	}
	//绑定值
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
}

// 测试文件调用
func InitConfigtest() {
	yamlFile, err := os.ReadFile("../config/config.yaml")
	//有错就down机
	if err != nil {
		panic(err)
	}
	//绑定值
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
}
