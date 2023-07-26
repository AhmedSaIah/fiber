package database

//
//import (
//	"context"
//	"fmt"
//
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//
//	"github.com/AhmedSaIah/fiber/models"
//)
//
//type UserRepository struct {
//	collection *mongo.Collection
//}
//
//func NewUserRepository(client *mongo.Client) *UserRepository {
//	return &UserRepository{
//		collection: client.Database("fiber").Collection("users"),
//	}
//}
//
//func (u *UserRepository) Save(user *models.User) error {
//	_, err := u.collection.InsertOne(context.Background(), user)
//	if err != nil {
//		fmt.Errorf("error creating user: %w", err)
//	}
//	return err
//}
//
//func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
//	var user models.User
//	err := u.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
//
//	if err != nil {
//		if err == mongo.ErrNoDocuments {
//			return nil, fmt.Errorf("error finding user with that email: %w", err)
//		}
//	}
//	return &user, nil
//}
//
//func (u *UserRepository) Delete(email string) error {
//	var user models.User
//	err := u.collection.FindOneAndDelete(context.Background(), bson.M{"email": email}).Decode(&user)
//
//	if err != nil {
//		if err == mongo.ErrNoDocuments {
//			return fmt.Errorf("user with that email not found to be deleted: %w", err)
//		}
//	}
//	return nil
//}
