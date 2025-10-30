package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"judgeMore_server/app/user/domain/model"
	"judgeMore_server/app/user/domain/repository"
	"judgeMore_server/pkg/constants"
	"judgeMore_server/pkg/errno"
)

type userDB struct {
	client *gorm.DB
}

func NewUserDB(client *gorm.DB) repository.UserDB {
	return &userDB{client: client}
}

func (db *userDB) IsUserExist(ctx context.Context, user *model.User) (bool, error) {
	var userInfo *User
	err := db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("role_id = ?", user.Uid).
		First(&userInfo).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //找到了说明用户存在
		return true, nil
	}
	if err != nil {
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query user: %v", err)
	}
	return false, nil
}
func (db *userDB) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	userInfo := &User{
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		RoleId:   user.Uid,
		Status:   0, //初始状态未激活
	}
	err := db.client.WithContext(ctx).
		Table(constants.TableUser).
		Create(userInfo).
		Error
	if err != nil {
		return 0, err
	}
	return userInfo.RoleId, nil
}

// 该函数调用前检验存在性
func (db *userDB) GetUserInfoByRoleId(ctx context.Context, role_id int64) (*model.User, error) {
	var userInfo *User
	_ = db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("role_id = ?", role_id).
		First(&userInfo).
		Error
	return &model.User{
		Uid:      userInfo.RoleId,
		UserName: userInfo.UserName,
		Grade:    userInfo.Grade,
		Major:    userInfo.Major,
		College:  userInfo.College,
		Status:   userInfo.Status,
		Role:     userInfo.UserRole,
	}, nil
}

func (db *userDB) UpdateInfoByRoleId(ctx context.Context, role_id int64, element ...string) (*model.User, error) {
	updateFields := make(map[string]interface{})
	for i, value := range element {
		if value == "" {
			continue // 跳过空值
		}
		switch i {
		case 0:
			updateFields["major"] = value
		case 1:
			updateFields["college"] = value
		case 2:
			updateFields["grade"] = value
		}
	}
	err := db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("role_id = ?", role_id).
		Updates(updateFields).
		Error
	if err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update userInfo: %v", err)
	}

	return db.GetUserInfoByRoleId(ctx, role_id)
}

func (db *userDB) ActivateUser(ctx context.Context, uid int64) error {
	err := db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("uid = ?", uid).
		Update("status", 1).
		Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to activate user: %v", err)
	}
	return nil
}
