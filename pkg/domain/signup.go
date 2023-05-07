package domain

import "context"

type SignupRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUseCase interface {

	// Create new user
	CreateUser(c context.Context, user *User) error

	// Create the access token for authentication
	CreateAccessToken(user *User, secret string, expiry int) (token string, err error)

	// Create the refresh token for authentication
	CreateRefreshToken(user *User, secret string, expiry int) (token string, err error)
}
