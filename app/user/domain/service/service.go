package service

import (
	"context"
	"errors"
	"fmt"
	"judgeMore_server/app/user/domain/model"
	"judgeMore_server/pkg/constants"
	"judgeMore_server/pkg/errno"
	"judgeMore_server/pkg/utils"
	"strconv"
)

func (svc *UserService) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	return svc.db.CreateUser(ctx, user)
}

func (svc *UserService) SendEmail(ctx context.Context, user *model.User) error {
	// 首先进行验证 学号即Uid 与fzu邮箱强绑定
	Correct := user.Email == strconv.FormatInt(user.Uid, 10)+constants.EmailSuffix
	if !Correct {
		return errno.NewErrNo(errno.ServiceEmailIncorrectCode, "Uid do not match email")
	}
	key := fmt.Sprintf("Email:%s", user.Email)
	code, err := svc.ca.PutCodeToCache(ctx, key)
	if err != nil {
		return err
	}
	// 发送邮箱
	err = utils.MailSendCode(user.Email, code)
	if err != nil {
		return err
	}
	return nil
}
func (svc *UserService) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// 由于这里需要对需更新的内容做选择 在svc处处理
	var updateParams []string

	if user.Major != "" {
		updateParams = append(updateParams, user.Major)
	}
	if user.College != "" {
		updateParams = append(updateParams, user.College)
	}
	if user.Grade != "" {
		updateParams = append(updateParams, user.Grade)
	}

	// 如果有需要更新的字段才执行
	if len(updateParams) > 0 {
		return svc.db.UpdateInfoByRoleId(ctx, user.Uid, updateParams...)
	}

	return nil, errno.Errorf(errno.InternalServiceErrorCode, "no element to update")

}

func (svc *UserService) VerifyEmail(ctx context.Context, data *model.EmailAuth) (err error) {
	// 判断存不存在
	key := fmt.Sprintf("Email:%s", data.Email)
	exist := svc.ca.IsKeyExist(ctx, key)
	if !exist {
		return errors.New("code expired")
	}
	code, err := svc.ca.GetCodeCache(ctx, key)
	if err != nil {
		return err
	}
	if code != data.Code {
		return errors.New("code not match")
	}
	// 更新user表的信息
	return svc.db.ActivateUser(ctx, data.Uid)
}
