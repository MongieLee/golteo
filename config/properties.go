package config

import (
	"github.com/spf13/viper"
	"os"
)

type MysqlConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type RateConfig struct {
	Enable bool `mapstructure:"enable"`
	Limit  int  `mapstructure:"limit"`
	Burst  int  `mapstructure:"burst"`
}

var CustomConfig struct {
	Mysql      MysqlConfig `mapstructure:"mysql"`
	Redis      RedisConfig `mapstructure:"redis"`
	RateConfig RateConfig  `mapstructure:"rate"`
}

func InitViperConfig() {
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置文件的格式
	viper.SetConfigType("yaml")
	// 设置查找配置文件的路径为当前路径 . 表示项目的工作目录，也就是main.go同级的那个目录
	viper.AddConfigPath(FlagConfig.Path)

	// 读取配置文件中的数据到viper中
	err := viper.ReadInConfig()
	if err != nil {
		os.Exit(1)
	}
	err = viper.Unmarshal(&CustomConfig)
	if err != nil {
		os.Exit(1)
	}
}
