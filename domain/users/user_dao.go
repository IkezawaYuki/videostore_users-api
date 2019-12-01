package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/utils/date_utils"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, nickname, email, age, date_created) VALUES(?,?,?,?,?,?)"
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.Age, user.DateCreated)
	if err != nil{
		return errors.NewInternalServerErr(
			fmt.Sprintf("error when trying to save user %s", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil{
		return errors.NewInternalServerErr(
			fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	user.ID = userID
	return nil
}
