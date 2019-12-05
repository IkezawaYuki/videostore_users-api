package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/utils/date_utils"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, nick_name, email, age, date_created) VALUES(?,?,?,?,?,?);"
	querySelectUser = "SELECT id, first_name, last_name, nick_name, email, age, date_created FROM users WHERE id = ?;"
	indexUniqueEmail = "EMAIL"
	errorNoRows = "no rows in result"
)


func (user *User) Get() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(querySelectUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.NickName, &user.Email, &user.Age, &user.DateCreated); err != nil{
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NewNotFoundErr(fmt.Sprintf("user %d not found", user.ID))
		}

		fmt.Println(err)
		return errors.NewInternalServerErr(fmt.Sprintf("error when trying to get user %d: %s", user.ID, err.Error()))
	}

	return nil
}

func (user *User) Save()*errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.Age, user.DateCreated)
	if err != nil{
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), indexUniqueEmail){
			return errors.NewBadRequestErr(fmt.Sprintf("email %s already exists", user.Email))
		}
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
