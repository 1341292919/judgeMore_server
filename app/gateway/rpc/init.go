package rpc

import "judgeMore_server/kitex_gen/user/userservice"

var (
	userClient userservice.Client
)

func Init() {
	InitUserRPC()
}
