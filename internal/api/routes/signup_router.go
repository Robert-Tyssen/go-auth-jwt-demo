package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/api/controllers"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/data/repos"
)

func InitSignupRouter(timeout time.Duration, gr *gin.RouterGroup) {

	// Create repository and controller
	userRepo := repos.NewUserRepository()
	sc := controllers.NewSignupController(timeout, userRepo)

	// Endpoints for email / password signup and signin
	gr.POST("/signup", sc.Signup)
	gr.POST("/signin", sc.Signin)

}
