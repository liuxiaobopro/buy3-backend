package main

import (
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"user_srv/global"
	"user_srv/handle"
	"user_srv/initialize"
	"user_srv/proto"
)

func main() {
	initialize.Config() // 初始化配置文件
	initialize.Logger() // 初始化日志
	initialize.DB()     // 初始化数据库

	zap.S().Info("user_srv ready to start")

	//#region 启动grpc
	grpcServer := grpc.NewServer()
	proto.RegisterUserServer(grpcServer, &handle.UserHandle{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", global.Conf.Srv.Ip, global.Conf.Srv.Port))
	if err != nil {
		panic(err)
	}

	zap.S().Infof("user_srv start at %s:%s", global.Conf.Srv.Ip, global.Conf.Srv.Port)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	//#endregion
}
