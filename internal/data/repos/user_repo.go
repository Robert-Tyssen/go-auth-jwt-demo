package repos

import (
	"context"

	"github.com/robert-tyssen/go-auth-jwt-demo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user models.User) (string, error)
	GetUserByEmail(email string) (models.User, error)
}

type userRepoImpl struct {
	userCol *mongo.Collection
}

func NewUserRepository(db *mongo.Client) UserRepository {

	userCol := db.Database("auth").Collection("users")
	return &userRepoImpl{
		userCol: userCol,
	}
}

// Creates a new user in the database, and returns the ID of the new user
// Returns an error if the user could not be created
func (ur *userRepoImpl) CreateUser(user models.User) (string, error) {

	// Create DTO for DB update
	userDto := bson.M{
		"email":    user.Email,
		"password": user.Password,
	}

	// Insert the user into the database
	res, err := ur.userCol.InsertOne(context.Background(), userDto)

	// Return error if insert failed
	if err != nil {
		return "", err
	}

	// Return the ID of the inserted user
	id := res.InsertedID.(primitive.ObjectID).String()
	return id, nil
}

func (ur *userRepoImpl) GetUserByEmail(email string) (models.User, error) {
	// TODO - implement function
	return models.User{}, nil
}
