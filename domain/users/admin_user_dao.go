package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/utils/date_utils"
	"github.com/IkezawaYuki/videostore_users-api/utils/mysql_utils"
	"github.com/IkezawaYuki/videostore_utils-go/logger"
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
	"strings"
)

const (
	queryInsertAdminUser                 = "INSERT INTO admin_users(user_id, first_name, last_name, nick_name, email, age, date_created) VALUES(?,?,?,?,?,?,?)"
	querySelectAdminUser                 = "SELECT id, user_id, first_name, last_name, nick_name, email, age, date_created FROM admin_users WHERE id = ?;"
	queryUpdateAdminUser                 = "UPDATE admin_users SET first_name=?, last_name=?, nick_name=?, email=?, age=? WHERE id=?;"
	queryDeleteAdminUser                 = "DELETE FROM admin_users WHERE id = ?"
	queryFindAdminUserByStatus           = "SELECT id, user_id, first_name, last_name, nick_name, email, age, date_created, status FROM admin_users WHERE status = ?;"
	queryFindAdminUserByEmailAndPassword = "SELECT id, user_id, first_name, last_name, nick_name, email, age, date_created, status, password FROM admin_users WHERE email=? AND password=?;"
	indexUniqueAdminEmail                = "ADMIN_EMAIL"
)

var (
	adminUserDB = make(map[int64]*AdminUser)
)

func (adminUser *AdminUser) Get() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(querySelectAdminUser)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(adminUser.ID)
	if getErr := result.Scan(&adminUser.ID,
		&adminUser.UserID,
		&adminUser.FirstName,
		&adminUser.LastName,
		&adminUser.NickName,
		&adminUser.Email,
		&adminUser.Age,
		&adminUser.DateCreated); getErr != nil {
		fmt.Println(getErr)
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (adminUser *AdminUser) Save() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertAdminUser)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	defer stmt.Close()
	adminUser.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(adminUser.UserID,
		adminUser.FirstName, adminUser.NickName, adminUser.Email, adminUser.Age, adminUser.DateCreated)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), indexUniqueAdminEmail) {
			return rest_errors.NewInternalServerError(fmt.Sprintf("email %s already exits", adminUser.Email), err)
		}
		return rest_errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user %s", err.Error()), err)
	}
	adminUserID, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user %s", err.Error()), err)
	}
	adminUser.ID = adminUserID

	return nil
}

func (adminUser *AdminUser) Update() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateAdminUser)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to update admin user", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(adminUser.FirstName, adminUser.LastName, adminUser.NickName, adminUser.Email, adminUser.Age, adminUser.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (adminUser *AdminUser) Delete() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteAdminUser)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to delete admin user", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(adminUser.ID); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (adminUser *AdminUser) FindByStatus(status string) ([]AdminUser, rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindAdminUserByStatus)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to find admin user", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to find admin user", err)
	}
	defer rows.Close()

	result := make([]AdminUser, 0)
	for rows.Next() {
		var adminUser AdminUser
		err := rows.Scan(&adminUser.ID, &adminUser.FirstName, &adminUser.LastName, &adminUser.NickName, &adminUser.Email, &adminUser.Age, &adminUser.DateCreated, &adminUser.Status)
		if err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, adminUser)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no admin user matching status %s", status))
	}

	return result, nil
}

func (adminUser *AdminUser) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error where trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerError("error when trying to find admin user", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(adminUser.Email, adminUser.Password)
	if getErr := result.Scan(&adminUser.ID,
		&adminUser.UserID,
		&adminUser.FirstName,
		&adminUser.LastName,
		&adminUser.NickName,
		&adminUser.Email,
		&adminUser.Age,
		&adminUser.DateCreated,
		&adminUser.Status,
		&adminUser.Password); getErr != nil {

		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewInternalServerError("admin users not found", err)
		}
		logger.Error("error where trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("error when trying to find admin user", err)
	}

	return nil
}
