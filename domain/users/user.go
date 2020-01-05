package users

import (
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
	"strings"
)

const (
	StatusActive = "active"
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
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.NickName = strings.TrimSpace(strings.ToLower(user.NickName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Password == "" {
		return rest_errors.NewBadRequestError("Invalid password")
	}

	return nil
}
