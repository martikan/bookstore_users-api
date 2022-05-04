package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martikan/bookstore_users-api/domain/user"
	"github.com/martikan/bookstore_users-api/errors"
	"github.com/martikan/bookstore_users-api/service"
)

// Controller for get all users
func GetAll(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// Controller for find a user by id
func Get(c *gin.Context) {

	userId, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := service.UserService.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

// Controller for search a user
func Search(c *gin.Context) {

	status := c.Query("status")

	users, err := service.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

// Controller for create a new user
func Create(c *gin.Context) {

	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid json body.")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := service.UserService.Create(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Controller for updae a user by id
func Update(c *gin.Context) {

	userId, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid json body.")
		c.JSON(restError.Status, restError)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := service.UserService.Update(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Controller for delete a user by id
func Delete(c *gin.Context) {

	userId, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if err := service.UserService.Delete(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Function to get the id of the user from path parameter
func getUserId(p string) (int64, *errors.RestError) {

	userId, userErr := strconv.ParseInt(p, 10, 64)
	if userErr != nil {
		return -1, errors.NewBadRequestError("User id should be a number.")
	}

	return userId, nil
}
