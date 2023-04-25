package router

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter(timeout time.Duration, gin *gin.Engine) {

	publicRouter := gin.Group("")
	NewPublicPlaceholderRouter(timeout, publicRouter)

	protectedRouter := gin.Group("")
	NewProtectedPlaceholderRouter(timeout, protectedRouter)
}
