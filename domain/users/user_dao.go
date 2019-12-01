package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)
var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{
	result := userDB[user.ID]
	if result == nil{
		return errors.NewNotFoundErr(fmt.Sprintf("user_id %d is not found", user.ID))
	}
	user.NickName = result.NickName
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Age = result.Age
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save()*errors.RestErr{
	current := userDB[user.ID]
	if current != nil{
		if current.Email == user.Email{
			return errors.NewBadRequestErr(fmt.Sprintf("email %s is already registered", user.Email))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user_id %d is already exists", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
