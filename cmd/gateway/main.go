package main

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app/server"
	"judgeMore_server/app/gateway/mw/jwt"
	"judgeMore_server/app/gateway/router"
	"judgeMore_server/app/gateway/rpc"
	"judgeMore_server/config"
	"judgeMore_server/pkg/constants"
	"judgeMore_server/pkg/utils"
)

var serviceName = constants.GatewayServiceName

func init() {
	config.Init(serviceName)
	jwt.Init()
	rpc.Init()
}

func main() {
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("get available port failed, err: %v", err)
	}

	h := server.New(
		server.WithHostPorts(listenAddr),
		server.WithHandleMethodNotAllowed(true),
	)

	router.GeneratedRegister(h)

	h.Spin()
}
