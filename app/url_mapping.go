package app

import (
	"github.com/martikan/bookstore_users-api/controller/user"
)

func mapUrls() {

	// Routes for users endpoints
	router.GET("/users", user.GetAllUsers)
	router.GET("/users/:id", user.GetUser)
	router.GET("/users/search", user.SearchUser)
	router.POST("/users", user.CreateUser)
	router.DELETE("/users/:id", user.DeleteUser)
}
