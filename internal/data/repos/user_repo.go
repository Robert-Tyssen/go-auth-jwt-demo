package repos

import "github.com/robert-tyssen/go-auth-jwt-demo/internal/models"

type UserRepository interface {
	CreateUser(user models.User) error
}

type userRepoImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepoImpl{}
}

func (ur *userRepoImpl) CreateUser(user models.User) error {
	return nil
}
