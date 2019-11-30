package users

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/services"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser ユーザーの取得
func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement GetUser()!\n")
}

// CreateUser ユーザーの登録
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "Invalid json body",
			Status:   http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusNotImplemented, result)
}

//// FindUser ユーザーの検索
////func FindUser(c *gin.Context){
////	c.String(http.StatusNotImplemented, "implement me!")
////}
