package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/api/router"
)

const maxTimeoutSeconds = 2

func main() {
	gin := gin.Default()

	timeout := maxTimeoutSeconds * time.Second
	router.SetupRouter(timeout, gin)
	gin.Run()
}
