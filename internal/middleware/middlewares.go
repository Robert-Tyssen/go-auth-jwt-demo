package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Not Implemented!")
}