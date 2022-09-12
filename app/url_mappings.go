package app

import (
	"github.com/dayroMeli/bookstore_userss-api/controllers/ping"
	"github.com/dayroMeli/bookstore_userss-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	//router.GET("/users:user_id", users.SearchUser)
	router.GET("/users/:user_id", users.GerUser)
	router.POST("/users", users.CreateUser)
}
