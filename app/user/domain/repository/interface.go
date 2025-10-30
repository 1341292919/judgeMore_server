package repository

import (
	"context"
	"judgeMore_server/app/user/domain/model"
)

type UserDB interface {
	IsUserExist(ctx context.Context, user *model.User) (bool, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	GetUserInfoByRoleId(ctx context.Context, role_id int64) (*model.User, error)
	UpdateInfoByRoleId(ctx context.Context, role_id int64, element ...string) (*model.User, error)
	ActivateUser(ctx context.Context, uid int64) error
}
type UserCache interface {
	GetCodeCache(ctx context.Context, key string) (code string, err error)
	PutCodeToCache(ctx context.Context, key string) (code string, err error)
	IsKeyExist(ctx context.Context, key string) bool
}
