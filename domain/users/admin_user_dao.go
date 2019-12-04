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
	indexUniqueAdminEmail = "ADMIN_EMAIL"
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
	stmt, err := users_db.Client.Prepare(queryInsertAdminUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	adminUser.DateCreated = date_utils.GetNowString()

	// todo ユーザー登録していないものは管理者ユーザーにはなれないようにする。

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
