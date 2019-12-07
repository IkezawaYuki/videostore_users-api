package users

import (
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"strings"
)

// User 一般ユーザーのEntity
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NickName    string `json:"nick_name"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status 		string `json:"status"`
	Password    string `json:"-"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.NickName = strings.TrimSpace(strings.ToLower(user.NickName))

	if user.Email == "" {
		return errors.NewBadRequestErr("Invalid email address")
	}

	return nil
}