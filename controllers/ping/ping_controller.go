package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping テスト用
func Ping(c *gin.Context){
	c.String(http.StatusOK, "pong\n")
}