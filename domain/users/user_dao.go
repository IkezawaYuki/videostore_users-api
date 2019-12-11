package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/logger"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"github.com/IkezawaYuki/videostore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, nick_name, email, age, date_created, status, password) VALUES(?,?,?,?,?,?,?,?);"
	querySelectUser = "SELECT id, first_name, last_name, nick_name, email, age, date_created, status, password FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, nick_name=?, email=?, age=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id = ?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, nick_name, email, age, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(querySelectUser)
	if err != nil {
		logger.Error("error where trying to prepare get user statement", err)
		return errors.NewInternalServerErr("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NickName,
		&user.Email,
		&user.Age,
		&user.DateCreated,
		&user.Status,
		&user.Password); getErr != nil{
		logger.Error("error where trying to get user by id", getErr)
		return errors.NewInternalServerErr("database error")
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.Age, user.DateCreated, user.Status, user.Password)
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
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next(){
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.NickName, &user.Email, &user.Age, &user.DateCreated, &user.Status)
		if  err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}

	if len(result) == 0{
		return nil, errors.NewNotFoundErr(fmt.Sprintf("no user matching status %s", status))
	}

	return result, nil
}