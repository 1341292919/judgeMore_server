package pack

import (
	"judgeMore_server/app/user/domain/model"
	rpc "judgeMore_server/kitex_gen/model"
	"strconv"
)

func User(user *model.User) *rpc.UserInfo {
	return &rpc.UserInfo{
		Username:  user.UserName,
		UserId:    user.Uid,
		Major:     user.Major,
		College:   user.College,
		Grade:     user.Grade,
		Role:      user.Role,
		Email:     user.Email,
		CreatedAt: strconv.FormatInt(user.CreateAT, 10),
		UpdatedAt: strconv.FormatInt(user.UpdateAT, 10),
		DeletedAt: strconv.FormatInt(user.DeleteAT, 10),
	}
}
