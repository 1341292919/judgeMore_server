package utils

import (
	"judgeMore_server/kitex_gen/model"
	"judgeMore_server/pkg/errno"
)

func IsSuccess(baseResp *model.BaseResp) bool {
	return baseResp.Code == errno.SuccessCode
}
