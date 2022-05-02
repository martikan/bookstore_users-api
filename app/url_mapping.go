package app

import (
	"github.com/martikan/bookstore_users-api/controller/user"
)

func mapUrls() {

	// Routes for users endpoints
	router.GET("/users", user.GetAll)
	router.GET("/internal/users/search", user.Search)
	router.GET("/users/:id", user.Get)
	router.POST("/users", user.Create)
	router.PUT("/users/:id", user.Update)
	router.PATCH("/users/:id", user.Update)
	router.DELETE("/users/:id", user.Delete)
}
