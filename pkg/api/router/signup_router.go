package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/api/controller"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/repos"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/usecases"
)

func NewSignupRouter(timeout time.Duration, gr *gin.RouterGroup) {
	ur := repos.NewUserRepository()
	sc := &controller.SignupController{SignupUsecase: usecases.NewSignupUseCase(ur, timeout)}

	gr.POST("/signup", sc.Signup)
}
