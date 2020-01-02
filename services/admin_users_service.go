package services

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/utils/crypto_utils"
	"github.com/IkezawaYuki/videostore_users-api/utils/date_utils"
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
)

var (
	AdminUsersService adminUsersService = adminUsersService{}
)

type adminUsersService struct {
}

type adminUserServiceInterface interface {
	GetAdminUser(int64)(*users.User, *rest_errors.RestErr)
	CreateAdminUser(users.User)(*users.User, *rest_errors.RestErr)
	UpdateAdminUser(bool, users.User)(*users.User, *rest_errors.RestErr)
	DeleteAdminUser(int64) *rest_errors.RestErr
	SearchAdminUser(string)(users.Users, *rest_errors.RestErr)
}

// GetAdminUser 管理者ユーザーの取得
func (as *adminUsersService)GetAdminUser(adminID int64)(*users.AdminUser, *rest_errors.RestErr){
	result := &users.AdminUser{ID: adminID}
	if err := result.Get(); err != nil{
		return nil, err
	}
	return result, nil
}

// CreateAdminUser 管理者ユーザーの新規追加
func (as *adminUsersService)CreateAdminUser(adminUser users.AdminUser)(*users.AdminUser, *rest_errors.RestErr){
	if err := adminUser.Validate(); err != nil{
		return nil, err
	}
	adminUser.Status = users.StatusActive
	adminUser.DateCreated = date_utils.GetNowString()
	adminUser.Password = crypto_utils.GetMd5(adminUser.Password)

	if err := adminUser.Save(); err != nil{
		return nil, err
	}
	return &adminUser, nil
}

// UpdateAdminUser ユーザー情報の変更
func (as *adminUsersService)UpdateAdminUser(isPartial bool, adminUser users.AdminUser)(*users.AdminUser, *rest_errors.RestErr){
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
func (as *adminUsersService)DeleteAdminUser(adminUserID int64) *rest_errors.RestErr{
	adminUser := &users.AdminUser{ID: adminUserID}
	return adminUser.Delete()
}

// Search ステータスによるユーザーの検索
func (as *adminUsersService)SearchAdminUser(status string)(users.AdminUsers, *rest_errors.RestErr){
	dao := users.AdminUser{}
	return dao.FindByStatus(status)
}

func (as *adminUsersService) LoginAdminUser(request users.LoginRequest)(*users.AdminUser, *rest_errors.RestErr){
	dao := &users.AdminUser{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil{
		return nil, err
	}
	return dao, nil
}