package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/utils/date_utils"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"github.com/IkezawaYuki/videostore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, nick_name, email, age, date_created) VALUES(?,?,?,?,?,?);"
	querySelectUser = "SELECT id, first_name, last_name, nick_name, email, age, date_created FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, nick_name=?, email=?, age=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id = ?"
)

func (user *User) Get() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(querySelectUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NickName,
		&user.Email,
		&user.Age,
		&user.DateCreated); getErr != nil{
		return mysql_utils.ParseError(getErr)
	}
	fmt.Println(result)
	return nil
}

func (user *User) Save()*errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.Age, user.DateCreated)
	if saveErr != nil{
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil{
		return mysql_utils.ParseError(err)
	}
	user.ID = userID
	return nil
}

func (user *User) Update()*errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.NickName,user.Email, user.Age, user.ID)
	if err != nil{
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil{
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr){
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil{
		return nil, errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil{
		return nil, errors.NewInternalServerErr(err.Error())
	}
	rows.Close()

	result := make([]User, 0)
	// todo

	return result, nil
}