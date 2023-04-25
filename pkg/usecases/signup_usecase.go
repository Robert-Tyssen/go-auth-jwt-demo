package usecases

import (
	"context"
	"time"

	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/domain"
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
