package users

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/services"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUser 管理者ユーザーの取得
func GetAdminUser(c *gin.Context) {
	adminID, userErr := strconv.ParseInt(c.Param("admin_id"), 10, 64)
	if userErr != nil{
		err := errors.NewBadRequestErr("admin user number should be number")
		c.JSON(err.Status, err)
	}
	adminUser, getErr := services.AdminUsersService.GetAdminUser(adminID)
	if getErr != nil{
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, adminUser)
}

// CreateUser 管理者ユーザーの登録
func CreateAdminUser(c *gin.Context) {
	var adminUser users.AdminUser
	if err := c.ShouldBindJSON(&adminUser); err != nil{
		restErr := errors.NewBadRequestErr("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.AdminUsersService.CreateAdminUser(adminUser)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}


// UpdateUser ユーザー情報の変更
func UpdateAdminUser(c *gin.Context){
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil{
		err := errors.NewBadRequestErr("user number should be number")
		c.JSON(err.Status, err)
	}

	var adminUser users.AdminUser
	if err := c.ShouldBindJSON(&adminUser); err != nil {
		restErr := errors.NewBadRequestErr("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	adminUser.ID = userID
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.AdminUsersService.UpdateAdminUser(isPartial, adminUser)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)

}

// DeleteUser ユーザー情報の削除
func DeleteAdminUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestErr("user number should be number")
		c.JSON(err.Status, err)
		return
	}

	if err := services.AdminUsersService.DeleteAdminUser(userID); err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func SearchAdminUser(c *gin.Context){
	status := c.Query("status")
	users, err := services.AdminUsersService.SearchAdminUser(status)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}


func LoginAdminUser(c *gin.Context){
	var loginRequest users.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil{
		restErr := errors.NewBadRequestErr("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	users, err := services.AdminUsersService.LoginAdminUser(loginRequest)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

}