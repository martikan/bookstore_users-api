package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martikan/bookstore_users-api/domain/user"
	"github.com/martikan/bookstore_users-api/error"
	"github.com/martikan/bookstore_users-api/service/user_service"
)

// Controller for get all users
func GetAllUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// Controller for find a user by id
func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := error.NewBadRequestError("User id should be a number.")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := user_service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Controller for search a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// Controller for create a new user
func CreateUser(c *gin.Context) {

	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := error.NewBadRequestError("Invalid json body.")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := user_service.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Controller for delete a user by id
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
