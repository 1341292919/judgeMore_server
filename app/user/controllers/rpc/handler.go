package rpc

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"judgeMore_server/app/user/domain/model"
	"judgeMore_server/app/user/pack"
	"judgeMore_server/app/user/usecase"
	"judgeMore_server/kitex_gen/user"
	"judgeMore_server/pkg/errno"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	r = new(user.RegisterResponse)
	u := &model.User{
		Uid:      req.Id,
		UserName: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	var uid string
	if uid, err = h.useCase.RegisterUser(ctx, u); err != nil {
		r.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	r.UserId = &uid
	r.Base = pack.BuildBaseResp(errno.Success)
	return
}

func (h *UserHandler) Login(ctx context.Context, req *user.LoginRequest) (r *user.LoginResponse, err error) {
	r = new(user.LoginResponse)
	u := &model.User{
		Uid:      req.Id,
		Password: req.Password,
	}
	hlog.Info(u)
	userInfo, err := h.useCase.Login(ctx, u)
	if err != nil {
		r.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	r.Data = pack.User(userInfo)
	r.Base = pack.BuildBaseResp(errno.Success)
	hlog.Info(r)
	return
}

func (h *UserHandler) Logout(ctx context.Context, req *user.LogoutReq) (r *user.LogoutResp, err error) {
	return nil, nil
}

func (h *UserHandler) QueryUserInfo(ctx context.Context, req *user.QueryUserInfoRequest) (r *user.QueryUserInfoResponse, err error) {
	r = new(user.QueryUserInfoResponse)
	u := &model.User{
		Uid: req.UserId,
	}
	userInfo, err := h.useCase.QueryUserInfo(ctx, u)
	if err != nil {
		r.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	r.Data = pack.User(userInfo)
	r.Base = pack.BuildBaseResp(errno.Success)
	return
}

func (h *UserHandler) VerifyEmail(ctx context.Context, req *user.VerifyEmailRequest) (r *user.VerifyEmailResponse, err error) {
	r = new(user.VerifyEmailResponse)
	u := &model.EmailAuth{
		Email: req.Email,
		Code:  req.Code,
	}
	err = h.useCase.VerifyEmail(ctx, u)
	if err != nil {
		r.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	r.Base = pack.BuildBaseResp(errno.Success)
	return
}

func (h *UserHandler) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (r *user.UpdateUserInfoResponse, err error) {
	r = new(user.UpdateUserInfoResponse)
	u := &model.User{}
	if req.Grade != nil {
		u.Grade = *req.Grade
	}
	if req.Major != nil {
		u.Major = *req.Major
	}
	if req.College != nil {
		u.College = *req.College
	}
	u.Uid = req.Id
	userInfo, err := h.useCase.UpdateUserInfo(ctx, u)
	if err != nil {
		r.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	r.Data = pack.User(userInfo)
	r.Base = pack.BuildBaseResp(errno.Success)
	return
}
