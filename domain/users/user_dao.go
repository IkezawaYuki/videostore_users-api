package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/logger"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
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

	return nil
}

func (user *User) Save()*errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error where trying to prepare save user statement", err)
		return errors.NewInternalServerErr("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.Age, user.DateCreated, user.Status, user.Password)
	if saveErr != nil{
		logger.Error("error where trying to save user by id", saveErr)
		return errors.NewInternalServerErr("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil{
		logger.Error("error where trying to get last insert id", err)
		return errors.NewInternalServerErr("database error")
	}
	user.ID = userID
	return nil
}

func (user *User) Update()*errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		logger.Error("error where trying to prepare update user statement", err)
		return errors.NewInternalServerErr("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.NickName,user.Email, user.Age, user.ID)
	if err != nil{
		logger.Error("error where trying to update user", err)
		return errors.NewInternalServerErr("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		logger.Error("error where trying to prepare delete user statement", err)
		return errors.NewInternalServerErr("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil{
		logger.Error("error where trying to delete user", err)
		return errors.NewInternalServerErr("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr){
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil{
		logger.Error("error where trying to prepare find users statement", err)
		return nil, errors.NewInternalServerErr("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil{
		logger.Error("error where trying to find users", err)
		return nil, errors.NewInternalServerErr("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next(){
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.NickName, &user.Email, &user.Age, &user.DateCreated, &user.Status)
		if  err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerErr("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0{
		return nil, errors.NewNotFoundErr(fmt.Sprintf("no user matching status %s", status))
	}

	return results, nil
}