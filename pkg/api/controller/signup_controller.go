package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/domain"
)

type SignupController struct {
	SignupUsecase domain.SignupUseCase
}

func (sc *SignupController) SignupEmail(c *gin.Context) {

	var user = domain.User{
		Email:    "test@example.com",
		Password: "123456",
	}

	err := sc.SignupUsecase.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	resp := domain.SignupEmailResponse{
		AccessToken:  "ACCESS-TOKEN-NOT-IMPLEMENTED",
		RefreshToken: "REFRESH-TOKEN-NOT-IMPLEMENTED",
	}

	c.JSON(http.StatusOK, resp)
}
