package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id, omitempty"`
	Name     string             `json:"name," bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type UserRepository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
	Delete(email string) error
}
