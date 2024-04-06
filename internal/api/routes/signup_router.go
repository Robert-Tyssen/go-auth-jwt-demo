package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/api/controllers"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/data/repos"
)

func NewSignupRouter(timeout time.Duration, gr *gin.RouterGroup) {

	userRepo := repos.NewUserRepository()
	sc := controllers.NewSignupController(timeout, userRepo)

	gr.POST("/signup", sc.Signup)
	
}
