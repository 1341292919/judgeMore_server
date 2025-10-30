package client

import (
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"judgeMore_server/config"
	"judgeMore_server/kitex_gen/user/userservice"
	"judgeMore_server/pkg/constants"
)

func initRPCClient[T any](serviceName string, newClientFunc func(string, ...client.Option) (T, error)) (*T, error) {
	if config.Etcd == nil || config.Etcd.Addr == "" {
		return nil, errors.New("config.Etcd.Addr is nil")
	}
	//初始化 Etcd
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		return nil, fmt.Errorf("initRPCClient etcd.NewEtcdResolver failed: %w", err)
	}
	client, err := newClientFunc(serviceName,
		client.WithResolver(r),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: fmt.Sprintf(constants.KitexClientEndpointInfoFormat, serviceName)}),
	)
	if err != nil {
		return nil, fmt.Errorf("initRPCClient NewClient failed: %w", err)
	}
	return &client, nil
}

func InitUserRPC() (*userservice.Client, error) {
	return initRPCClient("user", userservice.NewClient)
}
