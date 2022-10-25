package app

import (
	"github.com/dayroMeli/bookstore_userss-api/controllers/ping"
	"github.com/dayroMeli/bookstore_userss-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/internal/users/search", users.Search)
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

}
