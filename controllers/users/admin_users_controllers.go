package users

import (
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
	c.String(http.StatusNotImplemented, "implement me!\n")
}
