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
func CreateToken(user *models.User, expiryHours int) (string, error) {
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

func TokenIsValid(accessToken string) (bool, error) {
	return tokenIsValid(accessToken, getAccessTokenSecret())
}

func tokenIsValid(requestToken string, secret string) (bool, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(requestToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), errors.New("unexpected-signing-method")
	})

	// Token parsing error
	if err != nil {
		// TODO - additional handling for specific errors
		return false, err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return true, nil
	}

	// Token is invalid
	return false, errors.New("invalid-token")

}

// Get the secret for the access token from the environment
func getAccessTokenSecret() string {
	return os.Getenv("JWT_SECRET")
}
