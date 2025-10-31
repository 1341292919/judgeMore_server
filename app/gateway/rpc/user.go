package rpc

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	api "judgeMore_server/app/gateway/model/api/user"
	"judgeMore_server/app/gateway/pack"
	"judgeMore_server/kitex_gen/user"
	"judgeMore_server/pkg/base/client"
	"judgeMore_server/pkg/errno"
	"judgeMore_server/pkg/utils"
)

func InitUserRPC() {
	c, err := client.InitUserRPC()
	if err != nil {
		logger.Fatalf("api.rpc.user InitUserRPC failed, err is %v", err)
	}
	userClient = *c
}

func RegisterRPC(ctx context.Context, req *user.RegisterRequest) (apiResp *api.RegisterResponse, err error) {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		logger.Errorf("RegisterRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	apiResp = &api.RegisterResponse{UserId: *resp.UserId}
	return apiResp, nil
}
func LoginRPC(ctx context.Context, req *user.LoginRequest) (apiResp *api.LoginResponse, err error) {
	apiResp = new(api.LoginResponse)
	resp, err := userClient.Login(ctx, req)
	if err != nil {
		logger.Errorf("LoginRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	apiResp.Data = pack.UserInfo(resp.Data)
	return apiResp, nil
}
func LogoutRPC(ctx context.Context, req *user.LogoutReq) error {
	resp, err := userClient.Logout(ctx, req)
	if err != nil {
		logger.Errorf("LogoutRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return nil
}
func QueryUserInfoRpc(ctx context.Context, req *user.QueryUserInfoRequest) (apiResp *api.LoginResponse, err error) {
	resp, err := userClient.QueryUserInfo(ctx, req)
	if err != nil {
		logger.Errorf("QueryUserInfo: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	apiResp.Data = pack.UserInfo(resp.Data)
	return apiResp, nil
}
func UpdateUserInfoRpc(ctx context.Context, req *user.UpdateUserInfoRequest) (apiResp *api.UpdateUserInfoResponse, err error) {
	apiResp = new(api.UpdateUserInfoResponse)
	resp, err := userClient.UpdateUserInfo(ctx, req)
	if err != nil {
		logger.Errorf("QueryUserInfo: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	apiResp.Data = pack.UserInfo(resp.Data)
	return apiResp, nil
}
func VerifyEmailRpc(ctx context.Context, req *user.VerifyEmailRequest) error {
	resp, err := userClient.VerifyEmail(ctx, req)
	if err != nil {
		logger.Errorf("VerifyEmail: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return nil
}
