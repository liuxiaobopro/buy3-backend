package main

import (
	"user_srv/initialize"
)

func main() {
	initialize.Config() // 初始化配置文件
	initialize.Logger() // 初始化日志
	initialize.DB()     // 初始化数据库
	initialize.Grpc()   // 初始化grpc服务
}
