package repository

import (
	"context"
	"fmt"
	"github.com/AhmedSaIah/fiber/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AhmedSaIah/fiber/models"
)

const UsersCollection = "users"

type UserRepository interface {
	Save(user *models.User) error
	//Update(user *models.User) error
	//GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	//GetAll() (users []*models.User, err error)
	//Delete(id string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(conn database.Connection) UserRepository {
	return &userRepository{conn.DB().Collection(UsersCollection)}
}

func (u *userRepository) Save(user *models.User) error {
	_, err := u.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetByEmail(email string) (*models.User, error) {
	filer := bson.M{"email": email}
	opts := options.FindOne()
	result := u.collection.FindOne(context.Background(), filer, opts)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no user found with email provided %w", err)
		}
		return nil, err
	}

	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Delete(email string) error {
	var user models.User
	err := u.collection.FindOneAndDelete(context.Background(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user with that email not found to be deleted: %w", err)
		}
	}
	return nil
}
