package app

import (
	"github.com/gin-gonic/gin"
	"github.com/martikan/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Start the application...")
	router.Run(":8080")
}
