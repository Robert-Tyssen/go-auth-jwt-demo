package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/api/controllers"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/data/repos"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitSignupRouter(timeout time.Duration, gr *gin.RouterGroup, db *mongo.Client) {

	// Create repository and controller
	userRepo := repos.NewUserRepository(db)
	sc := controllers.NewSignupController(timeout, userRepo)

	// Endpoints for email / password signup and signin
	gr.POST("/signup", sc.Signup)
	gr.POST("/signin", sc.Signin)

}
