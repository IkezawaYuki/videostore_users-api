package services

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)

var (
	AdminUserService adminUserService = adminUserService{}
)

type adminUserService struct {
}

type adminUserServiceInterface interface {
	GetAdminUser(int64)(*users.User, *errors.RestErr)
	CreateAdminUser(users.User)(*users.User, *errors.RestErr)
	UpdateAdminUser(bool, users.User)(*users.User, *errors.RestErr)
	DeleteAdminUser(int64) *errors.RestErr
	SearchAdminUser(string)(users.Users, *errors.RestErr)
}

// GetAdminUser 管理者ユーザーの取得
func (as *adminUserService)GetAdminUser(adminID int64)(*users.AdminUser, *errors.RestErr){
	result := &users.AdminUser{ID: adminID}
	if err := result.Get(); err != nil{
		return nil, err
	}
	return result, nil
}

// CreateAdminUser 管理者ユーザーの新規追加
func (as *adminUserService)CreateAdminUser(user users.AdminUser)(*users.AdminUser, *errors.RestErr){
	if err := user.Validate(); err != nil{
		return nil, err
	}
	if err := user.Save(); err != nil{
		return nil, err
	}
	return &user, nil
}

// UpdateAdminUser ユーザー情報の変更
func (as *adminUserService)UpdateAdminUser(isPartial bool, adminUser users.AdminUser)(*users.AdminUser, *errors.RestErr){
	current := &users.AdminUser{ID: adminUser.ID}

	if err := current.Get();err != nil{
		return nil, err
	}

	if isPartial {
		if adminUser.FirstName != ""{
			current.FirstName = adminUser.FirstName
		}
		if adminUser.LastName != ""{
			current.LastName = adminUser.LastName
		}
		if adminUser.NickName != ""{
			current.NickName = adminUser.NickName
		}
		if adminUser.Email != ""{
			current.Email = adminUser.Email
		}
		if adminUser.Age != 0{
			current.Age = adminUser.Age
		}

	} else {
		current.FirstName = adminUser.FirstName
		current.LastName = adminUser.LastName
		current.NickName = adminUser.NickName
		current.Email = adminUser.Email
		current.Age = adminUser.Age
	}

	if err := current.Update(); err != nil{
		return nil, err
	}
	return current, nil
}

// DeleteAdminUser ユーザー情報の削除
func (as *adminUserService)DeleteAdminUser(adminUserID int64) *errors.RestErr{
	adminUser := &users.AdminUser{ID: adminUserID}
	return adminUser.Delete()
}

// Search ステータスによるユーザーの検索
func (as *adminUserService)SearchAdminUser(status string)(users.AdminUsers, *errors.RestErr){
	dao := users.AdminUser{}
	return dao.FindByStatus(status)
}