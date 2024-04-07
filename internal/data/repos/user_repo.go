package repos

import "github.com/robert-tyssen/go-auth-jwt-demo/internal/models"

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}

type userRepoImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepoImpl{}
}

func (ur *userRepoImpl) CreateUser(user models.User) error {
	// TODO - implement function
	return nil
}

func (ur *userRepoImpl) GetUserByEmail(email string) (models.User, error) {
	// TODO - implement function
	return models.User{}, nil
}
