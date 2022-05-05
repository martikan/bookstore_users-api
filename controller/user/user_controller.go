package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martikan/bookstore_users-api/domain/user"
	"github.com/martikan/bookstore_users-api/errors"
	"github.com/martikan/bookstore_users-api/service"
)

// GetAll Controller for get all users
func GetAll(c *gin.Context) {

	users, err := service.UserService.GetAll()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get Controller for find a user by id
func Get(c *gin.Context) {

	userId, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	usr, getErr := service.UserService.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, usr.Marshall(c.GetHeader("X-Public") == "true"))
}

// Search Controller for search a user
func Search(c *gin.Context) {

	status := c.Query("status")

	users, err := service.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

// Create Controller for create a new user
func Create(c *gin.Context) {

	var usr user.User

	if err := c.ShouldBindJSON(&usr); err != nil {
		restError := errors.NewBadRequestError("Invalid json body.")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := service.UserService.Create(usr)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Update Controller for update a user by id
func Update(c *gin.Context) {

	userId, userErr := getUserId(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var usr user.User

	if err := c.ShouldBindJSON(&usr); err != nil {
		restError := errors.NewBadRequestError("Invalid json body.")
		c.JSON(restError.Status, restError)
		return
	}

	usr.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := service.UserService.Update(isPartial, usr)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Delete Controller for delete a user by id
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
