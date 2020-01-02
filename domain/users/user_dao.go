package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/logger"
	"github.com/IkezawaYuki/videostore_users-api/utils/mysql_utils"
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, nick_name, email, age, date_created, status, password) VALUES(?,?,?,?,?,?,?,?);"
	querySelectUser = "SELECT id, first_name, last_name, nick_name, email, age, date_created, status, password FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, nick_name=?, email=?, age=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id = ?"
	queryFindByStatus = "SELECT id, first_name, last_name, nick_name, email, age, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, nick_name, email, age, date_created, status, password FROM users WHERE email=? AND password=?;"
)

func (user *User) Get() *rest_errors.RestErr{
	stmt, err := users_db.Client.Prepare(querySelectUser)
	if err != nil {
		logger.Error("error where trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
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
		return rest_errors.NewInternalServerError("database error", getErr)
	}

	return nil
}

func (user *User) Save()*rest_errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error where trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.Age, user.DateCreated, user.Status, user.Password)
	if saveErr != nil{
		logger.Error("error where trying to save user by id", saveErr)
		return rest_errors.NewInternalServerError("database error", err)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil{
		logger.Error("error where trying to get last insert id", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	user.ID = userID
	return nil
}

func (user *User) Update()*rest_errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		logger.Error("error where trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.NickName,user.Email, user.Age, user.ID)
	if err != nil{
		logger.Error("error where trying to update user", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	return nil
}

func (user *User) Delete() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		logger.Error("error where trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil{
		logger.Error("error where trying to delete user", err)
		return rest_errors.NewInternalServerError("database error", err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestErr){
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil{
		logger.Error("error where trying to prepare find users statement", err)
		return nil, rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil{
		logger.Error("error where trying to find users", err)
		return nil, rest_errors.NewInternalServerError("database error", err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next(){
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.NickName, &user.Email, &user.Age, &user.DateCreated, &user.Status)
		if  err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, rest_errors.NewInternalServerError("database error", err)
		}
		results = append(results, user)
	}

	if len(results) == 0{
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}

	return results, nil
}

func (user *User) FindByEmailAndPassword() *rest_errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error where trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password)
	if getErr := result.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NickName,
		&user.Email,
		&user.Age,
		&user.DateCreated,
		&user.Status,
		&user.Password); getErr != nil{

		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows){
			return rest_errors.NewInternalServerError("users not found", err)
		}
		logger.Error("error where trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("database error", err)
	}

	return nil
}
