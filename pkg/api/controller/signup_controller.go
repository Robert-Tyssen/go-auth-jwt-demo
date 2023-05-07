package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUseCase
}

// Signup a new user using email and password
func (sc *SignupController) Signup(c *gin.Context) {

	var request domain.SignupRequest

	// Validate that request is properly constructed
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Use bcrypt to generate an encrypted password from the user-provided one
	// Only the hashed password is stored in the backend
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Construct user to be created
	var user = domain.User{
		Id:       uuid.New().String(),
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	// Create the user
	err = sc.SignupUsecase.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Create access token
	// TODO - implement access token secret and expiry from ENV
	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, "TODO_ACCESS_TOKEN_SECRET", 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Create refresh token
	// TODO - implement refresh token secret and expiry from ENV
	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, "TODO_REFRESH_TOKEN_SECRET", 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Create response and return success status code
	response := domain.SignupResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	c.JSON(http.StatusOK, response)
}
