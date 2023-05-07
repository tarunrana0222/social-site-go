package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id" validate:"required"`
	Name     string             `json:"name" validate:"required"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"password" validate:"required"`
	UserID   string             `json:"userId" validate:"required"`
}
