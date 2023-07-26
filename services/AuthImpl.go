package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AhmedSaIah/fiber/models"
	"github.com/AhmedSaIah/fiber/utils"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthService(collection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{collection, ctx}
}

func (u *AuthServiceImpl) SignUpUser(user *models.SignUpRequest) (*models.DBResponse, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = true
	user.Role = "user"

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	res, err := u.collection.InsertOne(u.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exists")
		}
		return nil, err
	}
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: opt,
	}
	if _, err := u.collection.Indexes().CreateOne(u.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = u.collection.FindOne(u.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *AuthServiceImpl) SignInUser(user *models.SignInRequest) (*models.DBResponse, error) {
	return nil, nil
}
