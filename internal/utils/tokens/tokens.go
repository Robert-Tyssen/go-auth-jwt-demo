package tokens

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/models"
)

// Create a signed access token for a user, with a specified expiry time
// Returns the token as a string, or an error if the token could not be created
func CreateAccessToken(user *models.User, expiryHours int) (string, error) {
	// Set the expiry time for the token
	exp := time.Now().Add(time.Hour * time.Duration(expiryHours))

	// Create the claims for the token
	claims := &jwt.RegisteredClaims{
		Issuer:    "GO-AUTH-JWT-DEMO",
		ExpiresAt: jwt.NewNumericDate(exp),
		ID:        user.Id,
	}

	// Create a new token with the claims and get the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := getAccessTokenSecret()

	// Sign the token with the secret
	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	// Return the token
	return accessToken, nil
}

// Create a signed refresh token for a user, with a specified expiry time
// Returns the token as a string, or an error if the token could not be created
func CreateRefreshToken(user *models.User, expiryHours int) (string, error) {
	
	// Set the expiry time for the token
	exp := time.Now().Add(time.Hour * time.Duration(expiryHours))

	// Create the claims for the token
	claims := &jwt.RegisteredClaims{
		Issuer:    "GO-AUTH-JWT-DEMO",
		ExpiresAt: jwt.NewNumericDate(exp),
		ID:        user.Id,
	}

	// Create a new token with the claims and get the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := getRefreshTokenSecret()

	// Sign the token with the secret
	refreshToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	// Return the token
	return refreshToken, nil
}

func TokenIsAuthorized(requestToken string, secret string) (bool, error) {
	return false, errors.New("token-auth-not-implemented")
}

func GetIdFromToken(requestToken string, secret string) (string, error) {
	return "", errors.New("id-from-token-not-implemented")
}

// Get the secret for the access token from the environment
func getAccessTokenSecret() string {
	return os.Getenv("JWT_ACCESS_SECRET")
}

// Get the secret for the refresh token from the environment
func getRefreshTokenSecret() string {
	return os.Getenv("JWT_REFRESH_SECRET")
}
