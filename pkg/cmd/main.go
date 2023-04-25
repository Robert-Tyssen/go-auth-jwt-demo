package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/api/router"
)

func main() {
	gin := gin.Default()
	router.SetupRouter(gin)
	gin.Run()
}
