package router

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter(timeout time.Duration, gin *gin.Engine) {

	// All public APIs declared here (do not require auth)
	publicRouter := gin.Group("")
	NewPublicPlaceholderRouter(timeout, publicRouter)
	NewSignupRouter(timeout, publicRouter)

	// All protected APIs declared here (require auth)
	protectedRouter := gin.Group("")
	NewProtectedPlaceholderRouter(timeout, protectedRouter)
}
