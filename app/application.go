package app

import (
	"github.com/dayroMeli/bookstore_userss-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApllication() {

	mapUrls()
	logger.Info("about to start the application.....")

	router.Run(":8080")

}
