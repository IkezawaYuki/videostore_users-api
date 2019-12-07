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
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)

	// 管理者ユーザー関連
	router.GET("/admin/users/:admin_id", users.GetAdminUser)
	router.POST("/admin/users", users.CreateAdminUser)

}
