package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/middleware"
)

func InitRouter(timeout time.Duration) *gin.Engine {

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
	}

	//NewSignupRouter(timeout, publicRouter)

	// All protected APIs declared here (require auth)
	apiv1Private := r.Group("/api/v1")
	{
		apiv1Private.GET("/test", middleware.NotImplemented)
	}

	return r

}

func helloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Hello World, from Go JWT Demo!")
}
