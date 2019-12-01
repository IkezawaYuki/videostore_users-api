package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/utils/date_utils"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)

var (
	adminUserDB = make(map[int64]*AdminUser)
)

func (adminUser *AdminUser) Get() *errors.RestErr{
	result := adminUserDB[adminUser.ID]
	if result == nil{
		return errors.NewNotFoundErr(fmt.Sprintf("user_id %d is not found", adminUser.ID))
	}
	adminUser.UserID = result.UserID
	adminUser.NickName = result.NickName
	adminUser.FirstName = result.FirstName
	adminUser.LastName = result.LastName
	adminUser.Age = result.Age
	adminUser.Email = result.Email
	adminUser.DateCreated = result.DateCreated
	return nil
}


func (adminUser *AdminUser) Save()*errors.RestErr{
	current := adminUserDB[adminUser.ID]
	if current != nil{
		if current.Email == adminUser.Email{
			return errors.NewBadRequestErr(fmt.Sprintf("email %s is already registered", adminUser.Email))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user_id %d is already exists", adminUser.ID))
	}

	user := userDB[adminUser.UserID]
	if user == nil{
		return errors.NewBadRequestErr("cannot become an administrator without registering as a user")
	}
	adminUser.DateCreated = date_utils.GetNowString()

	adminUserDB[adminUser.ID] = adminUser
	return nil
}
