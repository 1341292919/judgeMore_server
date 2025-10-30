package service

import "judgeMore_server/app/user/domain/repository"

type UserService struct {
	db repository.UserDB
	ca repository.UserCache
}

func NewUserService(db repository.UserDB, ca repository.UserCache) *UserService {
	if db == nil {
		panic("userService`s db should not be nil")
	}
	if ca == nil {
		panic("userService`s cache should not be nil")
	}
	svc := &UserService{
		db: db,
		ca: ca,
	}
	return svc
}
