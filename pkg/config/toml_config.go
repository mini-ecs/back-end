package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type TomlConfig struct {
	AppName      string
	Service      Service
	MySQL        MySQLConfig
	Log          LogConfig
	ImageStorage ImageStorage
	NodeInfo     NodeInfo
	Debug        bool
}

type Service struct {
	Port string
}

// MySQL相关配置
type MySQLConfig struct {
	Host        string
	Name        string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

// 日志保存地址
type LogConfig struct {
	Path  string
	Level string
}

// ImageStorage 镜像存储地址
type ImageStorage struct {
	FilePath string
}

type NodeInfo struct {
	Ip   string
	Port uint
}

var c TomlConfig

func init() {
	// 设置文件名
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("toml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

}
func GetConfig() TomlConfig {
	return c
}
