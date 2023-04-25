package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/domain"
)

func NewProtectedPlaceholderRouter(timeout time.Duration, gr *gin.RouterGroup) {
	gr.GET("/protected-placeholder", handleProtectedPlaceholder)
}

func handleProtectedPlaceholder(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not Authorized"})
}
