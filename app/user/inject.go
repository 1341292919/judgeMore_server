package user

import (
	"judgeMore_server/app/user/controllers/rpc"
	"judgeMore_server/app/user/domain/service"
	"judgeMore_server/app/user/infrastructure/cache"
	"judgeMore_server/app/user/infrastructure/mysql"
	"judgeMore_server/app/user/usecase"
	"judgeMore_server/kitex_gen/user"
	"judgeMore_server/pkg/base/client"
	"judgeMore_server/pkg/constants"
)

func InjectUserHandler() user.UserService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	re, err := client.NewRedisClient(constants.RedisDBUser) // 使用和网关同一个数据库，目前仅用作登录登出
	if err != nil {
		panic(err)
	}
	db := mysql.NewUserDB(gormDB)
	ca := cache.NewUserCache(re)
	svc := service.NewUserService(db, ca)
	uc := usecase.NewUserUseCase(db, svc, ca)
	return rpc.NewUserHandler(uc)
}
