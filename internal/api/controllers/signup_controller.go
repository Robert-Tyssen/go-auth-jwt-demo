package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

	// Validate that the password meets complexity requirements
	if err := password.ValidatePassword(request.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "weak-password"})
		return
	}

	// Use bcrypt to generate an encrypted password from the user-provided one
	// Only the hashed password is stored in the backend
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Validate that the user does not already exist
	if foundUser, _ := sc.userRepo.GetUserByEmail(request.Email); foundUser.Email == request.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user-already-exists"})
		return
	}

	// Construct user to be created
	var user = models.User{
		Id:       "", // ID is generated by the database
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	// Create the user
	user.Id, err = sc.userRepo.CreateUser(user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user-creation-failed"})
		return
	}

	// Create access token with expiry of 1 hour
	accessToken, err := tokens.CreateAccessToken(&user, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Create refresh token with expiry of 2 hours
	refreshToken, err := tokens.CreateRefreshToken(&user, 2)
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

	_, cancel := context.WithTimeout(c, sc.timeout)
	defer cancel()

	var request models.SigninRequest

	// Validate that request is properly constructed
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Get the user by email
	user, err := sc.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Compare the password hash with the user-provided password
	if !password.ComparePasswordHash(request.Password, user.Password) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid email or password"})
		return
	}

	// TODO - implement function
	c.JSON(http.StatusInternalServerError, "not implemented")
}
