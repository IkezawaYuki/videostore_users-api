package users

import (
	"fmt"
	"github.com/IkezawaYuki/videostore_users-api/datasources/mysql/users_db"
	"github.com/IkezawaYuki/videostore_users-api/utils/date_utils"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"strings"
)

const (
	queryInsertAdminUser = "INSERT INTO admin_users(user_id, first_name, last_name, nick_name, email, age, date_created) VALUES(?,?,?,?,?,?,?)"
	querySelectAdminUser = "SELECT id, user_id, first_name, last_name, nick_name, email, age, date_created FROM admin_users WHERE id = ?;"
	indexUniqueAdminEmail = "ADMIN_EMAIL"
)

var (
	adminUserDB = make(map[int64]*AdminUser)
)

func (adminUser *AdminUser) Get() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(querySelectAdminUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(adminUser.ID)
	if err := result.Scan(&adminUser.ID,
		&adminUser.UserID,
		&adminUser.FirstName,
		&adminUser.LastName,
		&adminUser.NickName,
		&adminUser.Email,
		&adminUser.Age,
		&adminUser.DateCreated); err != nil{
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NewNotFoundErr(fmt.Sprintf("admin_user %d not found", adminUser.ID))
		}
		return errors.NewInternalServerErr(fmt.Sprintf("error when trying to get user %d: %s", adminUser.ID, err.Error()))
	}

	return nil
}


func (adminUser *AdminUser) Save()*errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryInsertAdminUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	adminUser.DateCreated = date_utils.GetNowString()


	insertResult, err := stmt.Exec(adminUser.UserID, adminUser.FirstName, adminUser.NickName, adminUser.Email, adminUser.Age, adminUser.DateCreated)
	if err != nil{
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), indexUniqueAdminEmail){
			return errors.NewInternalServerErr(fmt.Sprintf("email %s already exits", adminUser.Email))
		}
		return errors.NewInternalServerErr(
			fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	adminUserID, err := insertResult.LastInsertId()
	if err != nil{
		return errors.NewInternalServerErr(
			fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	adminUser.ID = adminUserID
	return nil


}
