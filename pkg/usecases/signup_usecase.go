package usecases

import (
	"context"
	"time"

	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/domain"
	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/internal/tokenutil"
)

type signupUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUseCase(ur domain.UserRepository, timeout time.Duration) domain.SignupUseCase {
	return &signupUseCase{
		userRepository: ur,
		contextTimeout: timeout,
	}
}

func (su *signupUseCase) CreateUser(c context.Context, user *domain.User) error {

	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	return su.userRepository.Create(ctx, user)
}

func (su *signupUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (token string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (token string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
