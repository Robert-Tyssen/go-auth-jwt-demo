package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(timeout time.Duration, db *mongo.Client) *gin.Engine {

	// Create top-level router without any middleware
	r := gin.New()

	// Setup global middlewares
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	// All public APIs declared here (do not require auth)
	apiv1Public := r.Group("/api/v1")
	{
		apiv1Public.GET("/hello-world", helloWorldHandler)
		InitSignupRouter(timeout, apiv1Public, db)
	}

	// All protected APIs declared here (require auth)
	apiv1Private := r.Group("/api/v1", middleware.ValidateToken)
	{
		apiv1Private.GET("/test", middleware.NotImplemented)
		apiv1Private.GET("/hello-world-private", helloWorldPrivateHandler)
	}

	return r

}

func helloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Hello World, from Go JWT Demo!")
}

func helloWorldPrivateHandler(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Hello World, from Go JWT Demo! This is a private route.")
}
