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

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil{
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

// CreateUser ユーザー情報の登録
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
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
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

// DeleteUser ユーザー情報の削除
func DeleteUser(c *gin.Context){
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil{
		err := errors.NewBadRequestErr("user number should be number")
		c.JSON(err.Status, err)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil{
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}


func Search(c *gin.Context){
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context){
	var loginRequest users.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil{
		restErr := errors.NewBadRequestErr("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	users, err := services.UsersService.LoginUser(loginRequest)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}