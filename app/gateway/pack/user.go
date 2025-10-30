package pack

import (
	"judgeMore_server/app/gateway/model/model"
	rpc "judgeMore_server/kitex_gen/model"
)

func UserInfo(data *rpc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		UserId:   data.UserId,
		Username: data.Username,
		College:  data.College,
		Grade:    data.Grade,
		Major:    data.Major,
	}
}
