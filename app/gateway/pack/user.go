package pack

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"judgeMore_server/app/gateway/model/model"
	rpc "judgeMore_server/kitex_gen/model"
	"strconv"
	"time"
)

func UserInfo(data *rpc.UserInfo) *model.UserInfo {
	hlog.Info(data)
	return &model.UserInfo{
		UserId:    data.UserId,
		Username:  data.Username,
		College:   data.College,
		Grade:     data.Grade,
		Major:     data.Major,
		Email:     data.Email,
		CreatedAt: ChangeFormat(data.CreatedAt),
		UpdatedAt: ChangeFormat(data.UpdatedAt),
	}
}
func ChangeFormat(timeStr string) string {
	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return ""
	}
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
