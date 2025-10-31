package usecase

import (
	"context"
	"fmt"
	"judgeMore_server/app/user/domain/model"
	"judgeMore_server/pkg/crypt"
	"judgeMore_server/pkg/errno"
)

func (uc *useCase) RegisterUser(ctx context.Context, u *model.User) (uid string, err error) {
	//这边应该完成用户注册的几个步骤 1.参数检验、2.用户存在检验、3.密码哈希、4.db create new user
	exist, err := uc.db.IsUserExist(ctx, u)
	if err != nil {
		return "", fmt.Errorf("check user exist failed: %w", err)
	}
	if exist {
		return "", errno.NewErrNo(errno.ServiceUserExistCode, "user already exist")
	}
	u.Password, err = crypt.PasswordHash(u.Password)
	if err != nil {
		return "", fmt.Errorf("hash password failed: %w", err)
	}
	//验证邮箱
	err = uc.svc.SendEmail(ctx, u)
	if err != nil {
		return "", fmt.Errorf("send email failed: %w", err)
	}
	// 创建账户
	uid, err = uc.svc.CreateUser(ctx, u)
	if err != nil {
		return "", fmt.Errorf("create user failed: %w", err)
	}
	return uid, nil
}

func (uc *useCase) Login(ctx context.Context, u *model.User) (UserInfo *model.User, err error) {
	exist, err := uc.db.IsUserExist(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("check user exist failed: %w", err)
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserExistCode, "user not exist")
	}
	// 密码检验
	userInfo, err := uc.db.GetUserInfoByRoleId(ctx, u.Uid)
	if err != nil {
		return nil, fmt.Errorf("get user Info failed: %w", err)
	}
	// 激活检验
	if userInfo.Status == 0 {
		return nil, errno.NewErrNo(errno.ServiceUserDeathCode, "user not active ")
	}
	if !crypt.VerifyPassword(u.Password, userInfo.Password) {
		return nil, errno.Errorf(errno.ServiceUserPasswordError, "password not match")
	}
	return userInfo, nil
}

func (uc *useCase) QueryUserInfo(ctx context.Context, u *model.User) (UserInfo *model.User, err error) {
	//存在性检验
	exist, err := uc.db.IsUserExist(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("check user exist failed: %w", err)
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserExistCode, "user not exist")
	}
	userInfo, err := uc.db.GetUserInfoByRoleId(ctx, u.Uid)
	if err != nil {
		return nil, fmt.Errorf("get user Info failed: %w", err)
	}
	return userInfo, nil
}

func (uc *useCase) UpdateUserInfo(ctx context.Context, u *model.User) (UserInfo *model.User, err error) {
	// 由于uid读取自token 所以不做存在性检验
	userInfo, err := uc.svc.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (uc *useCase) VerifyEmail(ctx context.Context, data *model.EmailAuth) (err error) {
	return uc.svc.VerifyEmail(ctx, data)
}
