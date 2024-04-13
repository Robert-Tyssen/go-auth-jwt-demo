package dto

import (
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Data transfer object for reading user data from MongoDB
type UserReadDto struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

// Convert a UserReadDto to a User
func (dto *UserReadDto) ToUser() models.User {
	return models.User{
		Id:       dto.ID.String(),
		Email:    dto.Email,
		Password: dto.Password,
	}
}
