package users

import (
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
	"strings"
)

// 管理者ユーザーのEntity UserIDを外部キーに持つ。
type AdminUser struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"admin_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NickName    string `json:"nick_name"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type AdminUsers []AdminUser

func (adminUser *AdminUser) Validate() rest_errors.RestErr {
	adminUser.Email = strings.TrimSpace(strings.ToLower(adminUser.Email))
	if adminUser.Email == "" {
		return rest_errors.NewBadRequestError("Invalid email address")
	}
	if adminUser.UserID > 0 {
		return rest_errors.NewBadRequestError("Invalid user id")
	}
	adminUser.Password = strings.TrimSpace(strings.ToLower(adminUser.Password))
	if adminUser.Password == "" {
		return rest_errors.NewBadRequestError("Invalid password")
	}

	return nil
}
