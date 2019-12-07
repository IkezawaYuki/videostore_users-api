package users

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/services"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)



// GetUser ユーザー情報の取得
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil{
		err := errors.NewBadRequestErr("user number should be number")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil{
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser ユーザー情報の登録
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdateUser ユーザー情報の変更
func UpdateUser(c *gin.Context){
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
	result, err := services.UpdateUser(isPartial, user)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)

}

// DeleteUser ユーザー情報の削除
func DeleteUser(c *gin.Context){
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil{
		err := errors.NewBadRequestErr("user number should be number")
		c.JSON(err.Status, err)
		return
	}

	if err := services.DeleteUser(userID); err != nil{
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}