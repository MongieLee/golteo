package config

import (
	"github.com/spf13/viper"
	"log"
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

type RabbitMqConfig struct {
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	VirtualHost string `mapstructure:"virtual_host"`
}

var CustomConfig struct {
	AppDebug   bool           `mapstructure:"app_debug"`
	Mysql      MysqlConfig    `mapstructure:"mysql"`
	Redis      RedisConfig    `mapstructure:"redis"`
	RateConfig RateConfig     `mapstructure:"rate"`
	Rmq        RabbitMqConfig `mapstructure:"rmq"`
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
		log.Fatalf("Viper读取配置是啊比，错误信息：%v", err)
	}
	err = viper.Unmarshal(&CustomConfig)
	if err != nil {
		log.Fatalf("Viper反解析Json失败，错误信息：%v", err)
	}
}
