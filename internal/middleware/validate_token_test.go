package middleware

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTokenString(t *testing.T) {
	tests := []struct {
		name           string
		authorization  string
		expectedResult string
	}{
		{
			name:           "Valid authorization header",
			authorization:  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			expectedResult: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
		},
		{
			name:           "Empty authorization header",
			authorization:  "",
			expectedResult: "",
		},
		{
			name:           "No Bearer schema",
			authorization:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			expectedResult: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &gin.Context{}
			c.Request = &http.Request{}
			c.Request.Header = http.Header{}
			c.Request.Header.Set("Authorization", test.authorization)

			result := getTokenString(c)
			if result != test.expectedResult {
				t.Errorf("Expected %s, but got %s", test.expectedResult, result)
			}
		})
	}
}