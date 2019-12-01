package app

import (
	"github.com/IkezawaYuki/videostore_users-api/controllers/ping"
	"github.com/IkezawaYuki/videostore_users-api/controllers/users"
)

func mapUrls(){
	router.GET("/ping", ping.Ping)

	// 一般ユーザー関連
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)

	// 管理者ユーザー関連
	router.GET("/admin/users/:user_id", users.GetAdminUser)
	router.POST("/admin/users", users.CreateAdminUser)
}
