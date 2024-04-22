package tokens

import (
	"errors"
	"fmt"
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

func TokenIsValid(tokenString string) (bool, error) {
	// Get the secret
	secret := getAccessTokenSecret()

	// Parse the token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret
			return []byte(secret), nil
		},
	)

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
