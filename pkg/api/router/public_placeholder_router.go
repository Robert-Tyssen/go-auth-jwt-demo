package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewPublicPlaceholderRouter(timeout time.Duration, gr *gin.RouterGroup) {
	gr.GET("/public-placeholder", _handler)
}

func _handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message" : "success"})
}
