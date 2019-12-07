package users

import (
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
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
	Status 		string `json:"status"`
	Password    string `json:"-"`
}

func (adminUser *AdminUser) Validate() *errors.RestErr{
	adminUser.Email = strings.TrimSpace(strings.ToLower(adminUser.Email))
	if adminUser.Email == "" {
		return errors.NewBadRequestErr("Invalid email address")
	}
	if adminUser.UserID > 0{
		return errors.NewBadRequestErr("Invalid user id")
	}

	return nil
}