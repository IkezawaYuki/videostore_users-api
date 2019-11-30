package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser 管理者ユーザーの取得
func GetAdminUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!\n")
}

// CreateUser 管理者ユーザーの登録
func CreateAdminUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!\n")
}
