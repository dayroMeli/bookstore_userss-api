package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApllication() {

	mapUrls()
	router.Run(":8080")

}
