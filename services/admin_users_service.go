package services

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)

// GetAdminUser 管理者ユーザーの取得
func GetAdminUser(adminID int64)(*users.AdminUser, *errors.RestErr){
	result := &users.AdminUser{ID: adminID}
	if err := result.Get(); err != nil{
		return nil, err
	}
	return result, nil
}

// CreateAdminUser 管理者ユーザーの新規追加
func CreateAdminUser(user users.AdminUser)(*users.AdminUser, *errors.RestErr){
	if err := user.Validate(); err != nil{
		return nil, err
	}
	if err := user.Save(); err != nil{
		return nil, err
	}
	return &user, nil
}
