package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dayroMeli/bookstore_userss-api/domain/users"
	"github.com/dayroMeli/bookstore_userss-api/services"
	"github.com/dayroMeli/bookstore_userss-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invaled json body")
		c.JSON(restErr.Status, restErr)
		return

	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func GerUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}

	c.JSON(http.StatusOK, user)

}

func SearchUser(c *gin.Context) {}
