package users

import (
	"net/http"
	"strconv"

	users "github.com/aravindvis/bookstore_users-api/domain/users"
	"github.com/aravindvis/bookstore_users-api/services"
	"github.com/aravindvis/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//CreateUser to create a user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser to get a user by id
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("Param User Id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, getError := services.GetUser(userID)

	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, result)
}

//SearchUsers searches for all the users
func SearchUsers(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Method implementation in progress")
}
