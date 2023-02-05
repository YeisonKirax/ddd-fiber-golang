package repositories

import (
	"context"
	"time"

	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/dtos"
	"github.com/YeisonKirax/fiber-api-rest/src/modules/users/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	userCollection *mongo.Collection
}

// Newtype returns new type.
func NewUsersRepository(userCollection *mongo.Collection) *UsersRepository {
	return &UsersRepository{userCollection: userCollection}
}

func (ur *UsersRepository) Create(newUser dtos.CreateUserDTO) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := models.NewUser(newUser.Name, newUser.Surname)
	insertionResult, err := ur.userCollection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	var userCreated models.User
	getCreatedUserError := ur.userCollection.FindOne(ctx, bson.D{{Key: "_id", Value: insertionResult.InsertedID}}).Decode(&userCreated)
	if getCreatedUserError != nil {
		return models.User{}, getCreatedUserError
	}
	return userCreated, nil
}
func (ur *UsersRepository) Find() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var users = []models.User{}
	results, err := ur.userCollection.Find(ctx, bson.M{})
	if err != nil {
		return []models.User{}, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		getUserError := results.Decode(&singleUser)
		if getUserError != nil {
			return users, getUserError
		}
		users = append(users, singleUser)
	}
	return users, nil
}

func (ur *UsersRepository) FindById(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user = models.User{}
	objectId, _ := primitive.ObjectIDFromHex(userId)
	err := ur.userCollection.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
