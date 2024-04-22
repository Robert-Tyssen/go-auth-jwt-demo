package repos

import (
	"context"
	"log"
	"strings"

	"github.com/robert-tyssen/go-auth-jwt-demo/internal/data/dto"
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

// Creates a UserRepository instance for database operations on users
func NewUserRepository(db *mongo.Client) UserRepository {

	// Create a new user repository instance
	repo := &userRepoImpl{
		userCol: db.Database("auth").Collection("users"),
	}

	// Setup indexes
	repo.setupIndexes()

	// Return the repo
	return repo
}

// Creates a new user in the database, and returns the ID of the new user
// Returns an error if the user could not be created
func (ur *userRepoImpl) CreateUser(user models.User) (string, error) {

	// Create DTO for DB update
	userDto := bson.M{
		"email":    strings.ToLower(user.Email),
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

// Fetches a user from the database based on their email address, and returns the user
// Returns an error if the user could not be found
func (ur *userRepoImpl) GetUserByEmail(email string) (models.User, error) {

	// Find the user in the database
	res := ur.userCol.FindOne(context.TODO(), bson.M{
		"email": strings.ToLower(email),
	})

	// Parse the result into a UserReadDto
	var dto = dto.UserReadDto{}
	err := res.Decode(&dto)

	// Convert the UserReadDto to a User and return result
	return dto.ToUser(), err
}


// This function sets up the necessary indexes for collections
// used by the user repository
func (ur *userRepoImpl) setupIndexes() {
	// Create index on email field
	_, err := ur.userCol.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{"email": 1}, Options: nil,
	})

	if err != nil {
		log.Fatal(err)
	}
}
