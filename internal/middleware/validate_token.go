package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/utils/tokens"
)

// ValidateToken is a middleware function that checks if the request contains a valid access token
// If the token is invalid the middleware will return an error response, otherwise it will continue to the next middleware
func ValidateToken(c *gin.Context) {

	// Extract the token string from the Authorization header
	tokenString := getTokenString(c)

	// Validate that the token is present in the HTTP request
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		c.Abort()
		return
	}

	// Check if the token is valid
	_, err := tokens.TokenIsValid(tokenString)

	// If the token is invalid, return an error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	// Token is valid - continue to the next middleware
	c.Next()

}

// Extract the token string from the HTTP authorization header
// The token string is expected to be in the format "Bearer <token>"
func getTokenString(c *gin.Context) string {
	const bearerSchema = "Bearer "
	authHeader := c.GetHeader("Authorization")
	return strings.TrimSpace(strings.TrimPrefix(authHeader, bearerSchema))
}
