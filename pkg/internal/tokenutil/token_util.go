package tokenutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/domain"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	claims := &jwt.RegisteredClaims{
		Issuer:    "GO-AUTH-JWT-DEMO",
		ExpiresAt: jwt.NewNumericDate(exp),
		ID:        user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	claims := &jwt.RegisteredClaims{
		Issuer:    "GO-AUTH-JWT-DEMO",
		ExpiresAt: jwt.NewNumericDate(exp),
		ID:        user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func TokenIsAuthorized(requestToken string, secret string) (bool, error) {
	return false, errors.New("token-auth-not-implemented")
}

func GetIdFromToken(requestToken string, secret string) (string, error) {
	return "", errors.New("id-from-token-not-implemented")
}
