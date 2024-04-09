package config

import (
	flag "github.com/spf13/pflag"
)

type Flag struct {
	Path string
}

var FlagConfig Flag

func InitFlag() {
	// 传入一个字符串指针，指定参数的全名，参数名的简写形式，参数的默认值，参数的使用说明
	// 这里会读取执行命令时 命令行传入的-path或者-p的值，如果没有传入，则默认值为./
	flag.StringVarP(&FlagConfig.Path, "path", "p", "./", "yaml config path")
	flag.Parse()
}
