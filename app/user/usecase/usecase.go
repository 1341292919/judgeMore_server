package usecase

import (
	"context"
	"judgeMore_server/app/user/domain/model"
	"judgeMore_server/app/user/domain/repository"
	"judgeMore_server/app/user/domain/service"
)

// UserUseCase 接口应该不应该定义在 domain 中，这属于 use case 层
type UserUseCase interface {
	RegisterUser(ctx context.Context, user *model.User) (uid string, err error)
	Login(ctx context.Context, user *model.User) (userInfo *model.User, err error)
	QueryUserInfo(ctx context.Context, user *model.User) (userInfo *model.User, err error)
	UpdateUserInfo(ctx context.Context, user *model.User) (userInfo *model.User, err error)
	VerifyEmail(ctx context.Context, data *model.EmailAuth) (err error)
}

// useCase 实现了 domain.UserUseCase
// 只会以接口的形式被调用, 所以首字母小写改为私有类型
type useCase struct {
	db  repository.UserDB
	svc *service.UserService
	ca  repository.UserCache
}

func NewUserUseCase(db repository.UserDB, svc *service.UserService, ca repository.UserCache) *useCase {
	return &useCase{
		db:  db,
		svc: svc,
		ca:  ca,
	}
}
