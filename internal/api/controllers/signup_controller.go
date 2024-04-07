package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/data/repos"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/models"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/utils/password"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/utils/tokens"
)

type SignupController struct {
	timeout  time.Duration
	userRepo repos.UserRepository
}

func NewSignupController(timeout time.Duration, ur repos.UserRepository) *SignupController {
	return &SignupController{
		timeout:  timeout,
		userRepo: ur,
	}
}

// Signup a new user using email and password
func (sc *SignupController) Signup(c *gin.Context) {

	_, cancel := context.WithTimeout(c, sc.timeout)
	defer cancel()

	var request models.SignupRequest

	// Validate that request is properly constructed
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Use bcrypt to generate an encrypted password from the user-provided one
	// Only the hashed password is stored in the backend
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Construct user to be created
	var user = models.User{
		Id:       uuid.New().String(),
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	// Create the user
	err = sc.userRepo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Create access token
	// TODO - implement access token secret and expiry from ENV
	accessToken, err := tokens.CreateAccessToken(&user, "TODO_ACCESS_TOKEN_SECRET", 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Create refresh token
	// TODO - implement refresh token secret and expiry from ENV
	refreshToken, err := tokens.CreateRefreshToken(&user, "TODO_REFRESH_TOKEN_SECRET", 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Create response and return success status code
	response := models.SignupResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	c.JSON(http.StatusOK, response)
}

// Signin a user using email and password
func (sc *SignupController) Signin(c *gin.Context) {
	// TODO - implement function

	_, cancel := context.WithTimeout(c, sc.timeout)
	defer cancel()

	var request models.SigninRequest

	// Validate that request is properly constructed
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := sc.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	if !password.ComparePasswordHash(request.Password, user.Password) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid email or password"})
		return
	}

	c.JSON(http.StatusInternalServerError, "not implemented")
}
