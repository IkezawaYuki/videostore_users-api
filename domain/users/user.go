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
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestErr("Invalid email address")
	}

	if user.NickName == ""{
		return errors.NewBadRequestErr("Invalid nickname")
	}

	if user.Age > 0{
		return errors.NewBadRequestErr("Invalid age")
	}

	return nil
}