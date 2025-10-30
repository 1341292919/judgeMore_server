package main

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"judgeMore_server/app/user"
	"judgeMore_server/config"
	"judgeMore_server/kitex_gen/user/userservice"
	"judgeMore_server/pkg/constants"
	"judgeMore_server/pkg/utils"
	"log"
	"net"
)

var serviceName = constants.UserServiceName

func init() {
	config.Init(serviceName)
}

func main() {

	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		logger.Fatalf("User: new etcd registry failed, err: %v", err)
	}
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("User: get available port failed, err: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", listenAddr) // 服务监听端口
	if err != nil {
		logger.Fatalf("User: resolve tcp addr failed, err: %v", err)
	}

	svr := userservice.NewServer(
		//只能注入一个handler
		user.InjectUserHandler(),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithRegistry(r), // 关键：注册到 ETCD
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName, // 关键：设置服务名称
		}),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
