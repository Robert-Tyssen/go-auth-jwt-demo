package main

import (
	"time"

	"github.com/robert-tyssen/go-auth-jwt-demo/internal/api/routes"
)

const maxTimeoutSeconds = 2

func main() {

	timeout := maxTimeoutSeconds * time.Second

	r := routes.InitRouter(timeout)
	r.Run()
}
