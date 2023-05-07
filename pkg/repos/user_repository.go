package repos

import (
	"context"

	"github.com/robert-tyssen/go-auth-jwt-demo/pkg/domain"
)

type userRepository struct{}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	// TODO - implement user creation
	//return errors.New("user-repo-create-not-implemented")
	return nil
}
