package pack

import (
	"judgeMore_server/kitex_gen/model"
	"judgeMore_server/pkg/errno"
)

type Base struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Base `json:"base"`
}

func BuildBaseResp(err errno.ErrNo) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
