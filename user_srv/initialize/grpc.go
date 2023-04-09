package initialize

import (
	"fmt"
	"net"
	"strconv"

	"user_srv/global"
	"user_srv/handle"
	"user_srv/proto"

	consulApi "github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func Grpc() {
	srvIp := global.Conf.Srv.Ip
	srvPost, _ := strconv.Atoi(global.Conf.Srv.Port)

	srvAddr := fmt.Sprintf("%s:%d", srvIp, srvPost)
	consulAddr := fmt.Sprintf("%s:%s", global.Conf.ConsulInfo.Host, global.Conf.ConsulInfo.Port)

	//#region 注册
	grpcServer := grpc.NewServer()
	proto.RegisterUserServer(grpcServer, &handle.UserHandle{})
	lis, err := net.Listen("tcp", srvAddr)
	if err != nil {
		panic(err)
	}
	//#endregion

	//#region 健康检查
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	//#endregion

	//#region 服务注册
	defaultConfig := consulApi.DefaultConfig()
	defaultConfig.Address = consulAddr

	client, err := consulApi.NewClient(defaultConfig)
	if err != nil {
		panic(err)
	}
	registration := consulApi.AgentServiceRegistration{
		ID:      global.Conf.Name,
		Name:    global.Conf.Name,
		Tags:    []string{global.Conf.Name},
		Address: "192.168.56.1", // 服务对应的地址
		Port:    srvPost,        // 服务对应的端口号
		Check: &consulApi.AgentServiceCheck{
			GRPC:                           "192.168.56.1:50001", // 服务对应的地址和端口号
			Interval:                       "1s",                 // 监控检测间隔时间
			Timeout:                        "3s",                 // 监控检测超时时间
			DeregisterCriticalServiceAfter: "1000s",              // 服务检测失败后,注销服务的时间
		},
	}
	if err = client.Agent().ServiceRegister(&registration); err != nil {
		panic(err)
	}
	//#endregion

	//#region 启动
	zap.S().Infof("user_srv start at %s:%s", global.Conf.Srv.Ip, global.Conf.Srv.Port)
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	//#endregion
}
