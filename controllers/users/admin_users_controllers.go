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
	adminUser, getErr := services.GetAdminUser(adminID)
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
	result, saveErr := services.CreateAdminUser(adminUser)
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

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UpdateAdminUser(isPartial, user)
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

	if err := services.DeleteAdminUser(userID); err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

